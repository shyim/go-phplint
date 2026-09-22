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
		{name: "fully qualified lowercase generator return type", source: `<?php function lines(): \generator { yield "line"; }`},
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
		{name: "enum keyword case names", source: "<?php enum Status: string { case NEW = 'new'; case DEFAULT = 'default'; case List = 'list'; }"},
		{name: "keyword class constant name", source: "<?php class A { const NEW = 1; }"},
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
		{name: "lone asymmetric property", source: "<?php class A { public(set) string $name; }"},
		{name: "asymmetric constructor promotion", source: "<?php class A { function __construct(public(set) int $name) {} }"},
		{name: "asymmetric earlier parameter", source: "<?php class A { function __construct(private int $other, public(set) int $name) {} }"},
		{name: "asymmetric method visibility", source: "<?php class A { protected function __construct(public(set) int $name) {} }"},
		{name: "stronger asymmetric set", source: "<?php class A { protected private(set) string $name; }"},
		{name: "dnf property type", source: "<?php class A { public (A&B)|C $name; }"},
		{name: "dnf promoted type", source: "<?php class A { function __construct(public (A&B)|C $name) {} }"},
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
		{name: "enum case named class", source: "<?php enum X { case CLASS; }"},
		{name: "class constant named class", source: "<?php class X { const CLASS = 1; }"},
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
		{name: "asymmetric method modifier", source: "<?php class A { public(set) function foo() {} }"},
		{name: "asymmetric constant modifier", source: "<?php class A { public(set) const X = 1; }"},
		{name: "asymmetric function parameter", source: "<?php function foo(public(set) int $x) {}"},
		{name: "asymmetric method parameter", source: "<?php class A { function foo(public(set) int $x) {} }"},
		{name: "asymmetric unknown operation", source: "<?php class A { public(foo) string $x; }"},
		{name: "asymmetric trait alias", source: "<?php class A { use T { foo as public(set); } }"},
		{name: "variadic asymmetric promotion", source: "<?php class A { function __construct(public(set) int ...$x) {} }"},
		{name: "set before stronger get", source: "<?php class A { public(set) private string $name; }"},
		{name: "promoted set before stronger get", source: "<?php class A { function __construct(public(set) private int $name) {} }"},
		{name: "untyped asymmetric property", source: "<?php class A { public(set) $name; }"},
		{name: "untyped combined asymmetric property", source: "<?php class A { public private(set) $name; }"},
		{name: "untyped asymmetric promotion", source: "<?php class A { function __construct(public(set) $name) {} }"},
		{name: "two asymmetric set visibilities", source: "<?php class A { public(set) private(set) string $name; }"},
		{name: "bare parenthesized intersection", source: "<?php function f((A&B) $name) {}"},
		{name: "bare parenthesized property intersection", source: "<?php class A { public (A&B) $name; }"},
		{name: "asymmetric static property", source: "<?php class A { public(set) static string $name; }"},
		{name: "public method call", source: "<?php class A { function public($name) {} } $a = new A; $a->public(set);"},
		{name: "static call with argument", source: "<?php class A { static function public($name) {} } A::public(1);"},
		{name: "abstract hooked property", source: "<?php abstract class A { abstract public string $name { set; } }"},
		{name: "interface hooked property", source: "<?php interface I { public string $name { get; } }"},
		{name: "interface hook with body", source: "<?php interface I { public string $name { get => 1; } }"},
		{name: "abstract property without abstract hook", source: "<?php abstract class A { abstract public string $name { set => $value; } }"},
		{name: "abstract property with semicolon and body", source: "<?php abstract class A { abstract public string $name { get; set => $value; } }"},
		{name: "final abstract hook", source: "<?php abstract class A { abstract public string $name { final get; } }"},
		{name: "final interface hook", source: "<?php interface I { public string $name { final get; } }"},
		{name: "final hook after abstract hook", source: "<?php abstract class A { abstract public string $name { get; final set; } }"},
		{name: "final hook with body", source: "<?php class A { public string $name { final get => 1; } }"},
		{name: "final body beside abstract hook", source: "<?php abstract class A { abstract public string $name { get; final set => $value; } }"},
		{name: "final body before abstract hook", source: "<?php abstract class A { abstract public string $name { final get => 1; set; } }"},
		{name: "intersection set hook parameter", source: "<?php class A { public A&B $name { set(A&B $value) => $value; } }"},
		{name: "dnf set hook parameter", source: "<?php class A { public (A&B)|C $name { set((A&B)|C $value) => $value; } }"},
		{name: "by-ref set hook parameter", source: "<?php class A { public string $name { set(string &$value) => $value; } }"},
		{name: "protected interface hook", source: "<?php interface I { protected string $name { get; } }"},
		{name: "private interface hook", source: "<?php interface I { private string $name { get; } }"},
		{name: "asymmetric get-only virtual property", source: "<?php class A { public private(set) string $name { get => 1; } }"},
		{name: "asymmetric property with set hook", source: "<?php class A { public private(set) string $name { get => 1; set => $value; } }"},
		{name: "asymmetric backed property", source: "<?php class A { public private(set) string $name = 'a' { get => $this->name; } }"},
		{name: "public set get-only", source: "<?php class A { public(set) string $name { get => 1; } }"},
		{name: "equivalent public set get-only", source: "<?php class A { public public(set) string $name { get => 1; } }"},
		{name: "equivalent private set get-only", source: "<?php class A { private private(set) string $name { get => 1; } }"},
		{name: "get reads property", source: "<?php class A { public private(set) string $name { get => $this->name; } }"},
		{name: "get block reads property", source: "<?php class A { public private(set) string $name { get { return $this->name; } } }"},
		{name: "nullsafe property read", source: "<?php class A { public private(set) string $name { get => $this?->name; } }"},
		{name: "wrong case property read", source: "<?php class A { public private(set) string $name { get => $this->Name; } }"},
		{name: "closure property read", source: "<?php class A { public private(set) string $name { get { return (function () { return $this->name; })(); } } }"},
		{name: "arrow property read", source: "<?php class A { public private(set) string $name { get => (fn () => $this->name)(); } }"},
		{name: "write-only virtual property", source: "<?php class A { public private(set) string $name { set { $GLOBALS['a'] = $value; } } }"},
		{name: "short set hook", source: "<?php class A { public private(set) string $name { set => $value; } }"},
		{name: "set writes property", source: "<?php class A { public private(set) string $name { set { $this->name = $value; } } }"},
		{name: "virtual hooked default", source: "<?php class A { public private(set) string $name = 'a' { get => 1; } }"},
		{name: "virtual hooked default without set visibility", source: "<?php class A { public string $name = 'a' { get => 1; } }"},
		{name: "interface protected set", source: "<?php interface I { protected(set) string $name { get; set; } }"},
		{name: "interface private set", source: "<?php interface I { private(set) string $name { get; set; } }"},
		{name: "interface public protected set", source: "<?php interface I { public protected(set) string $name { get; set; } }"},
		{name: "interface public private set", source: "<?php interface I { public private(set) string $name { get; set; } }"},
		{name: "by-ref get with set hook", source: "<?php class A { public string $name { &get => $this->name; set => $value; } }"},
		{name: "by-ref get on virtual property", source: "<?php class A { public string $name { &get => 1; set { $foo = $value; } } }"},
		{name: "abstract asymmetric get hook", source: "<?php abstract class A { abstract public private(set) string $name { get; } }"},
		{name: "abstract asymmetric by-ref get hook", source: "<?php abstract class A { abstract public private(set) string $name { &get; } }"},
		{name: "nested function property use", source: "<?php class A { public private(set) string $name { get { function f() { return $this->name; } return 1; } } }"},
		{name: "parenthesized this property", source: "<?php class A { public private(set) string $name { get => ($this)->name; } }"},
		{name: "double parenthesized this property", source: "<?php class A { public private(set) string $name { get => (($this))->name; } }"},
		{name: "parenthesized nullsafe this", source: "<?php class A { public private(set) string $name { get => ($this)?->name; } }"},
		{name: "constant computed property", source: "<?php class A { public private(set) string $name { get => $this->{\"name\"}; } }"},
		{name: "nullsafe constant computed property", source: "<?php class A { public private(set) string $name { get => $this?->{\"name\"}; } }"},
		{name: "single quoted computed property", source: "<?php class A { public private(set) string $name { get => $this->{'name'}; } }"},
		{name: "variable computed property", source: "<?php class A { public private(set) string $name { get => $this->{$n}; } }"},
		{name: "interpolated property in short hook", source: "<?php class A { public string $name { get => \"{$this->name}\"; } }"},
		{name: "interpolated property in hook block", source: "<?php class A { public string $name { get { return \"{$this->name}\"; } } }"},
		{name: "dollar curly property in short hook", source: "<?php class A { public string $name { get => \"${this->name}\"; } }"},
		{name: "interpolated property backs asymmetric hook", source: "<?php class A { public private(set) string $name { get => \"{$this->name}\"; } }"},
		{name: "unhooked abstract property", source: "<?php abstract class A { abstract public string $name; }"},
		{name: "unhooked interface property", source: "<?php interface I { public string $name; }"},
		{name: "enum case public(set)", source: "<?php enum E { case public(set); }"},
		{name: "enum case protected(set)", source: "<?php enum E { case protected(set); }"},
		{name: "enum case private(set)", source: "<?php enum E { case private(set); }"},
		{name: "const name public(set)", source: "<?php class A { const public(set) = 1; }"},
		{name: "const name protected(set)", source: "<?php class A { const protected(set) = 1; }"},
		{name: "const name private(set)", source: "<?php class A { const private(set) = 1; }"},
		{name: "fetch public(set)", source: "<?php class A { const X = 1; } echo A::public(set);"},
		{name: "fetch protected(set)", source: "<?php class A { const X = 1; } echo A::protected(set);"},
		{name: "fetch private(set)", source: "<?php class A { const X = 1; } echo A::private(set);"},
		{name: "trait method public(set)", source: "<?php class A { use T { public(set) as foo; } }"},
		{name: "trait method protected(set)", source: "<?php class A { use T { protected(set) as foo; } }"},
		{name: "trait method private(set)", source: "<?php class A { use T { private(set) as foo; } }"},
		{name: "namespace public(set)", source: "<?php namespace public(set);"},
		{name: "namespace protected(set)", source: "<?php namespace protected(set);"},
		{name: "namespace private(set)", source: "<?php namespace private(set);"},
		{name: "method name public(set)", source: "<?php class A { function public(set)() {} }"},
		{name: "method name protected(set)", source: "<?php class A { function protected(set)() {} }"},
		{name: "method name private(set)", source: "<?php class A { function private(set)() {} }"},
		{name: "hook name public(set)", source: "<?php class A { public string $name { public(set); } }"},
		{name: "hook name protected(set)", source: "<?php class A { public string $name { protected(set); } }"},
		{name: "hook name private(set)", source: "<?php class A { public string $name { private(set); } }"},
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
