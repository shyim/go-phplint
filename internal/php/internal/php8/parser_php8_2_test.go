package php8_test

import (
	"testing"

	"github.com/shyim/phplint-go/internal/php/internal/tester"
)

func TestClassReadonlyModifier(t *testing.T) {
	suite := tester.NewParserDumpTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<?php
readonly class Foo {
	public string $a;
}

final readonly class Boo {
	public string $a;
}
`

	suite.Expected = `&ast.Root{
	Stmts: []ast.Vertex{
		&ast.StmtClass{
			Modifiers: []ast.Vertex{
				&ast.Identifier{
					Val: []byte("readonly"),
				},
			},
			Name: &ast.Identifier{
				Val: []byte("Foo"),
			},
			Stmts: []ast.Vertex{
				&ast.StmtPropertyList{
					Modifiers: []ast.Vertex{
						&ast.Identifier{
							Val: []byte("public"),
						},
					},
					Type: &ast.Name{
						Parts: []ast.Vertex{
							&ast.NamePart{
								Val: []byte("string"),
							},
						},
					},
					Props: []ast.Vertex{
						&ast.StmtProperty{
							Var: &ast.ExprVariable{
								Name: &ast.Identifier{
									Val: []byte("$a"),
								},
							},
						},
					},
				},
			},
		},
		&ast.StmtClass{
			Modifiers: []ast.Vertex{
				&ast.Identifier{
					Val: []byte("final"),
				},
				&ast.Identifier{
					Val: []byte("readonly"),
				},
			},
			Name: &ast.Identifier{
				Val: []byte("Boo"),
			},
			Stmts: []ast.Vertex{
				&ast.StmtPropertyList{
					Modifiers: []ast.Vertex{
						&ast.Identifier{
							Val: []byte("public"),
						},
					},
					Type: &ast.Name{
						Parts: []ast.Vertex{
							&ast.NamePart{
								Val: []byte("string"),
							},
						},
					},
					Props: []ast.Vertex{
						&ast.StmtProperty{
							Var: &ast.ExprVariable{
								Name: &ast.Identifier{
									Val: []byte("$a"),
								},
							},
						},
					},
				},
			},
		},
	},
},`

	suite.Run()
}

func TestNullFalseTrueTypes(t *testing.T) {
	suite := tester.NewParserDumpTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<?php
function alwaysFalse(false $a): false {}

function alwaysTrue(true $a): true {}

function alwaysNull(null $a): null {}

class Bar {
    public null $nil;
    public true $true;
    public false $false;

    public function nil(null $nil): null {}
    public function true(true $true): true {}
    public function false(false $false): false {}
}
`

	suite.Expected = `&ast.Root{
	Stmts: []ast.Vertex{
		&ast.StmtFunction{
			Name: &ast.Identifier{
				Val: []byte("alwaysFalse"),
			},
			Params: []ast.Vertex{
				&ast.Parameter{
					Type: &ast.Name{
						Parts: []ast.Vertex{
							&ast.NamePart{
								Val: []byte("false"),
							},
						},
					},
					Var: &ast.ExprVariable{
						Name: &ast.Identifier{
							Val: []byte("$a"),
						},
					},
				},
			},
			ReturnType: &ast.Name{
				Parts: []ast.Vertex{
					&ast.NamePart{
						Val: []byte("false"),
					},
				},
			},
			Stmts: []ast.Vertex{},
		},
		&ast.StmtFunction{
			Name: &ast.Identifier{
				Val: []byte("alwaysTrue"),
			},
			Params: []ast.Vertex{
				&ast.Parameter{
					Type: &ast.Name{
						Parts: []ast.Vertex{
							&ast.NamePart{
								Val: []byte("true"),
							},
						},
					},
					Var: &ast.ExprVariable{
						Name: &ast.Identifier{
							Val: []byte("$a"),
						},
					},
				},
			},
			ReturnType: &ast.Name{
				Parts: []ast.Vertex{
					&ast.NamePart{
						Val: []byte("true"),
					},
				},
			},
			Stmts: []ast.Vertex{},
		},
		&ast.StmtFunction{
			Name: &ast.Identifier{
				Val: []byte("alwaysNull"),
			},
			Params: []ast.Vertex{
				&ast.Parameter{
					Type: &ast.Name{
						Parts: []ast.Vertex{
							&ast.NamePart{
								Val: []byte("null"),
							},
						},
					},
					Var: &ast.ExprVariable{
						Name: &ast.Identifier{
							Val: []byte("$a"),
						},
					},
				},
			},
			ReturnType: &ast.Name{
				Parts: []ast.Vertex{
					&ast.NamePart{
						Val: []byte("null"),
					},
				},
			},
			Stmts: []ast.Vertex{},
		},
		&ast.StmtClass{
			Name: &ast.Identifier{
				Val: []byte("Bar"),
			},
			Stmts: []ast.Vertex{
				&ast.StmtPropertyList{
					Modifiers: []ast.Vertex{
						&ast.Identifier{
							Val: []byte("public"),
						},
					},
					Type: &ast.Name{
						Parts: []ast.Vertex{
							&ast.NamePart{
								Val: []byte("null"),
							},
						},
					},
					Props: []ast.Vertex{
						&ast.StmtProperty{
							Var: &ast.ExprVariable{
								Name: &ast.Identifier{
									Val: []byte("$nil"),
								},
							},
						},
					},
				},
				&ast.StmtPropertyList{
					Modifiers: []ast.Vertex{
						&ast.Identifier{
							Val: []byte("public"),
						},
					},
					Type: &ast.Name{
						Parts: []ast.Vertex{
							&ast.NamePart{
								Val: []byte("true"),
							},
						},
					},
					Props: []ast.Vertex{
						&ast.StmtProperty{
							Var: &ast.ExprVariable{
								Name: &ast.Identifier{
									Val: []byte("$true"),
								},
							},
						},
					},
				},
				&ast.StmtPropertyList{
					Modifiers: []ast.Vertex{
						&ast.Identifier{
							Val: []byte("public"),
						},
					},
					Type: &ast.Name{
						Parts: []ast.Vertex{
							&ast.NamePart{
								Val: []byte("false"),
							},
						},
					},
					Props: []ast.Vertex{
						&ast.StmtProperty{
							Var: &ast.ExprVariable{
								Name: &ast.Identifier{
									Val: []byte("$false"),
								},
							},
						},
					},
				},
				&ast.StmtClassMethod{
					Modifiers: []ast.Vertex{
						&ast.Identifier{
							Val: []byte("public"),
						},
					},
					Name: &ast.Identifier{
						Val: []byte("nil"),
					},
					Params: []ast.Vertex{
						&ast.Parameter{
							Type: &ast.Name{
								Parts: []ast.Vertex{
									&ast.NamePart{
										Val: []byte("null"),
									},
								},
							},
							Var: &ast.ExprVariable{
								Name: &ast.Identifier{
									Val: []byte("$nil"),
								},
							},
						},
					},
					ReturnType: &ast.Name{
						Parts: []ast.Vertex{
							&ast.NamePart{
								Val: []byte("null"),
							},
						},
					},
					Stmt: &ast.StmtStmtList{
						Stmts: []ast.Vertex{},
					},
				},
				&ast.StmtClassMethod{
					Modifiers: []ast.Vertex{
						&ast.Identifier{
							Val: []byte("public"),
						},
					},
					Name: &ast.Identifier{
						Val: []byte("true"),
					},
					Params: []ast.Vertex{
						&ast.Parameter{
							Type: &ast.Name{
								Parts: []ast.Vertex{
									&ast.NamePart{
										Val: []byte("true"),
									},
								},
							},
							Var: &ast.ExprVariable{
								Name: &ast.Identifier{
									Val: []byte("$true"),
								},
							},
						},
					},
					ReturnType: &ast.Name{
						Parts: []ast.Vertex{
							&ast.NamePart{
								Val: []byte("true"),
							},
						},
					},
					Stmt: &ast.StmtStmtList{
						Stmts: []ast.Vertex{},
					},
				},
				&ast.StmtClassMethod{
					Modifiers: []ast.Vertex{
						&ast.Identifier{
							Val: []byte("public"),
						},
					},
					Name: &ast.Identifier{
						Val: []byte("false"),
					},
					Params: []ast.Vertex{
						&ast.Parameter{
							Type: &ast.Name{
								Parts: []ast.Vertex{
									&ast.NamePart{
										Val: []byte("false"),
									},
								},
							},
							Var: &ast.ExprVariable{
								Name: &ast.Identifier{
									Val: []byte("$false"),
								},
							},
						},
					},
					ReturnType: &ast.Name{
						Parts: []ast.Vertex{
							&ast.NamePart{
								Val: []byte("false"),
							},
						},
					},
					Stmt: &ast.StmtStmtList{
						Stmts: []ast.Vertex{},
					},
				},
			},
		},
	},
},`

	suite.Run()
}

func TestConstantsInTraits(t *testing.T) {
	suite := tester.NewParserDumpTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<?php
trait Foo {
    public const CONSTANT = 1;
}
`

	suite.Expected = `&ast.Root{
	Stmts: []ast.Vertex{
		&ast.StmtTrait{
			Name: &ast.Identifier{
				Val: []byte("Foo"),
			},
			Stmts: []ast.Vertex{
				&ast.StmtClassConstList{
					Modifiers: []ast.Vertex{
						&ast.Identifier{
							Val: []byte("public"),
						},
					},
					Consts: []ast.Vertex{
						&ast.StmtConstant{
							Name: &ast.Identifier{
								Val: []byte("CONSTANT"),
							},
							Expr: &ast.ScalarLnumber{
								Val: []byte("1"),
							},
						},
					},
				},
			},
		},
	},
},`

	suite.Run()
}
