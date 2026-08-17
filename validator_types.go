package phplint

import (
	"fmt"
	"strings"

	"github.com/shyim/go-phplint/internal/ast"
)

type typeContext uint8

const (
	typeContextParameter typeContext = iota + 1
	typeContextProperty
	typeContextReturn
)

func (v *compileValidator) validateNullable(nullable *ast.Nullable) {
	name := strings.ToLower(simpleTypeName(nullable.Expr))
	switch name {
	case "mixed":
		if v.version < PHP80 {
			return
		}
	case "never":
		if v.version < PHP81 {
			return
		}
	case "null":
		if v.version < PHP80 {
			return
		}
	case "void":
	default:
		return
	}
	if name != "" {
		v.report(nullable, fmt.Sprintf("type %s cannot be nullable", name))
	}
}

func (v *compileValidator) validateUnion(union *ast.Union) {
	seen := make(map[string]bool)
	hasIntersection := false
	for _, member := range union.Types {
		if _, ok := member.(*ast.Intersection); ok {
			hasIntersection = true
		}
		name := strings.ToLower(simpleTypeName(member))
		if name == "" {
			continue
		}
		if seen[name] {
			v.report(member, fmt.Sprintf("duplicate type %s is redundant", name))
		}
		seen[name] = true
	}

	if hasIntersection && v.version < PHP82 {
		v.report(union, "disjunctive normal form types require PHP 8.2")
	}
	if seen["mixed"] && len(seen) > 1 {
		v.report(union, "mixed cannot be part of a union type")
	}
	if seen["void"] {
		v.report(union, "void cannot be part of a union type")
	}
	if seen["never"] {
		v.report(union, "never cannot be part of a union type")
	}
	if seen["bool"] && (seen["false"] || seen["true"]) {
		v.report(union, "bool makes true or false redundant in a union type")
	}
	if seen["iterable"] && (seen["array"] || seen["traversable"]) {
		v.report(union, "iterable makes array or Traversable redundant in a union type")
	}
}

func (v *compileValidator) validateIntersection(intersection *ast.Intersection) {
	v.requireVersion(intersection, PHP81, "intersection types")

	seen := make(map[string]bool)
	for _, member := range intersection.Types {
		name := strings.ToLower(simpleTypeName(member))
		if name == "" {
			continue
		}
		if isBuiltinIntersectionType(name) {
			v.report(member, fmt.Sprintf("type %s cannot be part of an intersection type", name))
		}
		if seen[name] {
			v.report(member, fmt.Sprintf("duplicate type %s is redundant", name))
		}
		seen[name] = true
	}
}

func (v *compileValidator) validateType(node ast.Vertex, context typeContext) {
	if node == nil {
		return
	}

	switch current := node.(type) {
	case *ast.Nullable:
		v.validateNullable(current)
		v.validateType(current.Expr, context)
	case *ast.Union:
		v.validateUnion(current)
		for _, member := range current.Types {
			v.validateType(member, context)
		}
	case *ast.Intersection:
		v.validateIntersection(current)
		for _, member := range current.Types {
			v.validateType(member, context)
		}
	default:
		v.validateSimpleType(node, context)
	}
}

func (v *compileValidator) validateSimpleType(node ast.Vertex, context typeContext) {
	name := strings.ToLower(strings.TrimPrefix(simpleTypeName(node), `\`))
	if name == "" {
		return
	}

	switch name {
	case "true":
		v.requireVersion(node, PHP82, "the true type")
	case "false", "null":
		v.requireVersion(node, PHP82, fmt.Sprintf("standalone %s types", name))
	case "never":
		v.requireVersion(node, PHP81, "the never return type")
		if context != typeContextReturn {
			v.report(node, "never is only permitted as a return type")
		}
	case "void":
		if context != typeContextReturn {
			v.report(node, "void is only permitted as a return type")
		}
	case "callable":
		if context == typeContextProperty {
			v.report(node, "properties cannot have type callable")
		}
	}
}

func simpleTypeName(node ast.Vertex) string {
	switch current := node.(type) {
	case nil:
		return ""
	case *ast.Identifier:
		return string(current.Value)
	case *ast.NamePart:
		return string(current.Value)
	case *ast.Name:
		return joinNameParts(current.Parts)
	case *ast.NameFullyQualified:
		return `\` + joinNameParts(current.Parts)
	case *ast.NameRelative:
		return `namespace\` + joinNameParts(current.Parts)
	case *ast.Nullable:
		return simpleTypeName(current.Expr)
	default:
		return ""
	}
}

func isBuiltinIntersectionType(name string) bool {
	switch strings.TrimPrefix(name, `\`) {
	case "array", "bool", "callable", "false", "float", "int", "iterable",
		"mixed", "never", "null", "object", "string", "true", "void":
		return true
	default:
		return false
	}
}
