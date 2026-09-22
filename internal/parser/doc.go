/*
A Parser for PHP written in Go

Package usage example:

	package main

	import (
		"log"
		"os"

		"github.com/shyim/go-phplint/internal/conf"
		"github.com/shyim/go-phplint/internal/errors"
		"github.com/shyim/go-phplint/internal/parser"
		"github.com/shyim/go-phplint/internal/version"
		"github.com/shyim/go-phplint/internal/visitor/dumper"
	)

	func main() {
		src := []byte(`<? echo "Hello world";`)

		// Error handler

		var parserErrors []*errors.Error
		errorHandler := func(e *errors.Error) {
			parserErrors = append(parserErrors, e)
		}

		// Parse

		rootNode, err := parser.Parse(src, conf.Config{
			Version:          &version.Version{Major: 8, Minor: 4},
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
