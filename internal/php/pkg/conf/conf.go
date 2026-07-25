package conf

import (
	"github.com/shyim/phplint-go/internal/php/pkg/errors"
	"github.com/shyim/phplint-go/internal/php/pkg/version"
)

type Config struct {
	Version          *version.Version
	ErrorHandlerFunc func(e *errors.Error)
}
