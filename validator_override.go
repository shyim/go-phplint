package phplint

import (
	"fmt"
	"strings"

	"github.com/shyim/go-phplint/internal/ast"
)

// linkedClass contains only declarations PHP can bind while linting. Parent
// pointers capture the class table at the point of declaration.
type linkedClass struct {
	name, kind      string
	parent          *linkedClass
	final, readonly bool
	members         map[memberKey]classMember
}

type memberKey struct {
	kind, name string
}

type classMember struct {
	owner               string
	visibility          int
	final, static       bool
	abstractConstructor bool
}

func (c *linkedClass) lookup(key memberKey) (classMember, bool) {
	for ; c != nil; c = c.parent {
		if member, ok := c.members[key]; ok {
			return member, true
		}
	}
	return classMember{}, false
}

func (v *compileValidator) isOverride(attr *ast.Attribute) bool {
	name, _ := v.resolver.Namespace.ResolveName(attr.Name, "")
	return strings.EqualFold(name, "Override")
}

func (v *compileValidator) hasOverride(groups []ast.Vertex) bool {
	if v.version < PHP83 {
		return false
	}
	for _, group := range groups {
		for _, attr := range group.(*ast.AttributeGroup).Attrs {
			if v.isOverride(attr.(*ast.Attribute)) {
				return true
			}
		}
	}
	return false
}

// The attribute is visited inside its group, immediately below the target.
func (v *compileValidator) validateOverrideTarget(attr *ast.Attribute) {
	if v.version < PHP83 || !v.isOverride(attr) || len(v.stack) < 2 {
		return
	}
	target := v.stack[len(v.stack)-2]
	switch n := target.(type) {
	case *ast.StmtClassMethod:
	case *ast.Parameter:
		if len(n.Modifiers) == 0 && len(n.Hooks) == 0 {
			v.report(attr, "the #[Override] attribute cannot target a parameter")
			return
		}
		v.requireVersion(attr, PHP85, "#[Override] attributes on properties")
	case *ast.StmtPropertyList:
		v.requireVersion(attr, PHP85, "#[Override] attributes on properties")
	case *ast.StmtClassConstList, *ast.EnumCase:
		v.requireVersion(attr, PHP86, "#[Override] attributes on class constants and enum cases")
	default:
		v.report(attr, "invalid target for the #[Override] attribute")
	}
}

func (v *compileValidator) bindClass(node, name ast.Vertex, kind string, extends ast.Vertex, stmts []ast.Vertex, flags map[string]bool, hasInterfaces bool) {
	if hasInterfaces {
		return
	}
	for _, stmt := range stmts {
		if _, ok := stmt.(*ast.StmtTraitUse); ok {
			return
		}
	}

	var topLevel bool
	if len(v.stack) > 0 {
		switch v.stack[len(v.stack)-1].(type) {
		case *ast.Root, *ast.StmtNamespace:
			topLevel = name != nil
		}
	}

	display := nodeName(name)
	if ns := v.resolver.Namespace.Namespace; ns != "" && name != nil {
		display = ns + `\` + display
	}
	class := &linkedClass{name: display, kind: kind, final: flags["final"], readonly: flags["readonly"]}
	if extends != nil {
		if !topLevel {
			return
		}
		base, _ := v.resolver.Namespace.ResolveName(extends, "")
		class.parent = v.classes[strings.ToLower(base)]
		if class.parent == nil {
			return
		}
		parent := class.parent
		switch {
		case parent.kind != "class":
			v.report(node, fmt.Sprintf("class %s cannot extend %s %s", display, parent.kind, parent.name))
			return
		case parent.final:
			v.report(node, fmt.Sprintf("class %s cannot extend final class %s", display, parent.name))
			return
		case class.readonly != parent.readonly:
			v.report(node, fmt.Sprintf("class %s and parent %s must agree on readonly", display, parent.name))
			return
		}
	}

	// A future child only needs member information for an extensible class.
	saveMembers := topLevel && kind == "class" && !class.final
	if kind != "trait" {
		for _, stmt := range stmts {
			v.checkClassMembers(class, stmt, saveMembers)
		}
	}
	if topLevel {
		if v.classes == nil {
			v.classes = make(map[string]*linkedClass)
		}
		key := strings.ToLower(display)
		if _, exists := v.classes[key]; !exists {
			v.classes[key] = class
		}
	}
}

func (v *compileValidator) checkClassMembers(class *linkedClass, stmt ast.Vertex, save bool) {
	// Most members of a standalone final class need no inheritance checks.
	check := func(groups []ast.Vertex) (override, needed bool) {
		override = v.hasOverride(groups)
		return override, override || save || class.parent != nil
	}
	switch n := stmt.(type) {
	case *ast.StmtClassMethod:
		if override, needed := check(n.AttrGroups); needed {
			v.checkMember(class, n, "method", nodeName(n.Name), n.Modifiers, override, save)
		}
		if !strings.EqualFold(nodeName(n.Name), "__construct") {
			return
		}
		for _, param := range n.Params {
			p := param.(*ast.Parameter)
			if len(p.Modifiers) == 0 && len(p.Hooks) == 0 {
				continue
			}
			if override, needed := check(p.AttrGroups); needed {
				v.checkMember(class, p, "property", variableName(p.Var), p.Modifiers, override, save)
			}
		}
	case *ast.StmtPropertyList:
		if override, needed := check(n.AttrGroups); needed {
			for _, p := range n.Props {
				v.checkMember(class, p, "property", propertyName(p), n.Modifiers, override, save)
			}
		}
	case *ast.StmtClassConstList:
		if override, needed := check(n.AttrGroups); needed {
			for _, c := range n.Consts {
				v.checkMember(class, c, "constant", constantName(c), n.Modifiers, override, save)
			}
		}
	}
}

func (v *compileValidator) checkMember(class *linkedClass, node ast.Vertex, kind, name string, modifiers []ast.Vertex, override, save bool) {
	key := memberKey{kind, name}
	if kind == "method" {
		key.name = strings.ToLower(name)
	}
	flags := modifierSet(modifiers)
	visibility, ok := plainVisibility(flags)
	if !ok {
		visibility = "public"
	}
	parent, found := class.parent.lookup(key)
	constructor := kind == "method" && key.name == "__construct"
	member := classMember{
		owner:               class.name,
		visibility:          visibilityStrength(visibility),
		static:              flags["static"],
		final:               flags["final"] || (kind == "property" && flags["private(set)"] && visibility != "private"),
		abstractConstructor: constructor && (flags["abstract"] || parent.abstractConstructor),
	}
	if save {
		if class.members == nil {
			class.members = make(map[memberKey]classMember)
		}
		class.members[key] = member
	}

	private := parent.visibility == visibilityStrength("private")
	concreteConstructor := constructor && !parent.abstractConstructor
	if override && (!found || private || concreteConstructor) {
		v.report(node, fmt.Sprintf("%s::%s has #[\\Override] attribute, but no matching parent %s exists", class.name, name, kind))
	}
	if !found {
		return
	}
	if parent.final && !(private && kind == "method" && !constructor && v.version >= PHP80) {
		v.report(node, fmt.Sprintf("cannot override final %s %s::%s", kind, parent.owner, name))
	}
	if private && (kind != "method" || v.version >= PHP80) {
		return
	}
	if member.static != parent.static {
		v.report(node, fmt.Sprintf("static modifier on %s::%s must match %s", class.name, name, parent.owner))
	}
	if !private && !concreteConstructor && member.visibility < parent.visibility {
		v.report(node, fmt.Sprintf("cannot restrict visibility of %s::%s", parent.owner, name))
	}
}
