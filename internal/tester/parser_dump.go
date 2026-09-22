package tester

import (
	"bytes"
	"testing"

	"github.com/shyim/go-phplint/internal/conf"
	"github.com/shyim/go-phplint/internal/parser"
	"github.com/shyim/go-phplint/internal/version"
	"github.com/shyim/go-phplint/internal/visitor/dumper"
	"gotest.tools/assert"
)

type ParserDumpTestSuite struct {
	t *testing.T

	Code     string
	Expected string

	Version version.Version

	actualDump *bytes.Buffer
	dumper     *dumper.Dumper
}

func NewParserDumpTestSuite(t *testing.T) *ParserDumpTestSuite {
	actualDump := bytes.NewBuffer(nil)
	return &ParserDumpTestSuite{
		t: t,
		Version: version.Version{
			Major: 7,
			Minor: 4,
		},
		actualDump: actualDump,
		dumper:     dumper.NewDumper(actualDump),
	}
}

func (p *ParserDumpTestSuite) WithTokens() {
	p.dumper = p.dumper.WithTokens()
}

func (p *ParserDumpTestSuite) WithPositions() {
	p.dumper = p.dumper.WithPositions()
}

func (p *ParserDumpTestSuite) UsePHP8() {
	p.Version = version.Version{Major: 8, Minor: 0}
}

func (p *ParserDumpTestSuite) Run() {
	p.t.Helper()

	config := conf.Config{
		Fidelity: true,
		Version:  &p.Version,
	}

	actual, err := parser.Parse([]byte(p.Code), config)
	if err != nil {
		p.t.Fatalf("Error parse: %v", err)
	}

	p.dumper.Dump(actual)

	assert.DeepEqual(p.t, p.Expected+"\n", p.actualDump.String())
}
