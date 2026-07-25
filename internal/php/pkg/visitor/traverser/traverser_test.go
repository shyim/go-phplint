package traverser_test

import (
	"testing"

	"github.com/shyim/phplint-go/internal/php/pkg/ast"
	"github.com/shyim/phplint-go/internal/php/pkg/visitor"
	"github.com/shyim/phplint-go/internal/php/pkg/visitor/traverser"
)

// testCase is a traverser that does not go into any class statement or its children.
// If it does, the test fails.
type testCase struct {
	t *testing.T
	visitor.Null
	traversedFunction bool
}

var _ ast.Visitor = &testCase{}

func (t *testCase) EnterNode(n ast.Vertex) bool {
	t.t.Logf("EnterNode: %T", n)
	if _, ok := n.(*ast.StmtClass); ok {
		return false
	}

	return true
}

func (t *testCase) LeaveNode(n ast.Vertex) {
	t.t.Logf("LeaveNode: %T", n)
	if _, ok := n.(*ast.Root); ok {
		if !t.traversedFunction {
			t.t.Error("traverser did not traverse function")
		}
	}
}

func (t *testCase) StmtClass(n *ast.StmtClass) {
	t.t.Errorf("traverser got to class")
}

func (t *testCase) StmtClassMethod(n *ast.StmtClassMethod) {
	t.t.Errorf("traverser got to method")
}

func (t *testCase) StmtFunction(n *ast.StmtFunction) {
	t.traversedFunction = true
}

func TestEnterNodeIsRespected(t *testing.T) {
	tc := &testCase{t: t}
	tv := traverser.NewTraverser(tc)

	root := &ast.Root{
		Stmts: []ast.Vertex{
			&ast.StmtFunction{},
			&ast.StmtClass{
				Stmts: []ast.Vertex{
					&ast.StmtClassMethod{},
				},
			},
		},
	}

	root.Accept(tv)
}
