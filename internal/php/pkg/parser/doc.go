/*
A Parser for PHP written in Go

Package usage example:

	package main

	import (
		"log"
		"os"

		"github.com/shyim/phplint-go/internal/php/pkg/conf"
		"github.com/shyim/phplint-go/internal/php/pkg/errors"
		"github.com/shyim/phplint-go/internal/php/pkg/parser"
		"github.com/shyim/phplint-go/internal/php/pkg/version"
		"github.com/shyim/phplint-go/internal/php/pkg/visitor/dumper"
	)

	func main() {
		src := []byte(`<? echo "Hello world";`)

		// Error handler

		var parserErrors []*errors.Error
		errorHandler := func(e *errors.Error) {
			parsmakeerErrors = append(parserErrors, e)
		}

		// Parse

		rootNode, err := parser.Parse(src, conf.Config{
			Version:          &version.Version{Major: 5, Minor: 6},
			ErrorHandlerFunc: errorHandler,
		})

		if err != nil {
			log.Fatal("Error:" + err.Error())
		}

		// Dump

		goDumper := dumper.NewDumper(os.Stdout).
			WithTokens().
			WithPositions()

		rootNode.Accept(goDumper)
	}
*/
package parser
