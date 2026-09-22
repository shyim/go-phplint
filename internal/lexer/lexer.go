package lexer

import (
	"errors"

	php "github.com/shyim/go-phplint/internal/php"
	"github.com/shyim/go-phplint/internal/conf"
	"github.com/shyim/go-phplint/internal/token"
)

var ErrVersionOutOfRange = errors.New("the version is out of supported range")

type Lexer interface {
	Lex() *token.Token
}

func New(src []byte, config conf.Config) (Lexer, error) {
	// The unified lexer tokenizes the full PHP 7/8 superset and downgrades
	// version-gated keywords for older profiles.
	if config.Version.InPhp7Range() || config.Version.InPhp8Range() {
		return php.NewLexer(src, config), nil
	}

	return nil, ErrVersionOutOfRange
}
