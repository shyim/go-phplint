package phplint

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/shyim/go-phplint/internal/ast"
	"github.com/shyim/go-phplint/internal/conf"
	phperrors "github.com/shyim/go-phplint/internal/errors"
	phpparser "github.com/shyim/go-phplint/internal/parser"
	"github.com/shyim/go-phplint/internal/token"
)

// benchFixture is one PHP source file linted against one language profile.
type benchFixture struct {
	name    string
	file    string
	version Version
}

var benchFixtures = []benchFixture{
	// Procedural/OO PHP 7 code: no version gates, classic parser path.
	{name: "legacy-7.4", file: "legacy.php", version: PHP74},
	// Modern PHP 8 code: enums, attributes, promotion, typed constants.
	{name: "modern-8.4", file: "modern.php", version: PHP84},
	// HTML template: inline-HTML heavy lexing with alternative syntax blocks.
	{name: "template-8.3", file: "template.php", version: PHP83},
	// Broken source: error-recovery, diagnostic sorting and source-line attachment.
	{name: "invalid-8.4", file: "invalid.php", version: PHP84},
}

func readBenchFixture(tb testing.TB, file string) []byte {
	tb.Helper()

	source, err := os.ReadFile(filepath.Join("testdata", "bench", file))
	if err != nil {
		tb.Fatalf("read benchmark fixture %s: %v", file, err)
	}
	return source
}

// BenchmarkLint measures the full public entry point: source preparation,
// parsing and compile-time validation.
func BenchmarkLint(b *testing.B) {
	for _, fixture := range benchFixtures {
		source := readBenchFixture(b, fixture.file)
		options := Options{PHPVersion: fixture.version}

		b.Run(fixture.name, func(b *testing.B) {
			b.SetBytes(int64(len(source)))
			for b.Loop() {
				if _, err := Lint(fixture.file, source, options); err != nil {
					b.Fatalf("Lint() error = %v", err)
				}
			}
		})
	}
}

// BenchmarkLintVersionProfiles lints the same modern source against several
// language profiles. Older profiles additionally exercise the feature-gate
// diagnostics emitted during source preparation.
func BenchmarkLintVersionProfiles(b *testing.B) {
	source := readBenchFixture(b, "modern.php")

	for _, version := range []Version{PHP74, PHP80, PHP82, PHP84, PHP86} {
		options := Options{PHPVersion: version}

		b.Run(version.String(), func(b *testing.B) {
			b.SetBytes(int64(len(source)))
			for b.Loop() {
				if _, err := Lint("modern.php", source, options); err != nil {
					b.Fatalf("Lint() error = %v", err)
				}
			}
		})
	}
}

// BenchmarkLintLargeFile lints a generated file of a few thousand lines, which
// is where parser allocations and diagnostic bookkeeping dominate.
//
// The work runs as a sub-benchmark. Go splits -bench on '/' and '|', and an
// earlier alternative BenchmarkLint/… is a partial match of this function's
// name, so the top-level result is never printed.
func BenchmarkLintLargeFile(b *testing.B) {
	source := generateLargeSource(100)
	options := Options{PHPVersion: PHP84}

	b.Run("file", func(b *testing.B) {
		b.SetBytes(int64(len(source)))
		for b.Loop() {
			if _, err := Lint("large.php", source, options); err != nil {
				b.Fatalf("Lint() error = %v", err)
			}
		}
	})
}

// BenchmarkPrepareSource isolates the token-level rewriting stage that runs
// before the parser sees the source.
func BenchmarkPrepareSource(b *testing.B) {
	for _, fixture := range benchFixtures {
		source := readBenchFixture(b, fixture.file)

		b.Run(fixture.name, func(b *testing.B) {
			b.SetBytes(int64(len(source)))
			var tokens []*token.Token
			_, err := phpparser.Parse(source, conf.Config{
				Version:          fixture.version.internal(),
				Tokens:           &tokens,
				ErrorHandlerFunc: func(*phperrors.Error) {},
			})
			if err != nil {
				b.Fatalf("Parse() error = %v", err)
			}
			b.ResetTimer()
			for b.Loop() {
				_ = prepareSource(source, tokens, fixture.version, fixture.file)
			}
		})
	}
}

// BenchmarkParse isolates the lexer and parser without validation.
func BenchmarkParse(b *testing.B) {
	for _, fixture := range benchFixtures {
		source := readBenchFixture(b, fixture.file)
		config := conf.Config{
			Version:          fixture.version.internal(),
			ErrorHandlerFunc: func(*phperrors.Error) {},
		}

		b.Run(fixture.name, func(b *testing.B) {
			b.SetBytes(int64(len(source)))
			for b.Loop() {
				if _, err := phpparser.Parse(source, config); err != nil {
					b.Fatalf("Parse() error = %v", err)
				}
			}
		})
	}
}

// BenchmarkValidate isolates the compile-time checks by parsing once outside of
// the measured loop.
func BenchmarkValidate(b *testing.B) {
	for _, fixture := range benchFixtures {
		if fixture.file == "invalid.php" {
			// Validation is skipped when the parser reports diagnostics.
			continue
		}

		source := readBenchFixture(b, fixture.file)
		root := parseForBenchmark(b, source, fixture.version)

		b.Run(fixture.name, func(b *testing.B) {
			b.SetBytes(int64(len(source)))
			for b.Loop() {
				_ = validate(root, source, fixture.file, fixture.version)
			}
		})
	}
}

func parseForBenchmark(tb testing.TB, source []byte, version Version) ast.Vertex {
	tb.Helper()

	root, err := phpparser.Parse(source, conf.Config{
		Version:          version.internal(),
		ErrorHandlerFunc: func(*phperrors.Error) {},
	})
	if err != nil {
		tb.Fatalf("Parse() error = %v", err)
	}
	if root == nil {
		tb.Fatal("Parse() returned no root node")
	}
	return root
}

// generateLargeSource builds a syntactically valid PHP 8 file with the given
// number of classes, so declarations stay unique.
func generateLargeSource(classes int) []byte {
	var builder strings.Builder
	builder.WriteString("<?php\n\ndeclare(strict_types=1);\n\nnamespace App\\Generated;\n\n")

	for index := range classes {
		fmt.Fprintf(&builder, `final class Service%[1]d
{
    public const int VERSION = %[1]d;

    public function __construct(
        private readonly string $name = 'service-%[1]d',
        private array $options = [],
    ) {
    }

    public function handle(int|string $input, ?callable $next = null): string
    {
        $value = match (true) {
            is_int($input) => $input * %[1]d,
            $input === '' => 'empty',
            default => strtoupper($input),
        };

        foreach ($this->options as $key => $option) {
            $value .= sprintf('|%%s=%%s', $key, (string) $option);
        }

        return $next !== null ? $next($value) : $this->name . ':' . $value;
    }

    public function options(): \Generator
    {
        yield from $this->options;
    }
}

`, index)
	}

	return []byte(builder.String())
}
