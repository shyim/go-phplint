package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunRequiresVersionAndFiles(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		args []string
	}{
		{name: "version", args: []string{"test.php"}},
		{name: "file", args: []string{"--php-version", "8.5"}},
		{name: "supported version", args: []string{"--php-version", "8.6", "test.php"}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			var stdout, stderr bytes.Buffer
			if exitCode := run(test.args, &stdout, &stderr); exitCode != 2 {
				t.Fatalf("run() = %d, want 2", exitCode)
			}
			if stderr.Len() == 0 {
				t.Fatal("run() wrote no usage error")
			}
		})
	}
}

func TestRunIsSilentForCleanFiles(t *testing.T) {
	t.Parallel()

	path := writeFixture(t, "clean.php", "<?php echo 'clean';")
	var stdout, stderr bytes.Buffer
	exitCode := run([]string{"--php-version", "8.5", path}, &stdout, &stderr)

	if exitCode != 0 {
		t.Fatalf("run() = %d, want 0; stderr = %q", exitCode, stderr.String())
	}
	if stdout.Len() != 0 || stderr.Len() != 0 {
		t.Fatalf("clean run output stdout=%q stderr=%q, want none", stdout.String(), stderr.String())
	}
}

func TestRunPrintsDeterministicDiagnostics(t *testing.T) {
	t.Parallel()

	first := writeFixture(t, "first.php", "<?php break;")
	second := writeFixture(t, "second.php", "<?php continue;")

	var stdout, stderr bytes.Buffer
	exitCode := run(
		[]string{"--php-version", "8.5", first, second},
		&stdout,
		&stderr,
	)
	if exitCode != 1 {
		t.Fatalf("run() = %d, want 1; stderr = %q", exitCode, stderr.String())
	}
	if stderr.Len() != 0 {
		t.Fatalf("run() stderr = %q, want none", stderr.String())
	}

	output := stdout.String()
	firstIndex := strings.Index(output, first+":")
	secondIndex := strings.Index(output, second+":")
	if firstIndex < 0 || secondIndex < 0 || firstIndex >= secondIndex {
		t.Fatalf("diagnostics are not in input order: %q", output)
	}
	if !strings.HasSuffix(output, "2 error(s) in 2 file(s)\n") {
		t.Fatalf("run() summary = %q", output)
	}
}

func TestRunRejectsDirectories(t *testing.T) {
	t.Parallel()

	var stdout, stderr bytes.Buffer
	exitCode := run(
		[]string{"--php-version", "8.5", t.TempDir()},
		&stdout,
		&stderr,
	)
	if exitCode != 2 {
		t.Fatalf("run() = %d, want 2", exitCode)
	}
	if !strings.Contains(stderr.String(), "directories are not supported") {
		t.Fatalf("run() stderr = %q", stderr.String())
	}
}

func writeFixture(t *testing.T, name, source string) string {
	t.Helper()

	path := filepath.Join(t.TempDir(), name)
	if err := os.WriteFile(path, []byte(source), 0o600); err != nil {
		t.Fatalf("write fixture: %v", err)
	}
	return path
}
