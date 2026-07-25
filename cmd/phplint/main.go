package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"

	"github.com/shyim/go-phplint"
)

var buildVersion = "dev"

type fileResult struct {
	path        string
	diagnostics []phplint.Diagnostic
	err         error
}

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout, stderr io.Writer) int {
	flags := flag.NewFlagSet("phplint", flag.ContinueOnError)
	flags.SetOutput(stderr)

	var phpVersionValue string
	var showVersion bool
	flags.StringVar(&phpVersionValue, "php-version", "", "PHP minor version: 7.2 through 7.4 or 8.0 through 8.6 (8.6 is preview)")
	flags.BoolVar(&showVersion, "version", false, "print the phplint version")
	flags.Usage = func() {
		_, _ = fmt.Fprintln(stderr, "Usage: phplint --php-version <version> <file.php> [more.php ...]")
		flags.PrintDefaults()
	}

	if err := flags.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return 0
		}
		return 2
	}
	if showVersion {
		_, _ = fmt.Fprintf(stdout, "phplint %s\n", buildVersion)
		return 0
	}
	if phpVersionValue == "" {
		_, _ = fmt.Fprintln(stderr, "phplint: --php-version is required")
		flags.Usage()
		return 2
	}

	phpVersion, err := phplint.ParseVersion(phpVersionValue)
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "phplint: %v\n", err)
		return 2
	}

	paths := flags.Args()
	if len(paths) == 0 {
		_, _ = fmt.Fprintln(stderr, "phplint: at least one file is required")
		flags.Usage()
		return 2
	}

	results := lintFiles(paths, phpVersion)
	errorCount := 0
	failedFiles := 0
	operationalFailure := false

	for _, result := range results {
		if result.err != nil {
			_, _ = fmt.Fprintf(stderr, "phplint: %s: %v\n", result.path, result.err)
			operationalFailure = true
			continue
		}
		if len(result.diagnostics) == 0 {
			continue
		}

		failedFiles++
		errorCount += len(result.diagnostics)
		for _, diagnostic := range result.diagnostics {
			_, _ = fmt.Fprintf(
				stdout,
				"%s:%d:%d: PHP %s: %s\n",
				diagnostic.Filename,
				diagnostic.Start.Line,
				diagnostic.Start.Column,
				phpVersion,
				diagnostic.Message,
			)
		}
	}

	if errorCount > 0 {
		_, _ = fmt.Fprintf(stdout, "%d error(s) in %d file(s)\n", errorCount, failedFiles)
	}
	if operationalFailure {
		return 2
	}
	if errorCount > 0 {
		return 1
	}
	return 0
}

func lintFiles(paths []string, version phplint.Version) []fileResult {
	results := make([]fileResult, len(paths))
	jobs := runtime.GOMAXPROCS(0)
	if jobs > len(paths) {
		jobs = len(paths)
	}

	work := make(chan int)
	var workers sync.WaitGroup
	workers.Add(jobs)
	for range jobs {
		go func() {
			defer workers.Done()
			for index := range work {
				results[index] = lintFile(paths[index], version)
			}
		}()
	}

	for index := range paths {
		work <- index
	}
	close(work)
	workers.Wait()
	return results
}

func lintFile(path string, version phplint.Version) fileResult {
	result := fileResult{path: path}

	info, err := os.Stat(path)
	if err != nil {
		result.err = err
		return result
	}
	if info.IsDir() {
		result.err = errors.New("directories are not supported")
		return result
	}
	if !info.Mode().IsRegular() {
		result.err = errors.New("not a regular file")
		return result
	}

	source, err := os.ReadFile(path)
	if err != nil {
		result.err = err
		return result
	}

	result.diagnostics, result.err = phplint.Lint(
		path,
		source,
		phplint.Options{PHPVersion: version},
	)
	return result
}
