package phplint

import (
	"reflect"

	"github.com/shyim/go-phplint/internal/php/pkg/ast"
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
		return version >= PHP85 && current.EllipsisTkn != nil
	case *ast.ExprMethodCall:
		return version >= PHP85 && current.EllipsisTkn != nil
	case *ast.ExprStaticCall:
		return version >= PHP85 && current.EllipsisTkn != nil
	case *ast.ExprNew:
		if version < PHP81 || !newAllowedInConstantContext(context) {
			return false
		}
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
