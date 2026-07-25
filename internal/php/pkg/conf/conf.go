package conf

import (
	"github.com/shyim/go-phplint/internal/php/pkg/errors"
	"github.com/shyim/go-phplint/internal/php/pkg/version"
)

type Config struct {
	Version          *version.Version
	ErrorHandlerFunc func(e *errors.Error)
}
