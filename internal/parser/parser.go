package parser

import (
	"errors"

	php "github.com/shyim/go-phplint/internal/php"
	"github.com/shyim/go-phplint/internal/ast"
	"github.com/shyim/go-phplint/internal/conf"
	"github.com/shyim/go-phplint/internal/version"
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

	// The unified parser accepts the full PHP 7/8 syntax superset for every
	// supported profile. Version-specific rejection happens in the lexer
	// (reserved words, removed casts, heredoc rules) and in validation.
	if config.Version.InPhp7Range() || config.Version.InPhp8Range() {
		lexer := php.NewLexer(src, config)
		parser = php.NewParser(lexer, config)
		parser.Parse()
		return parser.GetRootNode(), nil
	}

	return nil, ErrVersionOutOfRange
}
