package php8_test

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/shyim/go-phplint/internal/php8"
	"github.com/shyim/go-phplint/internal/conf"
	"github.com/shyim/go-phplint/internal/version"
)

func BenchmarkPhp8(b *testing.B) {
	src, err := ioutil.ReadFile(filepath.Join("testdata", "test.php"))
	if err != nil {
		b.Fatal("can not read testdata/test.php: " + err.Error())
	}

	for n := 0; n < b.N; n++ {
		config := conf.Config{
			Version: &version.Version{
				Major: 8,
				Minor: 8,
			},
		}
		lexer := php8.NewLexer(src, config)
		php8parser := php8.NewParser(lexer, config)
		php8parser.Parse()
	}
}
