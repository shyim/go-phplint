package conf

import (
	"github.com/shyim/go-phplint/internal/errors"
	"github.com/shyim/go-phplint/internal/token"
	"github.com/shyim/go-phplint/internal/version"
)

type Config struct {
	Version          *version.Version
	ErrorHandlerFunc func(e *errors.Error)
	// Fidelity keeps whitespace and comment tokens and splits qualified names
	// into per-segment tokens. Lint leaves this false.
	Fidelity bool
	// Tokens, when non-nil, receives every token the parser shifts.
	Tokens *[]*token.Token
}
