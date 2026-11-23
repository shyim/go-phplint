package phplint

import (
	"reflect"
	"strings"

	"github.com/shyim/go-phplint/internal/ast"
)

type constantContext uint8

const (
	constantContextGlobalConstant constantContext = iota + 1
	constantContextClassConstant
	constantContextProperty
	constantContextParameter
	constantContextStatic
)

func isConstantExpression(node ast.Vertex, version Version, context constantContext) bool {
	if node == nil {
		return true
	}

	switch current := node.(type) {
	case *ast.Identifier,
		*ast.Name,
		*ast.NameFullyQualified,
		*ast.NameRelative,
		*ast.NamePart,
		*ast.ScalarDnumber,
		*ast.ScalarEncapsedStringPart,
		*ast.ScalarLnumber,
		*ast.ScalarMagicConstant,
		*ast.ScalarString:
		return true

	case *ast.ExprClosure:
		return version >= PHP85 && current.StaticTkn != nil && len(current.Uses) == 0
	case *ast.ExprArrowFunction:
		return false
	case *ast.ExprFunctionCall:
		return constantCallableName(current.Function) && constantCallableArgs(current.EllipsisTkn != nil, current.Args, version, context)
	case *ast.ExprStaticCall:
		return !strings.EqualFold(nodeName(current.Class), "static") && constantCallableName(current.Class) && constantCallableName(current.Call) &&
			constantCallableArgs(current.EllipsisTkn != nil, current.Args, version, context)
	case *ast.ExprExit:
		return constantCallableArgs(current.EllipsisTkn != nil, current.Args, version, context)
	case *ast.ExprClone:
		return constantCallableArgs(current.EllipsisTkn != nil, current.Args, version, context)
	case *ast.ExprNew:
		if version < PHP81 || !newAllowedInConstantContext(context) {
			return false
		}
	case *ast.Argument:
		// Argument unpacking is not supported in constant expressions.
		if current.VariadicTkn != nil {
			return false
		}
		return isConstantExpression(current.Expr, version, context)
	case *ast.ExprPropertyFetch:
		if version < PHP82 {
			return false
		}
		if _, ok := current.Var.(*ast.ExprClassConstFetch); !ok {
			return false
		}
	case *ast.ExprNullsafePropertyFetch:
		return false
	}

	kind := node.GetType()
	switch {
	case kind == ast.TypeScalarEncapsed,
		kind == ast.TypeScalarHeredoc,
		kind == ast.TypeExprArray,
		kind == ast.TypeExprArrayDimFetch,
		kind == ast.TypeExprArrayItem,
		kind == ast.TypeExprBitwiseNot,
		kind == ast.TypeExprBooleanNot,
		kind == ast.TypeExprBrackets,
		kind == ast.TypeExprClassConstFetch,
		kind == ast.TypeExprConstFetch,
		kind == ast.TypeExprTernary,
		kind == ast.TypeExprUnaryMinus,
		kind == ast.TypeExprUnaryPlus,
		kind == ast.TypeExprNew,
		kind == ast.TypeExprPropertyFetch:
		return constantChildrenAllowed(node, version, context)

	case kind >= ast.TypeExprBinaryBitwiseAnd && kind <= ast.TypeExprBinarySpaceship:
		return constantChildrenAllowed(node, version, context)

	case kind >= ast.TypeExprCastArray && kind <= ast.TypeExprCastUnset:
		return version >= PHP85 && constantChildrenAllowed(node, version, context)

	default:
		return false
	}
}

func newAllowedInConstantContext(context constantContext) bool {
	switch context {
	case constantContextGlobalConstant, constantContextParameter, constantContextStatic:
		return true
	default:
		return false
	}
}

func constantChildrenAllowed(node ast.Vertex, version Version, context constantContext) bool {
	value := reflect.ValueOf(node)
	if value.Kind() == reflect.Pointer {
		if value.IsNil() {
			return true
		}
		value = value.Elem()
	}

	vertexType := reflect.TypeOf((*ast.Vertex)(nil)).Elem()
	for index := 0; index < value.NumField(); index++ {
		field := value.Field(index)
		fieldType := value.Type().Field(index).Type

		if fieldType.Implements(vertexType) {
			if field.IsNil() {
				continue
			}
			child, ok := field.Interface().(ast.Vertex)
			if ok && !isConstantExpression(child, version, context) {
				return false
			}
			continue
		}

		if field.Kind() != reflect.Slice ||
			!fieldType.Elem().Implements(vertexType) {
			continue
		}
		for itemIndex := 0; itemIndex < field.Len(); itemIndex++ {
			child, ok := field.Index(itemIndex).Interface().(ast.Vertex)
			if ok && !isConstantExpression(child, version, context) {
				return false
			}
		}
	}

	return true
}

// PHP accepts ordinary names and literal string targets, including concatenated
// string literals. A variable or a constant fetch is a dynamic target.
func constantCallableName(node ast.Vertex) bool {
	switch n := node.(type) {
	case *ast.Name, *ast.NameRelative, *ast.NameFullyQualified, *ast.Identifier, *ast.ScalarString:
		return true
	case *ast.ExprBrackets:
		return constantCallableName(n.Expr)
	case *ast.ExprBinaryConcat:
		return literalConcatOperand(n.Left) && literalConcatOperand(n.Right)
	}
	return false
}

// Numeric literals in a concatenation are converted to strings by the parser.
func literalConcatOperand(node ast.Vertex) bool {
	switch n := node.(type) {
	case *ast.ScalarLnumber, *ast.ScalarDnumber:
		return true
	case *ast.ExprBrackets:
		return literalConcatOperand(n.Expr)
	default:
		return constantCallableName(node)
	}
}

func constantCallableArgs(ellipsis bool, args []ast.Vertex, version Version, context constantContext) bool {
	if ellipsis {
		return version >= PHP85
	}
	if version < PHP86 {
		return false
	}
	hasPlaceholder := false
	for _, argNode := range args {
		argument, ok := argNode.(*ast.Argument)
		if !ok {
			return false
		}
		if argument.PlaceholderTkn != nil {
			hasPlaceholder = true
			continue
		}
		if !isConstantExpression(argNode, version, context) {
			return false
		}
	}
	return hasPlaceholder
}
