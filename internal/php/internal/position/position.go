package position

import (
	"github.com/shyim/phplint-go/internal/php/pkg/ast"
	"github.com/shyim/phplint-go/internal/php/pkg/position"
	"github.com/shyim/phplint-go/internal/php/pkg/token"
)

type startPos struct {
	startLine int
	startPos  int
	startCol  int
}

type endPos struct {
	endLine int
	endPos  int
	endCol  int
}

type Builder struct {
	pool *position.Pool
}

func NewBuilder() *Builder {
	return &Builder{
		pool: position.NewPool(position.DefaultBlockSize),
	}
}

func getListStartPos(l []ast.Vertex) startPos {
	if l == nil {
		return startPos{-1, -1, -1}
	}

	if len(l) == 0 {
		return startPos{-1, -1, -1}
	}

	return getNodeStartPos(l[0])
}

func getNodeStartPos(n ast.Vertex) startPos {
	sl := -1
	sp := -1
	sc := -1

	if n == nil {
		return startPos{-1, -1, -1}
	}

	p := n.GetPosition()
	if p != nil {
		sl = p.StartLine
		sp = p.StartPos
		sc = p.StartCol
	}

	return startPos{sl, sp, sc}
}

func getListEndPos(l []ast.Vertex) endPos {
	if l == nil {
		return endPos{-1, -1, -1}
	}

	if len(l) == 0 {
		return endPos{-1, -1, -1}
	}

	return getNodeEndPos(l[len(l)-1])
}

func getNodeEndPos(n ast.Vertex) endPos {
	el := -1
	ep := -1
	ec := -1

	if n == nil {
		return endPos{-1, -1, -1}
	}

	p := n.GetPosition()
	if p != nil {
		el = p.EndLine
		ep = p.EndPos
		ec = p.EndCol
	}

	return endPos{el, ep, ec}
}

// NewNodeListPosition returns new Position
func (b *Builder) NewNodeListPosition(list []ast.Vertex) *position.Position {
	pos := b.pool.Get()

	pos.StartLine = getListStartPos(list).startLine
	pos.EndLine = getListEndPos(list).endLine
	pos.StartPos = getListStartPos(list).startPos
	pos.EndPos = getListEndPos(list).endPos
	pos.StartCol = getListStartPos(list).startCol
	pos.EndCol = getListEndPos(list).endCol

	return pos
}

// NewNodePosition returns new Position
func (b *Builder) NewNodePosition(n ast.Vertex) *position.Position {
	pos := b.pool.Get()

	pos.StartLine = getNodeStartPos(n).startLine
	pos.EndLine = getNodeEndPos(n).endLine
	pos.StartPos = getNodeStartPos(n).startPos
	pos.EndPos = getNodeEndPos(n).endPos
	pos.StartCol = getNodeStartPos(n).startCol
	pos.EndCol = getNodeEndPos(n).endCol

	return pos
}

// NewTokenPosition returns new Position
func (b *Builder) NewTokenPosition(t *token.Token) *position.Position {
	pos := b.pool.Get()

	pos.StartLine = t.Position.StartLine
	pos.EndLine = t.Position.EndLine
	pos.StartPos = t.Position.StartPos
	pos.EndPos = t.Position.EndPos
	pos.StartCol = t.Position.StartCol
	pos.EndCol = t.Position.EndCol

	return pos
}

// NewTokensPosition returns new Position
func (b *Builder) NewTokensPosition(
	startToken *token.Token,
	endToken *token.Token,
) *position.Position {
	pos := b.pool.Get()

	pos.StartLine = startToken.Position.StartLine
	pos.EndLine = endToken.Position.EndLine
	pos.StartPos = startToken.Position.StartPos
	pos.EndPos = endToken.Position.EndPos
	pos.StartCol = startToken.Position.StartCol
	pos.EndCol = endToken.Position.EndCol

	return pos
}

// NewTokenNodePosition returns new Position
func (b *Builder) NewTokenNodePosition(t *token.Token, n ast.Vertex) *position.Position {
	pos := b.pool.Get()

	pos.StartLine = t.Position.StartLine
	pos.EndLine = getNodeEndPos(n).endLine
	pos.StartPos = t.Position.StartPos
	pos.EndPos = getNodeEndPos(n).endPos
	pos.StartCol = t.Position.StartCol
	pos.EndCol = getNodeEndPos(n).endCol

	return pos
}

// NewNodeTokenPosition returns new Position
func (b *Builder) NewNodeTokenPosition(n ast.Vertex, t *token.Token) *position.Position {
	pos := b.pool.Get()

	pos.StartLine = getNodeStartPos(n).startLine
	pos.EndLine = t.Position.EndLine
	pos.StartPos = getNodeStartPos(n).startPos
	pos.EndPos = t.Position.EndPos
	pos.StartCol = getNodeStartPos(n).startCol
	pos.EndCol = t.Position.EndCol

	return pos
}

// NewNodesPosition returns new Position
func (b *Builder) NewNodesPosition(startNode ast.Vertex, endNode ast.Vertex) *position.Position {
	pos := b.pool.Get()

	pos.StartLine = getNodeStartPos(startNode).startLine
	pos.EndLine = getNodeEndPos(endNode).endLine
	pos.StartPos = getNodeStartPos(startNode).startPos
	pos.EndPos = getNodeEndPos(endNode).endPos
	pos.StartCol = getNodeStartPos(startNode).startCol
	pos.EndCol = getNodeEndPos(endNode).endCol

	return pos
}

// NewNodeListTokenPosition returns new Position
func (b *Builder) NewNodeListTokenPosition(list []ast.Vertex, t *token.Token) *position.Position {
	pos := b.pool.Get()

	pos.StartLine = getListStartPos(list).startLine
	pos.EndLine = t.Position.EndLine
	pos.StartPos = getListStartPos(list).startPos
	pos.EndPos = t.Position.EndPos
	pos.StartCol = getListStartPos(list).startCol
	pos.EndCol = t.Position.EndCol

	return pos
}

// NewTokenNodeListPosition returns new Position
func (b *Builder) NewTokenNodeListPosition(t *token.Token, list []ast.Vertex) *position.Position {
	pos := b.pool.Get()

	pos.StartLine = t.Position.StartLine
	pos.EndLine = getListEndPos(list).endLine
	pos.StartPos = t.Position.StartPos
	pos.EndPos = getListEndPos(list).endPos
	pos.StartCol = t.Position.StartCol
	pos.EndCol = getListEndPos(list).endCol

	return pos
}

// NewNodeNodeListPosition returns new Position
func (b *Builder) NewNodeNodeListPosition(n ast.Vertex, list []ast.Vertex) *position.Position {
	pos := b.pool.Get()

	pos.StartLine = getNodeStartPos(n).startLine
	pos.EndLine = getListEndPos(list).endLine
	pos.StartPos = getNodeStartPos(n).startPos
	pos.EndPos = getListEndPos(list).endPos
	pos.StartCol = getNodeStartPos(n).startCol
	pos.EndCol = getListEndPos(list).endCol

	return pos
}

// NewNodeListNodePosition returns new Position
func (b *Builder) NewNodeListNodePosition(list []ast.Vertex, n ast.Vertex) *position.Position {
	pos := b.pool.Get()

	pos.StartLine = getListStartPos(list).startLine
	pos.EndLine = getNodeEndPos(n).endLine
	pos.StartPos = getListStartPos(list).startPos
	pos.EndPos = getNodeEndPos(n).endPos
	pos.StartCol = getListStartPos(list).startCol
	pos.EndCol = getNodeEndPos(n).endCol

	return pos
}

// NewOptionalListTokensPosition returns new Position
func (b *Builder) NewOptionalListTokensPosition(
	list []ast.Vertex,
	t *token.Token,
	endToken *token.Token,
) *position.Position {
	pos := b.pool.Get()

	if list == nil {
		pos.StartLine = t.Position.StartLine
		pos.EndLine = endToken.Position.EndLine
		pos.StartPos = t.Position.StartPos
		pos.EndPos = endToken.Position.EndPos
		pos.StartCol = t.Position.StartCol
		pos.EndCol = endToken.Position.EndCol

		return pos
	}
	pos.StartLine = getListStartPos(list).startLine
	pos.EndLine = endToken.Position.EndLine
	pos.StartPos = getListStartPos(list).startPos
	pos.EndPos = endToken.Position.EndPos
	pos.StartCol = getListStartPos(list).startCol
	pos.EndCol = endToken.Position.EndCol

	return pos
}
