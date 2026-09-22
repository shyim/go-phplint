package token_test

import (
	"os"
	"regexp"
	"testing"

	"github.com/shyim/go-phplint/internal/token"
)

func TestTokenIDsMatchGrammarOrder(t *testing.T) {
	grammar, err := os.ReadFile("../php/php.y")
	if err != nil {
		t.Fatal(err)
	}
	source, err := os.ReadFile("token.go")
	if err != nil {
		t.Fatal(err)
	}

	grammarNames := regexp.MustCompile(`%token <token> (T_[A-Z0-9_]+)`).FindAllSubmatch(grammar, -1)
	sourceNames := regexp.MustCompile(`(?m)^\t(T_[A-Z0-9_]+)\b`).FindAllSubmatch(source, -1)
	if len(grammarNames) == 0 || len(grammarNames) != len(sourceNames) {
		t.Fatalf("grammar tokens %d, source tokens %d", len(grammarNames), len(sourceNames))
	}
	for i := range grammarNames {
		if string(grammarNames[i][1]) != string(sourceNames[i][1]) {
			t.Fatalf("token %d grammar %s source %s", i, grammarNames[i][1], sourceNames[i][1])
		}
	}

	if token.T_INCLUDE != 57346 {
		t.Fatalf("T_INCLUDE = %d, want goyacc's first private token 57346", token.T_INCLUDE)
	}
	if token.T_PIPE != token.T_AMPERSAND_NOT_FOLLOWED_BY_VAR_OR_VARARG+1 {
		t.Fatal("T_PIPE is not the token after the previous last identifier")
	}
	if token.T_VOID_CAST != token.T_PIPE+1 {
		t.Fatal("T_VOID_CAST does not follow T_PIPE")
	}
	if token.T_PUBLIC_SET != token.T_VOID_CAST+1 ||
		token.T_PROTECTED_SET != token.T_PUBLIC_SET+1 ||
		token.T_PRIVATE_SET != token.T_PROTECTED_SET+1 {
		t.Fatal("asymmetric visibility tokens are not at the end of the identifier list")
	}
	if token.T_PROPERTY_HOOKS != token.T_PRIVATE_SET+1 {
		t.Fatal("T_PROPERTY_HOOKS does not follow the asymmetric visibility tokens")
	}
}
