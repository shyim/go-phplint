package tester

import (
	"testing"

	"github.com/shyim/go-phplint/internal/conf"
	php "github.com/shyim/go-phplint/internal/php"
	"github.com/shyim/go-phplint/internal/version"
	"gotest.tools/assert"
)

type LexerTokenStringTestSuite struct {
	t *testing.T

	Code     string
	Expected []string

	Version version.Version
}

func NewLexerTokenStringTestSuite(t *testing.T) *LexerTokenStringTestSuite {
	return &LexerTokenStringTestSuite{
		t: t,
		Version: version.Version{
			Major: 7,
			Minor: 4,
		},
	}
}

func (l *LexerTokenStringTestSuite) UsePHP8() {
	l.Version = version.Version{Major: 8, Minor: 0}
}

func (l *LexerTokenStringTestSuite) Run() {
	config := conf.Config{
		Fidelity: true,
		Version:  &l.Version,
	}

	var lexer Lexer

	lexer = php.NewLexer([]byte(l.Code), config)

	for _, expected := range l.Expected {
		tkn := lexer.Lex()
		actual := string(tkn.Value)
		assert.DeepEqual(l.t, expected, actual)
	}
}
