package main

import (
	"io"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	phplint "github.com/shyim/go-phplint"
)

// benchCorpus writes count copies of the repository benchmark fixtures into a
// temporary directory and returns their paths.
func benchCorpus(b *testing.B, count int) []string {
	b.Helper()

	fixtures := []string{"legacy.php", "modern.php", "template.php", "invalid.php"}
	sources := make([][]byte, len(fixtures))
	for index, fixture := range fixtures {
		source, err := os.ReadFile(filepath.Join("..", "..", "testdata", "bench", fixture))
		if err != nil {
			b.Fatalf("read benchmark fixture %s: %v", fixture, err)
		}
		sources[index] = source
	}

	directory := b.TempDir()
	paths := make([]string, 0, count)
	for index := range count {
		source := sources[index%len(sources)]
		name := strconv.Itoa(index) + "-" + fixtures[index%len(fixtures)]
		path := filepath.Join(directory, name)
		if err := os.WriteFile(path, source, 0o600); err != nil {
			b.Fatalf("write benchmark fixture: %v", err)
		}
		paths = append(paths, path)
	}
	return paths
}

// BenchmarkLintFiles measures the concurrent file pipeline used by the CLI,
// including file I/O and worker fan-out.
func BenchmarkLintFiles(b *testing.B) {
	paths := benchCorpus(b, 24)

	for b.Loop() {
		results := lintFiles(paths, phplint.PHP84)
		if len(results) != len(paths) {
			b.Fatalf("lintFiles() returned %d results, want %d", len(results), len(paths))
		}
	}
}

// BenchmarkRunCLI measures a full CLI invocation, including argument parsing
// and diagnostic rendering.
func BenchmarkRunCLI(b *testing.B) {
	paths := benchCorpus(b, 8)
	args := append([]string{"--php-version", "8.4"}, paths...)

	for b.Loop() {
		if exitCode := run(args, io.Discard, io.Discard); exitCode != 1 {
			b.Fatalf("run() = %d, want 1", exitCode)
		}
	}
}
