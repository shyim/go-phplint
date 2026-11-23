package phplint

import (
	"github.com/shyim/go-phplint/internal/ast"
)

// validateWriteTarget mirrors PHP's compile-time rejection of writes to
// temporary expressions. Fetch chains rooted at a constant fetch are not
// writable: PHP 8.6 lifted the restriction for object property writes
// (const_object_property_write RFC), but direct array writes on constants
// and nullsafe writes remain invalid in every version.
func (v *compileValidator) validateWriteTarget(target ast.Vertex) {
	if target == nil {
		return
	}

	// Destructuring assignments write to every element.
	switch current := target.(type) {
	case *ast.ExprArray:
		for _, item := range current.Items {
			if element, ok := item.(*ast.ExprArrayItem); ok {
				v.validateWriteTarget(element.Val)
			}
		}
		return
	case *ast.ExprList:
		for _, item := range current.Items {
			if element, ok := item.(*ast.ExprArrayItem); ok {
				v.validateWriteTarget(element.Val)
			} else {
				v.validateWriteTarget(item)
			}
		}
		return
	}

	// PHP 8.6 allows mutating objects stored in constants: any fetch chain
	// that passes through a property fetch is writable. Pure array chains
	// modify the constant itself and stay invalid in every version.
	sawProperty := false
	current := target
	for {
		switch fetch := current.(type) {
		case *ast.ExprPropertyFetch:
			sawProperty = true
			current = fetch.Var
		case *ast.ExprBrackets:
			current = fetch.Expr
		case *ast.ExprMethodCall, *ast.ExprStaticCall, *ast.ExprStaticPropertyFetch:
			if hasNullsafeReceiver(current) {
				v.report(target, "Cannot use nullsafe operator in write context")
			}
			return
		case *ast.ExprNullsafePropertyFetch, *ast.ExprNullsafeMethodCall:
			v.report(target, "Cannot use nullsafe operator in write context")
			return
		case *ast.ExprArrayDimFetch:
			current = fetch.Var
		default:
			switch current.(type) {
			case *ast.ExprConstFetch, *ast.ExprClassConstFetch:
				if sawProperty && v.version >= PHP86 {
					return
				}
				v.report(target, "Cannot use temporary expression in write context")
			}
			return
		}
	}
}
