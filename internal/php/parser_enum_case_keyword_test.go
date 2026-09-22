package php_test

import (
	"testing"

	"github.com/shyim/go-phplint/internal/tester"
)

func TestEnumCaseSemiReservedKeywordNames(t *testing.T) {
	suite := tester.NewParserDumpTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<?php
enum Status: string {
    case NEW = 'new';
    case DEFAULT = 'default';
    case List = 'list';
    case Done = 'done';
}
`

	suite.Expected = `&ast.Root{
	Stmts: []ast.Vertex{
		&ast.StmtEnum{
			Name: &ast.Identifier{
				Val: []byte("Status"),
			},
			Type: &ast.Name{
				Parts: []ast.Vertex{
					&ast.NamePart{
						Val: []byte("string"),
					},
				},
			},
			Stmts: []ast.Vertex{
				&ast.EnumCase{
					Name: &ast.Identifier{
						Val: []byte("NEW"),
					},
					Expr: &ast.ScalarString{
						Val: []byte("'new'"),
					},
				},
				&ast.EnumCase{
					Name: &ast.Identifier{
						Val: []byte("DEFAULT"),
					},
					Expr: &ast.ScalarString{
						Val: []byte("'default'"),
					},
				},
				&ast.EnumCase{
					Name: &ast.Identifier{
						Val: []byte("List"),
					},
					Expr: &ast.ScalarString{
						Val: []byte("'list'"),
					},
				},
				&ast.EnumCase{
					Name: &ast.Identifier{
						Val: []byte("Done"),
					},
					Expr: &ast.ScalarString{
						Val: []byte("'done'"),
					},
				},
			},
		},
	},
},`

	suite.Run()
}
