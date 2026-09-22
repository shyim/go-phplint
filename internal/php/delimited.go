package php

import (
	"github.com/shyim/go-phplint/internal/ast"
	"github.com/shyim/go-phplint/internal/token"
)

// delimited is a parser-stack value for a comma-separated list. It is stored
// by value in the yacc stack so the list is not wrapped in a heap object that
// the next reduction immediately unpacks.
type delimited struct {
	items []ast.Vertex
	seps  []*token.Token
	open  *token.Token
	close *token.Token
	extra *token.Token
}

func singleItem(node ast.Vertex) delimited {
	if node == nil {
		return delimited{}
	}
	return delimited{items: []ast.Vertex{node}}
}

func appendItem(list delimited, sep *token.Token, node ast.Vertex) delimited {
	if node != nil {
		list.items = append(list.items, node)
	}
	if sep != nil {
		list.seps = append(list.seps, sep)
	}
	return list
}
