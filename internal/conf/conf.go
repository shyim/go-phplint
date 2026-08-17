package conf

import (
	"github.com/shyim/go-phplint/internal/errors"
	"github.com/shyim/go-phplint/internal/version"
)

type Config struct {
	Version          *version.Version
	ErrorHandlerFunc func(e *errors.Error)
}
