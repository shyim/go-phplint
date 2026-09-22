package php

import (
	"github.com/shyim/go-phplint/internal/ast"
	"github.com/shyim/go-phplint/internal/conf"
	"github.com/shyim/go-phplint/internal/errors"
	"github.com/shyim/go-phplint/internal/posbuilder"
	"github.com/shyim/go-phplint/internal/token"
	"github.com/shyim/go-phplint/internal/version"
)

// Parser structure
type Parser struct {
	Lexer          *Lexer
	currentToken   *token.Token
	rootNode       ast.Vertex
	errHandlerFunc func(*errors.Error)
	builder        *Builder
	phpVersion     *version.Version
	fidelity       bool
	tokens         *[]*token.Token
}

// NewParser creates and returns new Parser
func NewParser(lexer *Lexer, config conf.Config) *Parser {
	p := &Parser{
		Lexer:          lexer,
		errHandlerFunc: config.ErrorHandlerFunc,
		phpVersion:     config.Version,
		fidelity:       config.Fidelity,
		tokens:         config.Tokens,
	}
	p.builder = NewBuilder(posbuilder.NewBuilder(), p)
	return p
}

// versionAtLeast reports whether the target language profile is at least
// the given major.minor version. A nil version is treated as the newest
// profile so the unified grammar accepts the full superset.
func (p *Parser) versionAtLeast(major, minor uint64) bool {
	return p.phpVersion.AtLeast(major, minor)
}

// allowCurlyOffset reports whether `{expr}` array/string offsets are valid.
// They were removed in PHP 8.0, so the unified grammar only accepts them
// for PHP 7 profiles.
func (p *Parser) allowCurlyOffset() bool {
	return !p.versionAtLeast(8, 0)
}

func (p *Parser) Lex(lval *yySymType) int {
	t := p.Lexer.Lex()

	p.currentToken = t
	lval.token = t
	if p.tokens != nil && t != nil && t.ID != 0 {
		*p.tokens = append(*p.tokens, t)
	}

	return int(t.ID)
}

func (p *Parser) Error(msg string) {
	if p.errHandlerFunc == nil {
		return
	}

	p.errHandlerFunc(errors.NewError(humanizeSyntax(msg), p.currentToken.Position))
}

// Parse the php Parser entrypoint
func (p *Parser) Parse() int {
	p.rootNode = nil

	return yyParse(p)
}

// GetRootNode returns root node
func (p *Parser) GetRootNode() ast.Vertex {
	return p.rootNode
}
