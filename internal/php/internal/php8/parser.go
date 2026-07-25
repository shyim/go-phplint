package php8

import (
	"github.com/shyim/go-phplint/internal/php/internal/position"
	"github.com/shyim/go-phplint/internal/php/pkg/ast"
	"github.com/shyim/go-phplint/internal/php/pkg/conf"
	"github.com/shyim/go-phplint/internal/php/pkg/errors"
	"github.com/shyim/go-phplint/internal/php/pkg/token"
)

// Parser structure
type Parser struct {
	Lexer          *Lexer
	currentToken   *token.Token
	rootNode       ast.Vertex
	errHandlerFunc func(*errors.Error)
	builder        *Builder
}

// NewParser creates and returns new Parser
func NewParser(lexer *Lexer, config conf.Config) *Parser {
	p := &Parser{
		Lexer:          lexer,
		errHandlerFunc: config.ErrorHandlerFunc,
	}
	p.builder = NewBuilder(position.NewBuilder(), p)
	return p
}

func (p *Parser) Lex(lval *yySymType) int {
	t := p.Lexer.Lex()

	p.currentToken = t
	lval.token = t

	return int(t.ID)
}

func (p *Parser) Error(msg string) {
	if p.errHandlerFunc == nil {
		return
	}

	p.errHandlerFunc(errors.NewError(msg, p.currentToken.Position))
}

// Parse the php7 Parser entrypoint
func (p *Parser) Parse() int {
	p.rootNode = nil

	return yyParse(p)
}

// GetRootNode returns root node
func (p *Parser) GetRootNode() ast.Vertex {
	return p.rootNode
}
