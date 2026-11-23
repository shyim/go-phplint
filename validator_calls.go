package phplint

import (
	"github.com/shyim/go-phplint/internal/ast"
	"github.com/shyim/go-phplint/internal/token"
)

// callKind selects the compile-time treatment of a call argument list.
// PHP forbids Closure creation (first-class callables and PHP 8.6 partial
// application) in new expressions, nullsafe calls, and attribute arguments.
type callKind uint8

const (
	callNormal callKind = iota
	callNew
	callNullsafe
	callAttribute
)

// validateCallArguments mirrors the compile-time argument checks PHP applies
// to every call-like construct: named/positional ordering, and the PHP 8.6
// partial function application placeholder rules.
func (v *compileValidator) validateCallArguments(
	node ast.Vertex,
	args []ast.Vertex,
	ellipsis *token.Token,
	kind callKind,
) {
	usesNamed := false
	usesVariadicPlaceholder := false
	hasPlaceholder := false
	hasUnpack := false

	for _, argNode := range args {
		argument, ok := argNode.(*ast.Argument)
		if !ok {
			continue
		}

		named := argument.Name != nil
		placeholder := argument.PlaceholderTkn != nil
		variadicPlaceholder := placeholder && argument.PlaceholderTkn.ID == token.T_ELLIPSIS
		unpack := argument.VariadicTkn != nil && argument.Expr != nil

		if placeholder {
			hasPlaceholder = true
		}
		if hasUnpack && !unpack && !named && !placeholder {
			v.report(argument, "Cannot use positional argument after argument unpacking")
		}
		if unpack {
			if kind == callAttribute {
				v.report(argument, "Cannot use unpacking in attribute argument list")
			}
			hasUnpack = true
		}

		// A variadic placeholder stands for the remaining positional
		// arguments, so it may follow named arguments; every other
		// positional argument may not.
		if usesNamed && !named && !variadicPlaceholder {
			v.report(argument, "Cannot use positional argument after named argument")
		}
		if usesVariadicPlaceholder {
			if variadicPlaceholder {
				v.report(argument, "Variadic placeholder may only appear once")
			} else {
				v.report(argument, "Variadic placeholder must be last")
			}
		}

		if named {
			usesNamed = true
		}
		if variadicPlaceholder {
			usesVariadicPlaceholder = true
		}
	}

	if hasPlaceholder && hasUnpack {
		v.report(node, "Cannot combine partial application and unpacking")
	}
	if ellipsis == nil && !hasPlaceholder {
		return
	}
	switch kind {
	case callNew:
		v.report(node, "Cannot create Closure for new expression")
	case callAttribute:
		v.report(node, "Cannot create Closure as attribute argument")
	case callNullsafe:
		v.report(node, "Cannot combine nullsafe operator with Closure creation")
	case callNormal:
		if method, ok := node.(*ast.ExprMethodCall); ok && hasNullsafeReceiver(method.Var) {
			v.report(node, "Cannot combine nullsafe operator with Closure creation")
		}
	}
}

// Only the receiver chain propagates a nullsafe access. Arguments, index
// expressions and dynamic function calls start independent chains.
func hasNullsafeReceiver(node ast.Vertex) bool {
	for {
		switch n := node.(type) {
		case *ast.ExprNullsafePropertyFetch, *ast.ExprNullsafeMethodCall:
			return true
		case *ast.ExprPropertyFetch:
			node = n.Var
		case *ast.ExprMethodCall:
			node = n.Var
		case *ast.ExprArrayDimFetch:
			node = n.Var
		case *ast.ExprBrackets:
			node = n.Expr
		case *ast.ExprStaticPropertyFetch:
			node = n.Class
		case *ast.ExprStaticCall:
			node = n.Class
		default:
			return false
		}
	}
}
