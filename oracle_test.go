package phplint

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestNativePHPOracle(t *testing.T) {
	phpCommand := os.Getenv("PHPLINT_PHP_BINARY")
	if phpCommand == "" {
		phpCommand = "php"
	}
	phpBinary, err := exec.LookPath(phpCommand)
	if err != nil {
		t.Skip("php binary is not available")
	}

	versionOutput, err := exec.Command(
		phpBinary,
		"-n",
		"-r",
		"echo PHP_MAJOR_VERSION, '.', PHP_MINOR_VERSION;",
	).Output()
	if err != nil {
		t.Skipf("cannot determine native PHP version: %v", err)
	}

	version, err := ParseVersion(string(versionOutput))
	if err != nil {
		t.Skipf("native PHP %s is outside the supported oracle matrix", versionOutput)
	}

	tests := []struct {
		name   string
		source string
	}{
		{name: "basic", source: "<?php echo 'ok';"},
		{name: "mixed html", source: "<h1>Title</h1><?php echo 'body'; ?>"},
		{name: "short tag", source: "<? echo 'short';"},
		{name: "flexible heredoc", source: "<?php\n$value = <<<TEXT\n  content\n  TEXT;\n"},
		{name: "trailing call comma", source: "<?php trim('value',);"},
		{name: "list assignment reference", source: "<?php [$first, &$second] = $values;"},
		{name: "typed property", source: "<?php class Value { public int $number; }"},
		{name: "arrow function", source: "<?php $double = fn(int $value): int => $value * 2;"},
		{name: "null coalescing assignment", source: "<?php $value ??= 'default';"},
		{name: "array unpacking", source: "<?php $all = [0, ...$values];"},
		{name: "numeric literal separator", source: "<?php $million = 1_000_000;"},
		{name: "typed class", source: "<?php class A { public int|string $value; }"},
		{name: "union type", source: "<?php function normalize(int|string $value): int|string { return $value; }"},
		{name: "match expression", source: "<?php $label = match ($value) { 1 => 'one', default => 'other' };"},
		{name: "nullsafe operator", source: "<?php $name = $user?->profile()?->name;"},
		{name: "named arguments", source: "<?php trim(string: ' value ');"},
		{name: "constructor promotion", source: "<?php class C { function __construct(public string $name) {} }"},
		{name: "interface", source: "<?php interface A { public function run(): void; }"},
		{name: "abstract method", source: "<?php abstract class A { abstract public function run(): void; }"},
		{name: "enum", source: "<?php enum Status { case Active; }"},
		{name: "intersection type", source: "<?php function handle(Countable&Iterator $value): void {}"},
		{name: "first class callable", source: "<?php $callable = strlen(...);"},
		{name: "readonly property", source: "<?php class Value { public readonly int $number; }"},
		{name: "new parameter initializer", source: "<?php function make($value = new stdClass()) {}"},
		{name: "readonly class", source: "<?php readonly class Value { public int $number; }"},
		{name: "dnf type", source: "<?php function handle((Countable&Iterator)|Stringable $value): void {}"},
		{name: "true type", source: "<?php function yes(): true { return true; }"},
		{name: "trait constant", source: "<?php trait Values { public const ANSWER = 42; }"},
		{name: "typed class constant", source: "<?php class C { public const string NAME = 'c'; }"},
		{name: "dynamic class constant", source: "<?php class C { const NAME = 'c'; } $name = 'NAME'; echo C::{$name};"},
		{name: "readonly anonymous class", source: "<?php $value = new readonly class { public function x() {} };"},
		{name: "property hooks", source: "<?php class A { public string $name { get => $this->name; set => $value; } }"},
		{name: "reference property hook", source: "<?php class A { public string $name { &get => $this->name; } }"},
		{name: "attributed final hook", source: "<?php class A { public string $name { #[Deprecated] final get => $this->name; } }"},
		{name: "typed set hook", source: "<?php class A { public string $name { set(string $other) => $other; } }"},
		{name: "method body starting with get", source: "<?php class A { function run() { get(); } }"},
		{name: "asymmetric visibility", source: "<?php class A { public private(set) string $name; }"},
		{name: "new dereference", source: "<?php class C { function x() {} } new C()->x();"},
		{name: "pipe", source: "<?php $length = ' value ' |> trim(...) |> strlen(...);"},
		{name: "clone with", source: "<?php $copy = clone($value, ['name' => 'copy']);"},
		{name: "void cast", source: "<?php (void) do_work();"},
		{name: "constant attribute", source: "<?php #[Deprecated] const OLD_VALUE = 1;"},
		{name: "final promoted property", source: "<?php class C { function __construct(final public string $name) {} }"},
		{name: "override class constant", source: "<?php class ParentValue { const VALUE = 1; } class ChildValue extends ParentValue { #[\\Override] const VALUE = 2; }"},
		{name: "first class callable constant", source: "<?php const HANDLER = strlen(...);"},
		{name: "static closure constant", source: "<?php const HANDLER = static function () { return 1; };"},
		{name: "arbitrary static initializer", source: "<?php function f() { static $x = strlen('x'); }"},

		{name: "parse error", source: "<?php function broken( {"},
		{name: "duplicate function", source: "<?php function x() {} function X() {}"},
		{name: "duplicate class", source: "<?php class X {} class x {}"},
		{name: "duplicate method", source: "<?php class X { function y() {} function Y() {} }"},
		{name: "duplicate property", source: "<?php class X { public $y; public $y; }"},
		{name: "duplicate constant", source: "<?php class X { const Y = 1; const Y = 2; }"},
		{name: "duplicate parameter", source: "<?php function x($value, $value) {}"},
		{name: "break outside loop", source: "<?php break;"},
		{name: "continue outside loop", source: "<?php continue;"},
		{name: "too many break levels", source: "<?php while (true) { break 2; }"},
		{name: "zero break levels", source: "<?php while (true) { break 0; }"},
		{name: "multiple visibility", source: "<?php class X { public private $value; }"},
		{name: "abstract final class", source: "<?php abstract final class X {}"},
		{name: "abstract method body", source: "<?php abstract class X { abstract function y() {} }"},
		{name: "concrete method no body", source: "<?php class X { function y(); }"},
		{name: "interface method body", source: "<?php interface X { function y() {} }"},
		{name: "readonly untyped", source: "<?php class X { public readonly $value; }"},
		{name: "readonly default", source: "<?php class X { public readonly int $value = 1; }"},
		{name: "readonly static", source: "<?php class X { public static readonly int $value; }"},
		{name: "enum property", source: "<?php enum X { public int $value; }"},
		{name: "void parameter", source: "<?php function x(void $value) {}"},
		{name: "duplicate union", source: "<?php function x(A|A $value) {}"},
		{name: "redundant bool union", source: "<?php function x(bool|false $value) {}"},
		{name: "mixed union", source: "<?php function x(mixed|A $value) {}"},
		{name: "builtin intersection", source: "<?php function x(A&string $value) {}"},
		{name: "nullable mixed", source: "<?php function x(?mixed $value) {}"},
		{name: "constructor return", source: "<?php class X { function __construct(): void {} }"},
		{name: "destructor parameters", source: "<?php class X { function __destruct($value) {} }"},
		{name: "function call constant", source: "<?php const X = strlen('x');"},
		{name: "new class constant", source: "<?php class X { const VALUE = new stdClass(); }"},
		{name: "new property default", source: "<?php class X { public $value = new stdClass(); }"},
		{name: "nonstatic closure constant", source: "<?php const X = function () {};"},
		{name: "static hooked property", source: "<?php class X { public static string $value { get => ''; } }"},
		{name: "readonly hooked property", source: "<?php class X { public readonly string $value { get => ''; } }"},
		{name: "duplicate property hook", source: "<?php class X { public string $value { get => ''; get => ''; } }"},
		{name: "get hook parameters", source: "<?php class X { public string $value { get($arg) => $arg; } }"},
		{name: "set hook parameters", source: "<?php class X { public string $value { set($a, $b) => $a; } }"},
		{name: "bodyless concrete hook", source: "<?php class X { public string $value { get; } }"},
		{name: "weaker asymmetric visibility", source: "<?php class X { private public(set) string $value; }"},
		{name: "callable class constant type", source: "<?php class X { const callable VALUE = strlen(...); }"},
		{name: "typed class constant mismatch", source: "<?php class X { const int VALUE = 'wrong'; }"},
		{name: "second typed class constant mismatch", source: "<?php class X { const int A = 1, B = 'wrong'; }"},
		{name: "readonly method", source: "<?php class X { readonly function run() {} }"},
		{name: "static class constant", source: "<?php class X { static const VALUE = 1; }"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			nativePassed, nativeOutput := nativeLint(t, phpBinary, test.source)
			diagnostics, lintErr := Lint(
				"oracle.php",
				[]byte(test.source),
				Options{PHPVersion: version},
			)
			linterPassed := lintErr == nil && len(diagnostics) == 0
			if linterPassed == nativePassed {
				return
			}

			t.Fatalf(
				"pass/fail mismatch for PHP %s\nsource: %s\nnative passed: %t\nnative: %s\nlinter error: %v\nlinter diagnostics: %s",
				version,
				test.source,
				nativePassed,
				nativeOutput,
				lintErr,
				formatDiagnostics(diagnostics),
			)
		})
	}
}

func nativeLint(t *testing.T, phpBinary, source string) (bool, string) {
	t.Helper()

	command := exec.Command(phpBinary, "-n", "-d", "short_open_tag=1", "-l")
	command.Stdin = strings.NewReader(source)
	output, err := command.CombinedOutput()
	if err == nil {
		return true, strings.TrimSpace(string(output))
	}

	var exitError *exec.ExitError
	if !errors.As(err, &exitError) {
		t.Fatalf("run native PHP: %v", err)
	}
	return false, strings.TrimSpace(string(output))
}

func formatDiagnostics(diagnostics []Diagnostic) string {
	var output bytes.Buffer
	for index, diagnostic := range diagnostics {
		if index > 0 {
			output.WriteString("; ")
		}
		_, _ = fmt.Fprintf(
			&output,
			"%d:%d %s",
			diagnostic.Start.Line,
			diagnostic.Start.Column,
			diagnostic.Message,
		)
	}
	return output.String()
}
