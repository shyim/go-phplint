package dumper_test

import (
	"bytes"
	"testing"

	"github.com/shyim/go-phplint/internal/position"
	"github.com/shyim/go-phplint/internal/token"
	"github.com/shyim/go-phplint/internal/visitor/dumper"

	"github.com/shyim/go-phplint/internal/ast"
)

func TestDumper_root(t *testing.T) {
	o := bytes.NewBufferString("")

	p := dumper.NewDumper(o).WithTokens().WithPositions()
	n := &ast.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   2,
			StartPos:  3,
			EndPos:    4,
		},
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
		EndTkn: &token.Token{
			FreeFloating: []*token.Token{
				{
					ID:    token.T_WHITESPACE,
					Value: []byte(" "),
					Position: &position.Position{
						StartLine: 1,
						EndLine:   2,
						StartPos:  3,
						EndPos:    4,
					},
				},
			},
		},
	}
	n.Accept(p)

	expected := `&ast.Root{
	Position: &position.Position{
		StartLine: 1,
		EndLine:   2,
		StartCol:  0,
		EndCol:    0,
		StartPos:  3,
		EndPos:    4,
	},
	Stmts: []ast.Vertex{
		&ast.StmtNop{
		},
	},
	EndTkn: &token.Token{
		FreeFloating: []*token.Token{
			{
				ID: token.T_WHITESPACE,
				Val: []byte(" "),
				Position: &position.Position{
					StartLine: 1,
					EndLine:   2,
					StartCol:  0,
					EndCol:    0,
					StartPos:  3,
					EndPos:    4,
				},
			},
		},
	},
},
`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}
