package php_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/shyim/go-phplint/internal/conf"
	php "github.com/shyim/go-phplint/internal/php"
	"github.com/shyim/go-phplint/internal/version"
)

func BenchmarkPhp(b *testing.B) {
	src, err := os.ReadFile(filepath.Join("testdata", "test.php"))
	if err != nil {
		b.Fatal("can not read testdata/test.php: " + err.Error())
	}

	for n := 0; n < b.N; n++ {
		config := conf.Config{
			Version: &version.Version{
				Major: 8,
				Minor: 4,
			},
		}
		lexer := php.NewLexer(src, config)
		phpparser := php.NewParser(lexer, config)
		phpparser.Parse()
	}
}
