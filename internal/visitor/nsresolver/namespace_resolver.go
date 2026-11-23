// Package visitor contains walker.visitor implementations
package nsresolver

import (
	"bytes"
	"errors"
	"strings"

	"github.com/shyim/go-phplint/internal/ast"
	"github.com/shyim/go-phplint/internal/visitor"
)

// NamespaceResolver visitor
type NamespaceResolver struct {
	visitor.Null
	Namespace     *Namespace
	ResolvedNames map[ast.Vertex]string

	goDeep bool
}

// NewNamespaceResolver NamespaceResolver type constructor
func NewNamespaceResolver() *NamespaceResolver {
	return &NamespaceResolver{
		Namespace:     NewNamespace(""),
		ResolvedNames: map[ast.Vertex]string{},
		goDeep:        true,
	}
}

func (nsr *NamespaceResolver) EnterNode(n ast.Vertex) bool {
	n.Accept(nsr)

	if !nsr.goDeep {
		nsr.goDeep = true
		return false
	}

	return true
}

func (nsr *NamespaceResolver) StmtNamespace(n *ast.StmtNamespace) {
	if n.Name == nil {
		nsr.Namespace = NewNamespace("")
	} else {
		NSParts := nameNodeParts(n.Name)
		nsr.Namespace = NewNamespace(concatNameParts(NSParts))
	}
}

func (nsr *NamespaceResolver) StmtUse(n *ast.StmtUseList) {
	useType := ""
	if n.Type != nil {
		useType = string(n.Type.(*ast.Identifier).Value)
	}

	for _, nn := range n.Uses {
		nsr.AddAlias(useType, nn, nil)
	}

	nsr.goDeep = false
}

func (nsr *NamespaceResolver) StmtGroupUse(n *ast.StmtGroupUseList) {
	useType := ""
	if n.Type != nil {
		useType = string(n.Type.(*ast.Identifier).Value)
	}

	for _, nn := range n.Uses {
		nsr.AddAlias(useType, nn, nameNodeParts(n.Prefix))
	}

	nsr.goDeep = false
}

func (nsr *NamespaceResolver) StmtClass(n *ast.StmtClass) {
	if n.Extends != nil {
		nsr.ResolveName(n.Extends, "")
	}

	if n.Implements != nil {
		for _, interfaceName := range n.Implements {
			nsr.ResolveName(interfaceName, "")
		}
	}

	if n.Name != nil {
		nsr.AddNamespacedName(n, string(n.Name.(*ast.Identifier).Value))
	}
}

func (nsr *NamespaceResolver) StmtInterface(n *ast.StmtInterface) {
	if n.Extends != nil {
		for _, interfaceName := range n.Extends {
			nsr.ResolveName(interfaceName, "")
		}
	}

	nsr.AddNamespacedName(n, string(n.Name.(*ast.Identifier).Value))
}

func (nsr *NamespaceResolver) StmtTrait(n *ast.StmtTrait) {
	nsr.AddNamespacedName(n, string(n.Name.(*ast.Identifier).Value))
}

func (nsr *NamespaceResolver) StmtEnum(n *ast.StmtEnum) {
	for _, interfaceName := range n.Implements {
		nsr.ResolveName(interfaceName, "")
	}

	nsr.AddNamespacedName(n, string(n.Name.(*ast.Identifier).Value))
}

// Attribute resolves attribute names like class names so callers can
// recognize internal attributes such as \Override regardless of how they
// are spelled or imported.
func (nsr *NamespaceResolver) Attribute(n *ast.Attribute) {
	nsr.ResolveName(n.Name, "")
}

func (nsr *NamespaceResolver) StmtFunction(n *ast.StmtFunction) {
	nsr.AddNamespacedName(n, string(n.Name.(*ast.Identifier).Value))

	for _, parameter := range n.Params {
		nsr.ResolveType(parameter.(*ast.Parameter).Type)
	}

	if n.ReturnType != nil {
		nsr.ResolveType(n.ReturnType)
	}
}

func (nsr *NamespaceResolver) StmtClassMethod(n *ast.StmtClassMethod) {
	for _, parameter := range n.Params {
		nsr.ResolveType(parameter.(*ast.Parameter).Type)
	}

	if n.ReturnType != nil {
		nsr.ResolveType(n.ReturnType)
	}
}

func (nsr *NamespaceResolver) ExprClosure(n *ast.ExprClosure) {
	for _, parameter := range n.Params {
		nsr.ResolveType(parameter.(*ast.Parameter).Type)
	}

	if n.ReturnType != nil {
		nsr.ResolveType(n.ReturnType)
	}
}

func (nsr *NamespaceResolver) StmtPropertyList(n *ast.StmtPropertyList) {
	if n.Type != nil {
		nsr.ResolveType(n.Type)
	}
}

func (nsr *NamespaceResolver) StmtConstList(n *ast.StmtConstList) {
	for _, constant := range n.Consts {
		nsr.AddNamespacedName(constant, string(constant.(*ast.StmtConstant).Name.(*ast.Identifier).Value))
	}
}

func (nsr *NamespaceResolver) ExprStaticCall(n *ast.ExprStaticCall) {
	nsr.ResolveName(n.Class, "")
}

func (nsr *NamespaceResolver) ExprStaticPropertyFetch(n *ast.ExprStaticPropertyFetch) {
	nsr.ResolveName(n.Class, "")
}

func (nsr *NamespaceResolver) ExprClassConstFetch(n *ast.ExprClassConstFetch) {
	nsr.ResolveName(n.Class, "")
}

func (nsr *NamespaceResolver) ExprNew(n *ast.ExprNew) {
	nsr.ResolveName(n.Class, "")
}

func (nsr *NamespaceResolver) ExprInstanceOf(n *ast.ExprInstanceOf) {
	nsr.ResolveName(n.Class, "")
}

func (nsr *NamespaceResolver) StmtCatch(n *ast.StmtCatch) {
	for _, t := range n.Types {
		nsr.ResolveName(t, "")
	}
}

func (nsr *NamespaceResolver) ExprFunctionCall(n *ast.ExprFunctionCall) {
	nsr.ResolveName(n.Function, "function")
}

func (nsr *NamespaceResolver) ExprConstFetch(n *ast.ExprConstFetch) {
	nsr.ResolveName(n.Const, "const")
}

func (nsr *NamespaceResolver) StmtTraitUse(n *ast.StmtTraitUse) {
	for _, t := range n.Traits {
		nsr.ResolveName(t, "")
	}

	for _, a := range n.Adaptations {
		switch aa := a.(type) {
		case *ast.StmtTraitUsePrecedence:
			refTrait := aa.Trait
			if refTrait != nil {
				nsr.ResolveName(refTrait, "")
			}
			for _, insteadOf := range aa.Insteadof {
				nsr.ResolveName(insteadOf, "")
			}

		case *ast.StmtTraitUseAlias:
			refTrait := aa.Trait
			if refTrait != nil {
				nsr.ResolveName(refTrait, "")
			}
		}
	}
}

// LeaveNode is invoked after node process
func (nsr *NamespaceResolver) LeaveNode(n ast.Vertex) {
	switch nn := n.(type) {
	case *ast.StmtNamespace:
		if nn.Stmts != nil {
			nsr.Namespace = NewNamespace("")
		}
	}
}

// AddAlias adds a new alias
func (nsr *NamespaceResolver) AddAlias(useType string, nn ast.Vertex, prefix []ast.Vertex) {
	switch use := nn.(type) {
	case *ast.StmtUse:
		if use.Type != nil {
			useType = string(use.Type.(*ast.Identifier).Value)
		}

		useNameParts := nameNodeParts(use.Use)
		var alias string
		if use.Alias == nil {
			alias = string(useNameParts[len(useNameParts)-1].(*ast.NamePart).Value)
		} else {
			alias = string(use.Alias.(*ast.Identifier).Value)
		}

		nsr.Namespace.AddAlias(useType, concatNameParts(prefix, useNameParts), alias)
	}
}

// AddNamespacedName adds namespaced name by node
func (nsr *NamespaceResolver) AddNamespacedName(nn ast.Vertex, nodeName string) {
	if nsr.Namespace.Namespace == "" {
		nsr.ResolvedNames[nn] = nodeName
	} else {
		nsr.ResolvedNames[nn] = nsr.Namespace.Namespace + "\\" + nodeName
	}
}

// ResolveName adds a resolved fully qualified name by node
func (nsr *NamespaceResolver) ResolveName(nameNode ast.Vertex, aliasType string) {
	resolved, err := nsr.Namespace.ResolveName(nameNode, aliasType)
	if err == nil {
		nsr.ResolvedNames[nameNode] = resolved
	}
}

// ResolveType adds a resolved fully qualified type name
func (nsr *NamespaceResolver) ResolveType(n ast.Vertex) {
	switch nn := n.(type) {
	case *ast.Nullable:
		nsr.ResolveType(nn.Expr)
	case *ast.Name:
		nsr.ResolveName(n, "")
	case *ast.NameRelative:
		nsr.ResolveName(n, "")
	case *ast.NameFullyQualified:
		nsr.ResolveName(n, "")
	}
}

// Namespace context
type Namespace struct {
	Namespace string
	Aliases   map[string]map[string]string
}

// NewNamespace constructor
func NewNamespace(NSName string) *Namespace {
	return &Namespace{
		Namespace: NSName,
		Aliases: map[string]map[string]string{
			"":         {},
			"const":    {},
			"function": {},
		},
	}
}

// AddAlias adds a new alias
func (ns *Namespace) AddAlias(aliasType string, aliasName string, alias string) {
	aliasType = strings.ToLower(aliasType)

	if aliasType == "const" {
		ns.Aliases[aliasType][alias] = aliasName
	} else {
		ns.Aliases[aliasType][strings.ToLower(alias)] = aliasName
	}
}

// ResolveName returns a resolved fully qualified name
func (ns *Namespace) ResolveName(nameNode ast.Vertex, aliasType string) (string, error) {
	switch n := nameNode.(type) {
	case *ast.NameFullyQualified:
		// Fully qualifid name is already resolved
		return concatNameParts(nameNodeParts(n)), nil

	case *ast.NameRelative:
		parts := nameNodeParts(n)
		if ns.Namespace == "" {
			return concatNameParts(parts), nil
		}
		return ns.Namespace + "\\" + concatNameParts(parts), nil

	case *ast.Name:
		parts := nameNodeParts(n)
		if aliasType == "const" && len(parts) == 1 {
			part := strings.ToLower(string(parts[0].(*ast.NamePart).Value))
			if part == "true" || part == "false" || part == "null" {
				return part, nil
			}
		}

		if aliasType == "" && len(parts) == 1 {
			part := strings.ToLower(string(parts[0].(*ast.NamePart).Value))

			switch part {
			case "self":
				fallthrough
			case "static":
				fallthrough
			case "parent":
				fallthrough
			case "int":
				fallthrough
			case "float":
				fallthrough
			case "bool":
				fallthrough
			case "string":
				fallthrough
			case "void":
				fallthrough
			case "iterable":
				fallthrough
			case "object":
				return part, nil
			}
		}

		aliasName, err := ns.ResolveAlias(nameNode, aliasType)
		if err != nil {
			// resolve as relative name if alias not found
			if ns.Namespace == "" {
				return concatNameParts(parts), nil
			}
			return ns.Namespace + "\\" + concatNameParts(parts), nil
		}

		if len(parts) > 1 {
			// if name qualified, replace first part by alias
			return aliasName + "\\" + concatNameParts(parts[1:]), nil
		}

		return aliasName, nil
	}

	return "", errors.New("must be instance of name.Names")
}

// ResolveAlias returns alias or error if not found
func (ns *Namespace) ResolveAlias(nameNode ast.Vertex, aliasType string) (string, error) {
	aliasType = strings.ToLower(aliasType)
	nameParts := nameNodeParts(nameNode)
	if len(nameParts) == 0 {
		return "", errors.New("empty name")
	}

	firstPartStr := string(nameParts[0].(*ast.NamePart).Value)

	if len(nameParts) > 1 { // resolve aliases for qualified names, always against class alias type
		firstPartStr = strings.ToLower(firstPartStr)
		aliasType = ""
	} else {
		if aliasType != "const" { // constants are case-sensitive
			firstPartStr = strings.ToLower(firstPartStr)
		}
	}

	aliasName, ok := ns.Aliases[aliasType][firstPartStr]
	if !ok {
		return "", errors.New("Not found")
	}

	return aliasName, nil
}

func concatNameParts(parts ...[]ast.Vertex) string {
	str := ""

	for _, p := range parts {
		for _, n := range p {
			if str == "" {
				str = string(n.(*ast.NamePart).Value)
			} else {
				str = str + "\\" + string(n.(*ast.NamePart).Value)
			}
		}
	}

	return str
}

// nameNodeParts returns the parts of a name node. Parsers running without
// fidelity store only the raw Value, so synthesize parts from it when the
// Parts slice is empty. Leading separators and the relative "namespace"
// prefix are stripped exactly like the fidelity-mode builders do.
func nameNodeParts(node ast.Vertex) []ast.Vertex {
	var value []byte
	switch n := node.(type) {
	case *ast.Name:
		if len(n.Parts) > 0 {
			return n.Parts
		}
		value = n.Value
	case *ast.NameFullyQualified:
		if len(n.Parts) > 0 {
			return n.Parts
		}
		value = bytes.TrimPrefix(n.Value, []byte("\\"))
	case *ast.NameRelative:
		if len(n.Parts) > 0 {
			return n.Parts
		}
		value = n.Value
		if len(value) >= 9 && strings.EqualFold(string(value[:9]), "namespace") {
			value = value[9:]
		}
		value = bytes.TrimPrefix(value, []byte("\\"))
	default:
		return nil
	}

	if len(value) == 0 {
		return nil
	}
	segments := bytes.Split(value, []byte("\\"))
	parts := make([]ast.Vertex, 0, len(segments))
	for _, segment := range segments {
		if len(segment) == 0 {
			continue
		}
		parts = append(parts, &ast.NamePart{Value: segment})
	}
	return parts
}
