package php

import "github.com/shyim/go-phplint/internal/token"

// literalFitsInt reports whether the integer literal digits fit in a Go int
// after underscores are ignored. Invalid digits fail the same way strconv.ParseInt does.
func literalFitsInt(src []byte, base int) bool {
	if base < 2 || base > 36 {
		return false
	}
	maxInt := uint64(^uint(0) >> 1)
	var n uint64
	digits := 0
	for _, c := range src {
		if c == '_' {
			continue
		}
		var d byte
		switch {
		case c >= '0' && c <= '9':
			d = c - '0'
		case c >= 'a' && c <= 'z':
			d = c - 'a' + 10
		case c >= 'A' && c <= 'Z':
			d = c - 'A' + 10
		default:
			return false
		}
		if int(d) >= base {
			return false
		}
		digits++
		if n > (maxInt-uint64(d))/uint64(base) {
			return false
		}
		n = n*uint64(base) + uint64(d)
	}
	return digits > 0
}

func (lex *Lexer) integerToken(prefix, base int) token.ID {
	if literalFitsInt(lex.data[lex.ts+prefix:lex.te], base) {
		return token.T_LNUMBER
	}
	return token.T_DNUMBER
}

type keyword struct {
	text         string
	id           token.ID
	gated        bool
	major, minor uint64
}

func (lex *Lexer) identifierToken() token.ID {
	word := lex.data[lex.ts:lex.te]
	for _, kw := range keywordsByLen[len(word)] {
		if !equalFoldASCII(word, kw.text) {
			continue
		}
		if kw.gated && !lex.versionAtLeast(kw.major, kw.minor) {
			return token.T_STRING
		}
		if kw.id == token.T_PUBLIC || kw.id == token.T_PROTECTED || kw.id == token.T_PRIVATE {
			if lex.consumeSetVisibility() {
				switch kw.id {
				case token.T_PROTECTED:
					return token.T_PROTECTED_SET
				case token.T_PRIVATE:
					return token.T_PRIVATE_SET
				default:
					return token.T_PUBLIC_SET
				}
			}
		}
		return kw.id
	}
	return token.T_STRING
}

// consumeSetVisibility folds a following "(set)" into this keyword and reports
// whether it did. PHP accepts no whitespace inside public(set), protected(set),
// or private(set). Any other "(" stays in the token stream so the type grammar
// can see it. The caller has p at the last byte of the keyword and increments
// p once after return. The resulting token is not an identifier.
func (lex *Lexer) consumeSetVisibility() bool {
	end := lex.te
	if end+5 > lex.pe || lex.data[end] != '(' {
		return false
	}
	if !equalFoldASCII(lex.data[end+1:end+4], "set") || lex.data[end+4] != ')' {
		return false
	}
	lex.te = end + 5
	lex.p = lex.te - 1
	return true
}

// looksLikePropertyHookBlock reports whether the "{" just recognized opens a
// property hook block. Function and control-structure bodies that call get()
// or set() stay ordinary braces. te is one past the brace.
func (lex *Lexer) looksLikePropertyHookBlock() bool {
	i := lex.skipHookTrivia(lex.te)
	for {
		next := lex.skipHookAttribute(i)
		if next == i {
			break
		}
		i = lex.skipHookTrivia(next)
	}
	if lex.hasHookWord(i, "final") {
		i = lex.skipHookTrivia(i + len("final"))
	}
	if i < lex.pe && lex.data[i] == '&' {
		i = lex.skipHookTrivia(i + 1)
	}
	word, next := lex.hookWord(i)
	if word != "get" && word != "set" {
		return false
	}
	next = lex.skipHookTrivia(next)
	if next >= lex.pe || lex.data[next] != '(' {
		return true
	}
	after := lex.skipHookParens(next)
	if after < 0 {
		return false
	}
	after = lex.skipHookTrivia(after)
	return lex.hasHookPrefix(after, "=>") || (after < lex.pe && lex.data[after] == '{')
}

func (lex *Lexer) skipHookTrivia(i int) int {
	for i < lex.pe {
		switch lex.data[i] {
		case ' ', '\t', '\n', '\r', '\v', '\f':
			i++
		case '#':
			if i+1 < lex.pe && lex.data[i+1] == '[' {
				return i
			}
			i++
			for i < lex.pe && lex.data[i] != '\n' {
				i++
			}
		case '/':
			if i+1 >= lex.pe {
				return i
			}
			if lex.data[i+1] == '/' {
				i += 2
				for i < lex.pe && lex.data[i] != '\n' {
					i++
				}
				continue
			}
			if lex.data[i+1] == '*' {
				i += 2
				for i+1 < lex.pe && !(lex.data[i] == '*' && lex.data[i+1] == '/') {
					i++
				}
				if i+1 < lex.pe {
					i += 2
				}
				continue
			}
			return i
		default:
			return i
		}
	}
	return i
}

func (lex *Lexer) skipHookAttribute(i int) int {
	if i+1 >= lex.pe || lex.data[i] != '#' || lex.data[i+1] != '[' {
		return i
	}
	i += 2
	depth := 1
	for i < lex.pe && depth > 0 {
		switch lex.data[i] {
		case '\'', '"':
			i = lex.skipHookString(i)
		case '[':
			depth++
			i++
		case ']':
			depth--
			i++
		default:
			i++
		}
	}
	return i
}

func (lex *Lexer) skipHookString(i int) int {
	quote := lex.data[i]
	i++
	for i < lex.pe {
		if lex.data[i] == '\\' && i+1 < lex.pe {
			i += 2
			continue
		}
		if lex.data[i] == quote {
			return i + 1
		}
		i++
	}
	return i
}

func (lex *Lexer) skipHookParens(i int) int {
	if i >= lex.pe || lex.data[i] != '(' {
		return -1
	}
	i++
	depth := 1
	for i < lex.pe && depth > 0 {
		switch lex.data[i] {
		case '\'', '"':
			i = lex.skipHookString(i)
		case '(':
			depth++
			i++
		case ')':
			depth--
			i++
		default:
			i++
		}
	}
	if depth != 0 {
		return -1
	}
	return i
}

func (lex *Lexer) hookWord(i int) (string, int) {
	if i >= lex.pe || !isHookIdentStart(lex.data[i]) {
		return "", i
	}
	start := i
	i++
	for i < lex.pe && isHookIdentPart(lex.data[i]) {
		i++
	}
	word := lex.data[start:i]
	lower := make([]byte, len(word))
	for n, c := range word {
		if c >= 'A' && c <= 'Z' {
			c += 'a' - 'A'
		}
		lower[n] = c
	}
	return string(lower), i
}

func (lex *Lexer) hasHookWord(i int, word string) bool {
	got, next := lex.hookWord(i)
	return got == word && (next >= lex.pe || !isHookIdentPart(lex.data[next]))
}

func (lex *Lexer) hasHookPrefix(i int, prefix string) bool {
	if i+len(prefix) > lex.pe {
		return false
	}
	return string(lex.data[i:i+len(prefix)]) == prefix
}

func isHookIdentStart(c byte) bool {
	return c == '_' || (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')
}

func isHookIdentPart(c byte) bool {
	return isHookIdentStart(c) || (c >= '0' && c <= '9')
}

func equalFoldASCII(value []byte, word string) bool {
	if len(value) != len(word) {
		return false
	}
	for i := 0; i < len(value); i++ {
		c := value[i]
		if c >= 'A' && c <= 'Z' {
			c += 'a' - 'A'
		}
		if c != word[i] {
			return false
		}
	}
	return true
}

var keywordsByLen = map[int][]keyword{
	2: {
		{"as", token.T_AS, false, 0, 0},
		{"do", token.T_DO, false, 0, 0},
		{"fn", token.T_FN, true, 7, 4},
		{"if", token.T_IF, false, 0, 0},
		{"or", token.T_LOGICAL_OR, false, 0, 0},
	},
	3: {
		{"and", token.T_LOGICAL_AND, false, 0, 0},
		{"die", token.T_EXIT, false, 0, 0},
		{"for", token.T_FOR, false, 0, 0},
		{"new", token.T_NEW, false, 0, 0},
		{"try", token.T_TRY, false, 0, 0},
		{"use", token.T_USE, false, 0, 0},
		{"var", token.T_VAR, false, 0, 0},
		{"xor", token.T_LOGICAL_XOR, false, 0, 0},
	},
	4: {
		{"case", token.T_CASE, false, 0, 0},
		{"echo", token.T_ECHO, false, 0, 0},
		{"else", token.T_ELSE, false, 0, 0},
		{"enum", token.T_ENUM, true, 8, 0},
		{"exit", token.T_EXIT, false, 0, 0},
		{"eval", token.T_EVAL, false, 0, 0},
		{"goto", token.T_GOTO, false, 0, 0},
		{"list", token.T_LIST, false, 0, 0},
	},
	5: {
		{"array", token.T_ARRAY, false, 0, 0},
		{"break", token.T_BREAK, false, 0, 0},
		{"catch", token.T_CATCH, false, 0, 0},
		{"class", token.T_CLASS, false, 0, 0},
		{"clone", token.T_CLONE, false, 0, 0},
		{"const", token.T_CONST, false, 0, 0},
		{"empty", token.T_EMPTY, false, 0, 0},
		{"endif", token.T_ENDIF, false, 0, 0},
		{"final", token.T_FINAL, false, 0, 0},
		{"isset", token.T_ISSET, false, 0, 0},
		{"match", token.T_MATCH, true, 8, 0},
		{"print", token.T_PRINT, false, 0, 0},
		{"throw", token.T_THROW, false, 0, 0},
		{"trait", token.T_TRAIT, false, 0, 0},
		{"unset", token.T_UNSET, false, 0, 0},
		{"while", token.T_WHILE, false, 0, 0},
		{"yield", token.T_YIELD, false, 0, 0},
	},
	6: {
		{"elseif", token.T_ELSEIF, false, 0, 0},
		{"endfor", token.T_ENDFOR, false, 0, 0},
		{"public", token.T_PUBLIC, false, 0, 0},
		{"return", token.T_RETURN, false, 0, 0},
		{"static", token.T_STATIC, false, 0, 0},
		{"switch", token.T_SWITCH, false, 0, 0},
		{"global", token.T_GLOBAL, false, 0, 0},
	},
	7: {
		{"declare", token.T_DECLARE, false, 0, 0},
		{"default", token.T_DEFAULT, false, 0, 0},
		{"extends", token.T_EXTENDS, false, 0, 0},
		{"finally", token.T_FINALLY, false, 0, 0},
		{"foreach", token.T_FOREACH, false, 0, 0},
		{"include", token.T_INCLUDE, false, 0, 0},
		{"private", token.T_PRIVATE, false, 0, 0},
		{"require", token.T_REQUIRE, false, 0, 0},
		{"__dir__", token.T_DIR, false, 0, 0},
	},
	8: {
		{"abstract", token.T_ABSTRACT, false, 0, 0},
		{"callable", token.T_CALLABLE, false, 0, 0},
		{"continue", token.T_CONTINUE, false, 0, 0},
		{"endwhile", token.T_ENDWHILE, false, 0, 0},
		{"function", token.T_FUNCTION, false, 0, 0},
		{"readonly", token.T_READONLY, true, 8, 0},
		{"__file__", token.T_FILE, false, 0, 0},
		{"__line__", token.T_LINE, false, 0, 0},
	},
	9: {
		{"endswitch", token.T_ENDSWITCH, false, 0, 0},
		{"interface", token.T_INTERFACE, false, 0, 0},
		{"insteadof", token.T_INSTEADOF, false, 0, 0},
		{"namespace", token.T_NAMESPACE, false, 0, 0},
		{"protected", token.T_PROTECTED, false, 0, 0},
		{"cfunction", token.T_FUNCTION, false, 0, 0},
		{"__class__", token.T_CLASS_C, false, 0, 0},
		{"__trait__", token.T_TRAIT_C, false, 0, 0},
	},
	10: {
		{"enddeclare", token.T_ENDDECLARE, false, 0, 0},
		{"endforeach", token.T_ENDFOREACH, false, 0, 0},
		{"instanceof", token.T_INSTANCEOF, false, 0, 0},
		{"implements", token.T_IMPLEMENTS, false, 0, 0},
		{"__method__", token.T_METHOD_C, false, 0, 0},
	},
	12: {
		{"include_once", token.T_INCLUDE_ONCE, false, 0, 0},
		{"require_once", token.T_REQUIRE_ONCE, false, 0, 0},
		{"__function__", token.T_FUNC_C, false, 0, 0},
	},
	13: {
		{"__namespace__", token.T_NS_C, false, 0, 0},
	},
	15: {
		{"__halt_compiler", token.T_HALT_COMPILER, false, 0, 0},
	},
}
