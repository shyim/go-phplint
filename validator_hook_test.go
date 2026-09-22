package phplint

import (
	"testing"

	"github.com/shyim/go-phplint/internal/ast"
	"github.com/shyim/go-phplint/internal/token"
)

func TestVerifyHookedProperty(t *testing.T) {
	t.Parallel()

	thisX := thisProperty("x")
	thisXCase := thisProperty("X")
	thisNullsafe := &ast.ExprNullsafePropertyFetch{
		Var:  thisVariable(),
		Prop: &ast.Identifier{Value: []byte("x")},
	}
	closureUse := &ast.ExprClosure{
		Stmts: []ast.Vertex{&ast.StmtReturn{Expr: thisX}},
	}
	arrowUse := &ast.ExprArrowFunction{Expr: thisX}

	tests := []struct {
		name       string
		kind       string
		modifiers  []string
		hooks      []ast.PropertyHook
		hasDefault bool
		want       string
	}{
		{
			name:      "public(set) get-only",
			modifiers: []string{"public(set)"},
			hooks:     []ast.PropertyHook{{Name: "get", Kind: ast.PropertyHookShort, Body: &ast.ScalarLnumber{Value: []byte("1")}}},
		},
		{
			name:      "public public(set) get-only",
			modifiers: []string{"public", "public(set)"},
			hooks:     []ast.PropertyHook{{Name: "get", Kind: ast.PropertyHookShort, Body: &ast.ScalarLnumber{Value: []byte("1")}}},
		},
		{
			name:      "private private(set) get-only",
			modifiers: []string{"private", "private(set)"},
			hooks:     []ast.PropertyHook{{Name: "get", Kind: ast.PropertyHookShort, Body: &ast.ScalarLnumber{Value: []byte("1")}}},
		},
		{
			name:      "get short uses property",
			modifiers: []string{"public", "private(set)"},
			hooks:     []ast.PropertyHook{{Name: "get", Kind: ast.PropertyHookShort, Body: thisX}},
		},
		{
			name:      "get block uses property",
			modifiers: []string{"public", "private(set)"},
			hooks: []ast.PropertyHook{{
				Name:  "get",
				Kind:  ast.PropertyHookBlock,
				Stmts: []ast.Vertex{&ast.StmtReturn{Expr: thisX}},
			}},
		},
		{
			name:      "nullsafe use",
			modifiers: []string{"public", "private(set)"},
			hooks:     []ast.PropertyHook{{Name: "get", Kind: ast.PropertyHookShort, Body: thisNullsafe}},
		},
		{
			name:      "wrong case is still virtual",
			modifiers: []string{"public", "private(set)"},
			hooks:     []ast.PropertyHook{{Name: "get", Kind: ast.PropertyHookShort, Body: thisXCase}},
			want:      "read-only virtual property",
		},
		{
			name:      "closure use is still virtual",
			modifiers: []string{"public", "private(set)"},
			hooks:     []ast.PropertyHook{{Name: "get", Kind: ast.PropertyHookShort, Body: closureUse}},
			want:      "read-only virtual property",
		},
		{
			name:      "arrow use is still virtual",
			modifiers: []string{"public", "private(set)"},
			hooks:     []ast.PropertyHook{{Name: "get", Kind: ast.PropertyHookShort, Body: arrowUse}},
			want:      "read-only virtual property",
		},
		{
			name:      "write-only virtual set block",
			modifiers: []string{"public", "private(set)"},
			hooks: []ast.PropertyHook{{
				Name: "set",
				Kind: ast.PropertyHookBlock,
				Stmts: []ast.Vertex{&ast.StmtExpression{Expr: &ast.ExprAssign{
					Var:  &ast.ExprVariable{Name: &ast.Identifier{Value: []byte("$foo")}},
					Expr: &ast.ExprVariable{Name: &ast.Identifier{Value: []byte("$value")}},
				}}},
			}},
			want: "write-only virtual property",
		},
		{
			name:      "short set is backed",
			modifiers: []string{"public", "private(set)"},
			hooks:     []ast.PropertyHook{{Name: "set", Kind: ast.PropertyHookShort, Body: &ast.ExprVariable{Name: &ast.Identifier{Value: []byte("$value")}}}},
		},
		{
			name:      "set block writes property",
			modifiers: []string{"public", "private(set)"},
			hooks: []ast.PropertyHook{{
				Name: "set",
				Kind: ast.PropertyHookBlock,
				Stmts: []ast.Vertex{&ast.StmtExpression{Expr: &ast.ExprAssign{
					Var:  thisX,
					Expr: &ast.ExprVariable{Name: &ast.Identifier{Value: []byte("$value")}},
				}}},
			}},
		},
		{
			name:       "virtual default",
			modifiers:  []string{"public", "private(set)"},
			hasDefault: true,
			hooks:      []ast.PropertyHook{{Name: "get", Kind: ast.PropertyHookShort, Body: &ast.ScalarLnumber{Value: []byte("1")}}},
			want:       "default value for virtual hooked property",
		},
		{
			name:       "backed default",
			modifiers:  []string{"public", "private(set)"},
			hasDefault: true,
			hooks:      []ast.PropertyHook{{Name: "get", Kind: ast.PropertyHookShort, Body: thisX}},
		},
		{
			name:      "by-ref get on backed property",
			modifiers: []string{"public"},
			hooks: []ast.PropertyHook{
				{Name: "get", Kind: ast.PropertyHookShort, ByRef: true, Body: thisX},
				{Name: "set", Kind: ast.PropertyHookShort, Body: &ast.ExprVariable{Name: &ast.Identifier{Value: []byte("$value")}}},
			},
			want: "may not return by reference",
		},
		{
			name:      "by-ref get on virtual property",
			modifiers: []string{"public"},
			hooks: []ast.PropertyHook{
				{Name: "get", Kind: ast.PropertyHookShort, ByRef: true, Body: &ast.ScalarLnumber{Value: []byte("1")}},
				{
					Name: "set",
					Kind: ast.PropertyHookBlock,
					Stmts: []ast.Vertex{&ast.StmtExpression{Expr: &ast.ExprAssign{
						Var:  &ast.ExprVariable{Name: &ast.Identifier{Value: []byte("$foo")}},
						Expr: &ast.ExprVariable{Name: &ast.Identifier{Value: []byte("$value")}},
					}}},
				},
			},
		},
		{
			name:      "get-only asymmetric",
			modifiers: []string{"public", "private(set)"},
			hooks:     []ast.PropertyHook{{Name: "get", Kind: ast.PropertyHookShort, Body: &ast.ScalarLnumber{Value: []byte("1")}}},
			want:      "read-only virtual property",
		},
		{
			name:      "abstract semicolon get",
			modifiers: []string{"abstract", "public", "private(set)"},
			hooks:     []ast.PropertyHook{{Name: "get", Kind: ast.PropertyHookSemi}},
			want:      "read-only virtual property",
		},
		{
			name:      "abstract by-ref semicolon get",
			modifiers: []string{"abstract", "public", "private(set)"},
			hooks:     []ast.PropertyHook{{Name: "get", Kind: ast.PropertyHookSemi, ByRef: true}},
			want:      "read-only virtual property",
		},
		{
			name:      "nested function use is still virtual",
			modifiers: []string{"public", "private(set)"},
			hooks: []ast.PropertyHook{{
				Name: "get",
				Kind: ast.PropertyHookBlock,
				Stmts: []ast.Vertex{
					&ast.StmtFunction{
						Stmts: []ast.Vertex{&ast.StmtReturn{Expr: thisX}},
					},
					&ast.StmtReturn{Expr: &ast.ScalarLnumber{Value: []byte("1")}},
				},
			}},
			want: "read-only virtual property",
		},
		{
			name:      "parenthesized this",
			modifiers: []string{"public", "private(set)"},
			hooks: []ast.PropertyHook{{
				Name: "get",
				Kind: ast.PropertyHookShort,
				Body: &ast.ExprPropertyFetch{
					Var:  &ast.ExprBrackets{Expr: thisVariable()},
					Prop: &ast.Identifier{Value: []byte("x")},
				},
			}},
		},
		{
			name:      "constant computed name",
			modifiers: []string{"public", "private(set)"},
			hooks: []ast.PropertyHook{{
				Name: "get",
				Kind: ast.PropertyHookShort,
				Body: &ast.ExprPropertyFetch{
					Var:                 thisVariable(),
					OpenCurlyBracketTkn: &token.Token{Value: []byte("{")},
					Prop:                &ast.ScalarString{Value: []byte("\"x\"")},
				},
			}},
		},
		{
			name:      "parser shaped constant computed name",
			modifiers: []string{"public", "private(set)"},
			hooks: []ast.PropertyHook{{
				Name: "get",
				Kind: ast.PropertyHookShort,
				Body: &ast.ExprPropertyFetch{
					Var: thisVariable(),
					Prop: &ast.ExprBrackets{
						Expr: &ast.ScalarString{Value: []byte("\"x\"")},
					},
				},
			}},
		},
		{
			name:      "parser shaped single quoted computed name",
			modifiers: []string{"public", "private(set)"},
			hooks: []ast.PropertyHook{{
				Name: "get",
				Kind: ast.PropertyHookShort,
				Body: &ast.ExprNullsafePropertyFetch{
					Var: thisVariable(),
					Prop: &ast.ExprBrackets{
						Expr: &ast.ScalarString{Value: []byte("'x'")},
					},
				},
			}},
		},
		{
			name:      "variable computed name is still virtual",
			modifiers: []string{"public", "private(set)"},
			hooks: []ast.PropertyHook{{
				Name: "get",
				Kind: ast.PropertyHookShort,
				Body: &ast.ExprPropertyFetch{
					Var:                 thisVariable(),
					OpenCurlyBracketTkn: &token.Token{Value: []byte("{")},
					Prop:                &ast.ExprVariable{Name: &ast.Identifier{Value: []byte("$n")}},
				},
			}},
			want: "read-only virtual property",
		},
		{
			name:      "parser shaped variable computed name is still virtual",
			modifiers: []string{"public", "private(set)"},
			hooks: []ast.PropertyHook{{
				Name: "get",
				Kind: ast.PropertyHookShort,
				Body: &ast.ExprPropertyFetch{
					Var: thisVariable(),
					Prop: &ast.ExprBrackets{
						Expr: &ast.ExprVariable{Name: &ast.Identifier{Value: []byte("$n")}},
					},
				},
			}},
			want: "read-only virtual property",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			property := hookedProperty(test.modifiers, test.hooks, test.hasDefault)
			validator := &compileValidator{version: PHP84}
			validator.verifyHookedProperty(property, "class")
			if test.want == "" {
				if len(validator.diagnostics) != 0 {
					t.Fatalf("diagnostics = %#v, want none", validator.diagnostics)
				}
				return
			}
			if !diagnosticsContain(validator.diagnostics, test.want) {
				t.Fatalf("diagnostics = %#v, want phrase %q", validator.diagnostics, test.want)
			}
		})
	}

	if !interfacePropertyVisibilityForbidden(modifierNodes("protected")) ||
		!interfacePropertyVisibilityForbidden(modifierNodes("private")) {
		t.Fatal("plain interface visibility was accepted")
	}
	if interfacePropertyVisibilityForbidden(modifierNodes("protected(set)")) ||
		interfacePropertyVisibilityForbidden(modifierNodes("private(set)")) ||
		interfacePropertyVisibilityForbidden(modifierNodes("public", "protected(set)")) {
		t.Fatal("asymmetric interface set visibility was rejected")
	}
}

func thisVariable() ast.Vertex {
	return &ast.ExprVariable{Name: &ast.Identifier{Value: []byte("$this")}}
}

func thisProperty(name string) ast.Vertex {
	return &ast.ExprPropertyFetch{
		Var:  thisVariable(),
		Prop: &ast.Identifier{Value: []byte(name)},
	}
}

func modifierNodes(names ...string) []ast.Vertex {
	nodes := make([]ast.Vertex, len(names))
	for index, name := range names {
		nodes[index] = &ast.Identifier{Value: []byte(name)}
	}
	return nodes
}

func hookedProperty(modifiers []string, hooks []ast.PropertyHook, hasDefault bool) *ast.StmtPropertyList {
	var expr ast.Vertex
	if hasDefault {
		expr = &ast.ScalarString{Value: []byte("a")}
	}
	return &ast.StmtPropertyList{
		Hooked:    true,
		Modifiers: modifierNodes(modifiers...),
		Hooks:     hooks,
		Props: []ast.Vertex{&ast.StmtProperty{
			Var:  &ast.ExprVariable{Name: &ast.Identifier{Value: []byte("$x")}},
			Expr: expr,
		}},
	}
}
