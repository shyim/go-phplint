package phplint

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/shyim/go-phplint/internal/conf"
	phperrors "github.com/shyim/go-phplint/internal/errors"
	phpparser "github.com/shyim/go-phplint/internal/parser"
	"github.com/shyim/go-phplint/internal/token"
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

	var tokens []*token.Token
	var parserDiagnostics []Diagnostic
	root, parseErr := phpparser.Parse(source, conf.Config{
		Version: options.PHPVersion.internal(),
		Tokens:  &tokens,
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

	diagnostics = append(diagnostics, prepareSource(source, tokens, options.PHPVersion, filename)...)
	diagnostics = append(diagnostics, parserDiagnostics...)

	if len(parserDiagnostics) == 0 && root != nil {
		diagnostics = append(
			diagnostics,
			validate(root, source, filename, options.PHPVersion)...,
		)
	}

	sortDiagnostics(diagnostics)
	diagnostics = deduplicateDiagnostics(diagnostics)
	attachSourceLines(diagnostics, source)
	return diagnostics, nil
}

func attachSourceLines(diagnostics []Diagnostic, source []byte) {
	if len(diagnostics) == 0 {
		return
	}

	lines := bytes.Split(source, []byte("\n"))
	for index := range diagnostics {
		line := diagnostics[index].Start.Line
		if line < 1 || line > len(lines) {
			continue
		}
		diagnostics[index].SourceLine = string(
			bytes.TrimSuffix(lines[line-1], []byte("\r")),
		)
	}
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
