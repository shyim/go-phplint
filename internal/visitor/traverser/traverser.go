package traverser

import (
	"github.com/shyim/go-phplint/internal/ast"
)

//go:generate go run traverser_gen.go

type CheckEntrance interface {
	EnterNode(ast.Vertex) bool
}

type NotifyLeave interface {
	LeaveNode(ast.Vertex)
}

type Traverser struct {
	v ast.Visitor
}

func NewTraverser(v ast.Visitor) *Traverser {
	return &Traverser{
		v: v,
	}
}

func (t *Traverser) Traverse(n ast.Vertex) {
	if n != nil {
		n.Accept(t)
	}
}

func (t *Traverser) checkEntrance(n ast.Vertex) bool {
	if ssv, ok := t.v.(CheckEntrance); ok {
		return ssv.EnterNode(n)
	}

	return true
}

func (t *Traverser) leave(n ast.Vertex) {
	if ssv, ok := t.v.(NotifyLeave); ok {
		ssv.LeaveNode(n)
	}
}
