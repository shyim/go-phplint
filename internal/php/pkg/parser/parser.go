package parser

import (
	"errors"

	"github.com/shyim/go-phplint/internal/php/internal/php7"
	"github.com/shyim/go-phplint/internal/php/internal/php8"
	"github.com/shyim/go-phplint/internal/php/pkg/ast"
	"github.com/shyim/go-phplint/internal/php/pkg/conf"
	"github.com/shyim/go-phplint/internal/php/pkg/version"
)

// ErrVersionOutOfRange is returned if the version is not supported
var ErrVersionOutOfRange = errors.New("the version is out of supported range")

// Parser interface
type Parser interface {
	Parse() int
	GetRootNode() ast.Vertex
}

func Parse(src []byte, config conf.Config) (ast.Vertex, error) {
	var parser Parser

	if config.Version == nil {
		config.Version = &version.Version{Major: 7, Minor: 4}
	}

	if config.Version.InPhp7Range() {
		lexer := php7.NewLexer(src, config)
		parser = php7.NewParser(lexer, config)
		parser.Parse()
		return parser.GetRootNode(), nil
	}

	if config.Version.InPhp8Range() {
		lexer := php8.NewLexer(src, config)
		parser = php8.NewParser(lexer, config)
		parser.Parse()
		return parser.GetRootNode(), nil
	}

	return nil, ErrVersionOutOfRange
}
