package php_test

// Version-gating tests for the unified PHP 7/8 parser.
//
// The unified parser accepts the full syntax superset for every supported
// profile. Where PHP itself rejects a construct for a given version, the
// rejection happens in the lexer (reserved words, removed casts, heredoc
// rules, `#` comments) or in the grammar actions (curly-brace offsets).
// Newer syntax that still parses on old profiles (match expressions, the
// nullsafe operator, union types, ...) is gated by compile-time validation
// in the linter instead.

import (
	"testing"

	"github.com/shyim/go-phplint/internal/errors"
	"github.com/shyim/go-phplint/internal/position"
	"github.com/shyim/go-phplint/internal/tester"
	"github.com/shyim/go-phplint/internal/token"
	"github.com/shyim/go-phplint/internal/version"
)

func phpVersion(major, minor uint64) version.Version {
	return version.Version{Major: major, Minor: minor}
}

func TestUnifiedMatchKeyword(t *testing.T) {
	// `match` is a plain identifier below PHP 8.0.
	ok := tester.NewParserErrorTestSuite(t)
	ok.Version = phpVersion(7, 4)
	ok.Code = `<?php function match() {}`
	ok.Expected = nil
	ok.Run()

	// Reserved word on PHP 8.0+.
	bad := tester.NewParserErrorTestSuite(t)
	bad.UsePHP8()
	bad.Code = `<?php function match() {}`
	bad.Expected = []*errors.Error{
		{
			Msg: "syntax error: unexpected match, expecting '('",
			Pos: position.NewPosition(1, 1, 15, 20, 15, 20),
		},
	}
	bad.Run()
}

func TestUnifiedEnumReadonlyKeywords(t *testing.T) {
	for _, keyword := range []string{"enum", "readonly"} {
		// Plain identifiers on PHP 7 profiles.
		for _, v := range []version.Version{phpVersion(7, 2), phpVersion(7, 4)} {
			ok := tester.NewParserErrorTestSuite(t)
			ok.Version = v
			ok.Code = `<?php function ` + keyword + `() {}`
			ok.Expected = nil
			ok.Run()
		}

		// Reserved words on PHP 8.0+ (matching the previous PHP 8 parser).
		bad := tester.NewParserErrorTestSuite(t)
		bad.UsePHP8()
		bad.Code = `<?php function ` + keyword + `() {}`
		if bad.Code == `<?php function enum() {}` {
			bad.Expected = []*errors.Error{
				{
					Msg: "syntax error: unexpected enum, expecting '('",
					Pos: position.NewPosition(1, 1, 15, 19, 15, 19),
				},
			}
		} else {
			bad.Expected = []*errors.Error{
				{
					Msg: "syntax error: unexpected readonly, expecting '('",
					Pos: position.NewPosition(1, 1, 15, 23, 15, 23),
				},
			}
		}
		bad.Run()
	}
}

func TestUnifiedFnKeyword(t *testing.T) {
	// `fn` is a plain identifier below PHP 7.4.
	ok := tester.NewParserErrorTestSuite(t)
	ok.Version = phpVersion(7, 2)
	ok.Code = `<?php function fn() {}`
	ok.Expected = nil
	ok.Run()

	// Reserved word on PHP 7.4+.
	bad := tester.NewParserErrorTestSuite(t)
	bad.Code = `<?php function fn() {}`
	bad.Expected = []*errors.Error{
		{
			Msg: "syntax error: unexpected fn, expecting '('",
			Pos: position.NewPosition(1, 1, 15, 17, 15, 17),
		},
	}
	bad.Run()
}

func TestUnifiedRemovedCasts(t *testing.T) {
	// `(real)` and `(unset)` remain valid below PHP 8.0.
	for _, cast := range []string{"(real)$a;", "(unset)$a;"} {
		ok := tester.NewParserErrorTestSuite(t)
		ok.Version = phpVersion(7, 4)
		ok.Code = `<?php $x = ` + cast
		ok.Expected = nil
		ok.Run()
	}

	// Removed on PHP 8.0+.
	real := tester.NewParserErrorTestSuite(t)
	real.UsePHP8()
	real.Code = "<?php \n(real)$a;"
	real.Expected = []*errors.Error{
		{
			Msg: "The (real) cast has been removed, use (float) instead",
			Pos: position.NewPosition(2, 2, 7, 13, 1, 7),
		},
	}
	real.Run()

	unset := tester.NewParserErrorTestSuite(t)
	unset.UsePHP8()
	unset.Code = "<?php \n(unset)$a;"
	unset.Expected = []*errors.Error{
		{
			Msg: "The (unset) cast is no longer supported",
			Pos: position.NewPosition(2, 2, 7, 14, 1, 1),
		},
	}
	unset.Run()
}

func TestUnifiedCurlyOffsets(t *testing.T) {
	// `{expr}` offsets are valid below PHP 8.0.
	ok := tester.NewParserErrorTestSuite(t)
	ok.Version = phpVersion(7, 4)
	ok.Code = `<?php echo $a{0};`
	ok.Expected = nil
	ok.Run()

	// Rejected on PHP 8.0+.
	bad := tester.NewParserErrorTestSuite(t)
	bad.UsePHP8()
	bad.Code = `<?php echo $a{0};`
	bad.Expected = []*errors.Error{
		{
			Msg: "Array and string offset access syntax with curly braces is no longer supported",
			Pos: position.NewPosition(1, 1, 15, 16, 15, 16),
		},
	}
	bad.Run()
}

func TestUnifiedAttributeComment(t *testing.T) {
	// On PHP 7 profiles `#[...]` is a `#` line comment: the rest of the
	// line is skipped and the next line lexes normally.
	suite := tester.NewLexerTokenStructTestSuite(t)
	suite.Version = phpVersion(7, 4)
	suite.Code = "<?php #[Attr]\nfoo();"
	suite.Expected = []*token.Token{
		{ID: token.T_STRING, Value: []byte("foo")},
		{ID: token.ID('('), Value: []byte("(")},
		{ID: token.ID(')'), Value: []byte(")")},
		{ID: token.ID(';'), Value: []byte(";")},
	}
	suite.Run()

	// On PHP 8.0+ it lexes as a real attribute.
	attr := tester.NewLexerTokenStructTestSuite(t)
	attr.UsePHP8()
	attr.Code = `<?php #[Attr]`
	attr.Expected = []*token.Token{
		{ID: token.T_ATTRIBUTE, Value: []byte("#[")},
		{ID: token.T_STRING, Value: []byte("Attr")},
		{ID: token.ID(']'), Value: []byte("]")},
	}
	attr.Run()
}

func TestUnifiedHeredoc(t *testing.T) {
	// Indented closing markers require PHP 7.3+.
	flexible := "<?php\n$value = <<<TEXT\n  content\n  TEXT;\n"
	ok := tester.NewParserErrorTestSuite(t)
	ok.Version = phpVersion(7, 3)
	ok.Code = flexible
	ok.Expected = nil
	ok.Run()

	// Strict markers keep working everywhere, including PHP 7.2.
	strict := "<?php\n$value = <<<TEXT\ncontent\nTEXT;\n"
	ok72 := tester.NewParserErrorTestSuite(t)
	ok72.Version = phpVersion(7, 2)
	ok72.Code = strict
	ok72.Expected = nil
	ok72.Run()
}

func TestUnifiedNullsafeInInterpolation(t *testing.T) {
	// Unbraced `?->` inside interpolations is literal text on PHP 7 and
	// parses cleanly on every profile at parser level.
	for _, v := range []version.Version{phpVersion(7, 4), phpVersion(8, 0)} {
		ok := tester.NewParserErrorTestSuite(t)
		ok.Version = v
		ok.Code = `<?php $a = "x$y?->z";`
		ok.Expected = nil
		ok.Run()
	}

	// Braced interpolations contain full expressions and also parse; the
	// version gate for the nullsafe operator lives in validation.
	braced := tester.NewParserErrorTestSuite(t)
	braced.Version = phpVersion(7, 4)
	braced.Code = `<?php $a = "x{$y?->z}";`
	braced.Expected = nil
	braced.Run()
}
