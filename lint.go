package phplint

import (
	"fmt"
	"sort"

	"github.com/shyim/go-phplint/internal/conf"
	phperrors "github.com/shyim/go-phplint/internal/errors"
	phpparser "github.com/shyim/go-phplint/internal/parser"
)

// Options configures a lint operation.
type Options struct {
	PHPVersion Version
}

// Lint parses and compile-validates one PHP source file.
//
// Syntax and compile failures are returned as diagnostics. The error return is
// reserved for invalid options or an internal linter failure.
func Lint(filename string, source []byte, options Options) (diagnostics []Diagnostic, err error) {
	if !options.PHPVersion.valid() {
		return nil, fmt.Errorf("invalid PHP version %q", options.PHPVersion)
	}

	defer func() {
		if recovered := recover(); recovered != nil {
			diagnostics = nil
			err = fmt.Errorf("internal parser failure: %v", recovered)
		}
	}()

	parseSource, modernDiagnostics := prepareSource(source, options.PHPVersion, filename)
	diagnostics = append(diagnostics, modernDiagnostics...)

	var parserDiagnostics []Diagnostic
	root, parseErr := phpparser.Parse(parseSource, conf.Config{
		Version: options.PHPVersion.internal(),
		ErrorHandlerFunc: func(parseError *phperrors.Error) {
			start, end := positionFromInternal(parseError.Pos)
			parserDiagnostics = append(parserDiagnostics, Diagnostic{
				Filename: filename,
				Message:  normalizeParserMessage(parseError.Msg),
				Phase:    PhaseParse,
				Start:    start,
				End:      end,
			})
		},
	})
	if parseErr != nil {
		return nil, fmt.Errorf("initialize PHP %s parser: %w", options.PHPVersion, parseErr)
	}

	diagnostics = append(diagnostics, parserDiagnostics...)

	if len(parserDiagnostics) == 0 && root != nil {
		diagnostics = append(
			diagnostics,
			validate(root, source, filename, options.PHPVersion)...,
		)
	}

	sortDiagnostics(diagnostics)
	return deduplicateDiagnostics(diagnostics), nil
}

func normalizeParserMessage(message string) string {
	if message == "" {
		return "syntax error"
	}
	return message
}

func sortDiagnostics(diagnostics []Diagnostic) {
	sort.SliceStable(diagnostics, func(i, j int) bool {
		left, right := diagnostics[i], diagnostics[j]
		if left.Start.Offset != right.Start.Offset {
			return left.Start.Offset < right.Start.Offset
		}
		if left.Phase != right.Phase {
			return left.Phase < right.Phase
		}
		return left.Message < right.Message
	})
}

func deduplicateDiagnostics(diagnostics []Diagnostic) []Diagnostic {
	if len(diagnostics) < 2 {
		return diagnostics
	}

	result := diagnostics[:0]
	var previous Diagnostic
	for index, diagnostic := range diagnostics {
		if index > 0 &&
			diagnostic.Phase == previous.Phase &&
			diagnostic.Start.Offset == previous.Start.Offset &&
			diagnostic.End.Offset == previous.End.Offset &&
			diagnostic.Message == previous.Message {
			continue
		}

		result = append(result, diagnostic)
		previous = diagnostic
	}

	return result
}
