package php7_test

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/shyim/phplint-go/internal/php/internal/php7"
	"github.com/shyim/phplint-go/internal/php/pkg/conf"
	"github.com/shyim/phplint-go/internal/php/pkg/version"
)

func BenchmarkPhp7(b *testing.B) {
	src, err := ioutil.ReadFile(filepath.Join("testdata", "test.php"))
	if err != nil {
		b.Fatal("can not read testdata/test.php: " + err.Error())
	}

	for n := 0; n < b.N; n++ {
		config := conf.Config{
			Version: &version.Version{
				Major: 7,
				Minor: 4,
			},
		}
		lexer := php7.NewLexer(src, config)
		php7parser := php7.NewParser(lexer, config)
		php7parser.Parse()
	}
}
