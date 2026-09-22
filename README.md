# go-phplint

[![CodSpeed](https://img.shields.io/endpoint?url=https://codspeed.io/badge.json)](https://app.codspeed.io/shyim/go-phplint?utm_source=badge)

`phplint` is a source-only PHP syntax and compile-time linter written in Go. It
selects the PHP language profile explicitly, so one executable can check code
for PHP 7.2 through 7.4 or any PHP 8 minor from 8.0 through 8.6 without
installing those PHP runtimes. PHP 8.6 support is currently a preview profile.

The normal build is pure Go. A single unified PHP lexer/parser in
`internal/php` accepts the full PHP 7/8 syntax superset for every supported
profile, and the linter adds version gates and recoverable compile-time
checks on top. Version-specific rejection lives in the lexer (reserved
words, removed casts, heredoc rules, `#` comments), in a few
version-conditional grammar actions, and in validation.

## Build

Go 1.24 or newer is required.

```sh
go build -trimpath -o phplint ./cmd/phplint
```

To stamp a release version:

```sh
go build -trimpath -ldflags "-X main.buildVersion=v0.1.0" -o phplint ./cmd/phplint
```

## CLI

Pass exactly one supported PHP minor and one or more files:

```sh
phplint --php-version 7.2 oldest-supported.php
phplint --php-version 7.4 legacy.php
phplint --php-version 8.4 src/App.php src/Domain.php
phplint --php-version 8.6 preview.php
```

A successful run is silent. A source failure is concise:

```text
src/App.php:12:9: PHP 8.4: syntax error
    function broken( {
            ^
1 error(s) in 1 file(s)
```

The CLI accepts explicit regular files only. It deliberately has no directory
walking, configuration discovery, standard-input mode, or implicit version
default. PHP short tags are always recognized.

Exit codes:

- `0`: every file passed
- `1`: one or more source diagnostics
- `2`: invalid arguments, file I/O failure, or internal failure

## Go API

```go
package main

import (
	"fmt"

	phplint "github.com/shyim/go-phplint"
)

func main() {
	diagnostics, err := phplint.Lint(
		"example.php",
		[]byte("<?php function answer(): int { return 42; }"),
		phplint.Options{PHPVersion: phplint.PHP84},
	)
	if err != nil {
		panic(err)
	}
	for _, diagnostic := range diagnostics {
		fmt.Println(diagnostic)
	}
}
```

`Lint` returns source failures as diagnostics and reserves `error` for invalid
options or an internal failure. Diagnostic offsets are zero-based byte offsets;
line and column values are one-based. Each diagnostic also carries
`SourceLine`, the text of the offending source line, so callers can render
excerpts without re-reading the file.

## Compatibility scope

The stable profiles are exactly `7.2`, `7.3`, `7.4`, `8.0`, `8.1`, `8.2`,
`8.3`, `8.4`, and `8.5`. The `8.6` profile follows PHP 8.6.0 Alpha 2 and
remains a preview until PHP 8.6 reaches general availability. Patch versions
such as `8.4.2` are rejected because syntax profiles are maintained at
minor-version granularity.

The goal is native `php -l` pass/fail behavior for syntax errors and
single-file compile-time fatal errors, not identical diagnostic wording.
Warnings and deprecations for which `php -l` exits successfully are not
reported. Tests compare the linter with native PHP binaries across the full
version matrix. PHP itself remains the final authority for edge cases that
depend on engine semantics rather than parsing or single-file compilation.

## Development

```sh
go test ./...
go vet . ./cmd/phplint
```

If a supported `php` executable is on `PATH`, `go test` also runs the native
pass/fail oracle. Set `PHPLINT_PHP_BINARY` to select another executable:

```sh
PHPLINT_PHP_BINARY=/opt/php/8.4/bin/php go test -run TestNativePHPOracle
```

The CI oracle builds one Go test binary and runs it in official PHP CLI
containers for every supported minor.

The generated parser (`internal/php/php.go`) and scanner
(`internal/php/scanner.go`) are checked in. To regenerate them after editing
the grammars, install `goyacc` and `ragel` and run:

```sh
goyacc -o internal/php/php.go internal/php/php.y
ragel -Z -G2 -o internal/php/scanner.go internal/php/scanner.rl
```

`goyacc` is `golang.org/x/tools/cmd/goyacc`. After it runs, set
`yyInitialStackSize` to 128 and `yyErrorVerbose` to true in
`internal/php/php.go`. Ragel must be invoked with `-Z` so the host language
is Go. Keywords are classified in Go after an identifier match, not as
case-folded literals in `scanner.rl`.

Hand edits to the generated files must be mirrored in `php.y`/`scanner.rl`
so regeneration preserves them. The version-conditional actions, the parser
stack size, and verbose syntax errors are the intentional divergence from a
plain regeneration.

## Benchmarks

Benchmarks live next to the code they measure and use the standard `testing`
package. They cover the public `Lint` entry point, the source preparation,
parsing and validation stages, and the concurrent CLI pipeline:

```sh
go test -bench=. . ./cmd/phplint
```

Every push and pull request runs them on
[CodSpeed](https://app.codspeed.io/shyim/go-phplint) with the walltime
instrument, so performance changes show up in the pull request.

## Origin and license

The embedded parser is adapted from
[`laytan/php-parser` v0.10.0](https://github.com/laytan/php-parser), which is
MIT-licensed. Its original license is retained in
[`internal/LICENSE`](internal/LICENSE). This project is also released
under the MIT License.
