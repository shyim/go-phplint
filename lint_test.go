package phplint

import (
	"strings"
	"testing"
)

func TestParseVersion(t *testing.T) {
	t.Parallel()

	for _, expected := range SupportedVersions() {
		expected := expected
		t.Run(expected.String(), func(t *testing.T) {
			t.Parallel()
			actual, err := ParseVersion(expected.String())
			if err != nil {
				t.Fatalf("ParseVersion() error = %v", err)
			}
			if actual != expected {
				t.Fatalf("ParseVersion() = %v, want %v", actual, expected)
			}
		})
	}

	for _, invalid := range []string{"", "7.1", "8", "8.6.1", "9.0"} {
		if _, err := ParseVersion(invalid); err == nil {
			t.Errorf("ParseVersion(%q) unexpectedly succeeded", invalid)
		}
	}
}

func TestLintBasicSyntaxAndCompileErrors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		source     string
		wantPhrase string
	}{
		{
			name:   "valid",
			source: "<?php function greet(string $name): string { return \"Hi $name\"; }",
		},
		{
			name:   "fully qualified lowercase generator return type",
			source: `<?php function lines(): \generator { yield "line"; }`,
		},
		{
			name:       "syntax error",
			source:     "<?php function broken( {",
			wantPhrase: "syntax error",
		},
		{
			name:       "break outside loop",
			source:     "<?php break;",
			wantPhrase: "not in a loop",
		},
		{
			name:       "duplicate method",
			source:     "<?php class A { function x() {} function X() {} }",
			wantPhrase: "cannot redeclare method",
		},
		{
			name:       "void return value",
			source:     "<?php function x(): void { return 1; }",
			wantPhrase: "must not return a value",
		},
		{
			name:       "invalid readonly property",
			source:     "<?php class A { public readonly $value; }",
			wantPhrase: "must have a type",
		},
		{
			name:   "semi-reserved keywords as enum case names",
			source: "<?php enum Status: string { case NEW = 'new'; case DEFAULT = 'default'; case List = 'list'; case Function = 'function'; }",
		},
		{
			name:   "semi-reserved keyword as attributed enum case name",
			source: "<?php enum Status { #[Deprecated] case Print; }",
		},
		{
			name:   "semi-reserved keyword as class constant name",
			source: "<?php class A { const NEW = 1; const Default = 2; }",
		},
		{
			name:       "enum case named class",
			source:     "<?php enum Status { case CLASS; }",
			wantPhrase: "must not be called 'class'",
		},
		{
			name:       "class constant named class",
			source:     "<?php class A { const class = 1; }",
			wantPhrase: "must not be called 'class'",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			diagnostics, err := Lint(
				"test.php",
				[]byte(test.source),
				Options{PHPVersion: PHP85},
			)
			if err != nil {
				t.Fatalf("Lint() error = %v", err)
			}
			if test.wantPhrase == "" {
				if len(diagnostics) != 0 {
					t.Fatalf("Lint() diagnostics = %#v, want none", diagnostics)
				}
				return
			}
			if !diagnosticsContain(diagnostics, test.wantPhrase) {
				t.Fatalf("Lint() diagnostics = %#v, want phrase %q", diagnostics, test.wantPhrase)
			}
		})
	}
}

func TestLintAttachesSourceLines(t *testing.T) {
	t.Parallel()

	diagnostics, err := Lint(
		"test.php",
		[]byte("<?php\nclass A {\n\tfunction x() {}\n\tfunction X() {}\n}\n"),
		Options{PHPVersion: PHP85},
	)
	if err != nil {
		t.Fatalf("Lint() error = %v", err)
	}
	if len(diagnostics) != 1 {
		t.Fatalf("Lint() diagnostics = %#v, want one", diagnostics)
	}
	if got, want := diagnostics[0].SourceLine, "\tfunction X() {}"; got != want {
		t.Fatalf("Lint() SourceLine = %q, want %q", got, want)
	}
}

func TestLintVersionBoundaries(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		source     string
		before     Version
		introduced Version
	}{
		{
			name:       "flexible heredoc",
			source:     "<?php\n$value = <<<TEXT\n  content\n  TEXT;\n",
			before:     PHP72,
			introduced: PHP73,
		},
		{
			name:       "trailing call comma",
			source:     "<?php trim('value',);",
			before:     PHP72,
			introduced: PHP73,
		},
		{
			name:       "list assignment reference",
			source:     "<?php [$first, &$second] = $values;",
			before:     PHP72,
			introduced: PHP73,
		},
		{
			name:       "typed property",
			source:     "<?php class Value { public int $number; }",
			before:     PHP73,
			introduced: PHP74,
		},
		{
			name:       "arrow function",
			source:     "<?php $double = fn(int $value): int => $value * 2;",
			before:     PHP73,
			introduced: PHP74,
		},
		{
			name:       "null coalescing assignment",
			source:     "<?php $value ??= 'default';",
			before:     PHP73,
			introduced: PHP74,
		},
		{
			name:       "array unpacking",
			source:     "<?php $all = [0, ...$values];",
			before:     PHP73,
			introduced: PHP74,
		},
		{
			name:       "numeric literal separator",
			source:     "<?php $million = 1_000_000;",
			before:     PHP73,
			introduced: PHP74,
		},
		{
			name:       "union type",
			source:     "<?php function normalize(int|string $value): int|string { return $value; }",
			before:     PHP74,
			introduced: PHP80,
		},
		{
			name:       "match expression",
			source:     "<?php $label = match ($value) { 1 => 'one', default => 'other' };",
			before:     PHP74,
			introduced: PHP80,
		},
		{
			name:       "nullsafe operator",
			source:     "<?php $name = $user?->profile()?->name;",
			before:     PHP74,
			introduced: PHP80,
		},
		{
			name:       "named arguments",
			source:     "<?php trim(string: ' value ');",
			before:     PHP74,
			introduced: PHP80,
		},
		{
			name:       "constructor promotion",
			source:     "<?php class C { function __construct(public string $name) {} }",
			before:     PHP74,
			introduced: PHP80,
		},
		{
			name:       "enum",
			source:     "<?php enum Status { case Active; }",
			before:     PHP80,
			introduced: PHP81,
		},
		{
			name:       "intersection type",
			source:     "<?php function handle(Countable&Iterator $value): void {}",
			before:     PHP80,
			introduced: PHP81,
		},
		{
			name:       "first class callable",
			source:     "<?php $callable = strlen(...);",
			before:     PHP80,
			introduced: PHP81,
		},
		{
			name:       "readonly property",
			source:     "<?php class Value { public readonly int $number; }",
			before:     PHP80,
			introduced: PHP81,
		},
		{
			name:       "new in initializer",
			source:     "<?php function make($value = new stdClass()) {}",
			before:     PHP80,
			introduced: PHP81,
		},
		{
			name:       "new in initializer with arguments",
			source:     "<?php class C { public function __construct(protected Config $config = new Config(2, 0.01, true)) {} }",
			before:     PHP80,
			introduced: PHP81,
		},
		{
			name:       "readonly class",
			source:     "<?php readonly class Value { public int $number; }",
			before:     PHP81,
			introduced: PHP82,
		},
		{
			name:       "dnf type",
			source:     "<?php function handle((Countable&Iterator)|Stringable $value): void {}",
			before:     PHP81,
			introduced: PHP82,
		},
		{
			name:       "true type",
			source:     "<?php function yes(): true { return true; }",
			before:     PHP81,
			introduced: PHP82,
		},
		{
			name:       "standalone null type",
			source:     "<?php function nothing(): null { return null; }",
			before:     PHP81,
			introduced: PHP82,
		},
		{
			name:       "standalone nullable false type",
			source:     "<?php function no(): ?false { return false; }",
			before:     PHP81,
			introduced: PHP82,
		},
		{
			name:       "union of only null and false",
			source:     "<?php function maybe(): false|null { return false; }",
			before:     PHP81,
			introduced: PHP82,
		},
		{
			name:       "union with null member",
			source:     "<?php function pick(string|callable|null $arrow = null): ?array { return null; }",
			before:     PHP74,
			introduced: PHP80,
		},
		{
			name:       "trait constant",
			source:     "<?php trait Values { public const ANSWER = 42; }",
			before:     PHP81,
			introduced: PHP82,
		},
		{
			name:       "typed class constant",
			source:     "<?php class C { public const string NAME = 'c'; }",
			before:     PHP82,
			introduced: PHP83,
		},
		{
			name:       "arbitrary static initializer",
			source:     "<?php function value() { static $item = strlen('value'); return $item; }",
			before:     PHP82,
			introduced: PHP83,
		},
		{
			name:       "dynamic class constant",
			source:     "<?php class C { const NAME = 'c'; } $name = 'NAME'; echo C::{$name};",
			before:     PHP82,
			introduced: PHP83,
		},
		{
			name:       "readonly anonymous class",
			source:     "<?php $value = new readonly class { public function x() {} };",
			before:     PHP82,
			introduced: PHP83,
		},
		{
			name:       "property hooks",
			source:     "<?php class C { public string $name { get => $this->name; set => $value; } }",
			before:     PHP83,
			introduced: PHP84,
		},
		{
			name:       "asymmetric visibility",
			source:     "<?php class C { public private(set) string $name; }",
			before:     PHP83,
			introduced: PHP84,
		},
		{
			name:       "new dereference",
			source:     "<?php class C { function x() {} } new C()->x();",
			before:     PHP83,
			introduced: PHP84,
		},
		{
			name:       "pipe",
			source:     "<?php $length = ' hello ' |> trim(...) |> strlen(...);",
			before:     PHP84,
			introduced: PHP85,
		},
		{
			name:       "clone with",
			source:     "<?php $copy = clone($value, ['name' => 'new']);",
			before:     PHP84,
			introduced: PHP85,
		},
		{
			name:       "void cast",
			source:     "<?php (void) do_work();",
			before:     PHP84,
			introduced: PHP85,
		},
		{
			name:       "constant attribute",
			source:     "<?php #[Deprecated] const OLD_VALUE = 1;",
			before:     PHP84,
			introduced: PHP85,
		},
		{
			name:       "final promoted property",
			source:     "<?php class C { function __construct(final public string $name) {} }",
			before:     PHP84,
			introduced: PHP85,
		},
		{
			name:       "override class constant",
			source:     "<?php class A { const VALUE = 1; } class C extends A { #[\\Override] const VALUE = 2; }",
			before:     PHP85,
			introduced: PHP86,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			before, err := Lint(
				"before.php",
				[]byte(test.source),
				Options{PHPVersion: test.before},
			)
			if err != nil {
				t.Fatalf("Lint(before) error = %v", err)
			}
			if len(before) == 0 {
				t.Fatalf("Lint(before PHP %s) unexpectedly passed", test.before)
			}

			after, err := Lint(
				"after.php",
				[]byte(test.source),
				Options{PHPVersion: test.introduced},
			)
			if err != nil {
				t.Fatalf("Lint(introduced) error = %v", err)
			}
			if len(after) != 0 {
				t.Fatalf("Lint(PHP %s) diagnostics = %#v, want none", test.introduced, after)
			}
		})
	}
}

func TestModernCompileValidation(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		source     string
		wantPhrase string
	}{
		{
			name:   "reference get hook",
			source: "<?php class C { public string $name { &get => $this->name; } }",
		},
		{
			name:   "attributed final hook",
			source: "<?php class C { public string $name { #[Deprecated] final get => $this->name; } }",
		},
		{
			name:   "typed set hook",
			source: "<?php class C { public string $name { set(string $other) => $other; } }",
		},
		{
			name:   "method body beginning with get",
			source: "<?php class C { function run() { get(); } }",
		},
		{
			name:       "static hooked property",
			source:     "<?php class C { public static string $name { get => $this->name; } }",
			wantPhrase: "cannot be static",
		},
		{
			name:       "duplicate hook",
			source:     "<?php class C { public string $name { get => $this->name; get => ''; } }",
			wantPhrase: "cannot redeclare",
		},
		{
			name:       "get parameters",
			source:     "<?php class C { public string $name { get($value) => $value; } }",
			wantPhrase: "must not have a parameter",
		},
		{
			name:       "invalid set parameters",
			source:     "<?php class C { public string $name { set($a, $b) => $a; } }",
			wantPhrase: "exactly one",
		},
		{
			name:       "bodyless concrete hook",
			source:     "<?php class C { public string $name { get; } }",
			wantPhrase: "must have a body",
		},
		{
			name:       "weaker property visibility",
			source:     "<?php class C { private public(set) string $name; }",
			wantPhrase: "visibility",
		},
		{
			name:       "invalid typed constant type",
			source:     "<?php class C { const callable HANDLER = strlen(...); }",
			wantPhrase: "cannot have type callable",
		},
		{
			name:       "typed constant value mismatch",
			source:     "<?php class C { const int VALUE = 'wrong'; }",
			wantPhrase: "not compatible",
		},
		{
			name:   "typed constant union",
			source: "<?php class C { const int|string VALUE = 'right'; }",
		},
		{
			name:       "second typed constant mismatch",
			source:     "<?php class C { const int FIRST = 1, SECOND = 'wrong'; }",
			wantPhrase: "not compatible",
		},
		{
			name:       "readonly method",
			source:     "<?php class C { readonly function run() {} }",
			wantPhrase: "cannot be readonly",
		},
		{
			name:       "static class constant",
			source:     "<?php class C { static const VALUE = 1; }",
			wantPhrase: "cannot be static",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			diagnostics, err := Lint(
				"modern.php",
				[]byte(test.source),
				Options{PHPVersion: PHP85},
			)
			if err != nil {
				t.Fatalf("Lint() error = %v", err)
			}
			if test.wantPhrase == "" {
				if len(diagnostics) != 0 {
					t.Fatalf("Lint() diagnostics = %#v, want none", diagnostics)
				}
				return
			}
			if !diagnosticsContain(diagnostics, test.wantPhrase) {
				t.Fatalf("Lint() diagnostics = %#v, want phrase %q", diagnostics, test.wantPhrase)
			}
		})
	}
}

func TestSemiReservedMethodNamesAreNotOperators(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		source string
	}{
		{
			name:   "method declaration named clone",
			source: "<?php class C { public function clone(string $id, string $other, ?string $more = null): string { return $id . $other . $more; } }",
		},
		{
			name:   "static call to method named clone",
			source: "<?php class C { public static function clone(string $a, string $b): string { return $a . $b; } } echo C::clone('a', 'b');",
		},
		{
			name:   "dereferenced static call to method named new",
			source: "<?php Builder::new()->qux();",
		},
		{
			name:   "dereferenced static call to method named new with arguments",
			source: "<?php Builder::new('a', 'b')->qux();",
		},
		{
			name:   "dereferenced instance call to method named new",
			source: "<?php $builder->new()->qux();",
		},
		{
			name:   "method declaration named new",
			source: "<?php class C { public static function new(): self { return new self(); } }",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			for _, version := range []Version{PHP72, PHP84, PHP85} {
				diagnostics, err := Lint(
					"clone.php",
					[]byte(test.source),
					Options{PHPVersion: version},
				)
				if err != nil {
					t.Fatalf("Lint(PHP %s) error = %v", version, err)
				}
				if len(diagnostics) != 0 {
					t.Fatalf("Lint(PHP %s) diagnostics = %#v, want none", version, diagnostics)
				}
			}
		})
	}
}

func diagnosticsContain(diagnostics []Diagnostic, phrase string) bool {
	for _, diagnostic := range diagnostics {
		if strings.Contains(strings.ToLower(diagnostic.Message), strings.ToLower(phrase)) {
			return true
		}
	}
	return false
}
