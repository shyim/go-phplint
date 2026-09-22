package php

import (
	"bytes"

	"github.com/shyim/go-phplint/internal/conf"
	"github.com/shyim/go-phplint/internal/errors"
	"github.com/shyim/go-phplint/internal/posbuilder"
	"github.com/shyim/go-phplint/internal/position"
	"github.com/shyim/go-phplint/internal/token"
	"github.com/shyim/go-phplint/internal/version"
)

type Lexer struct {
	data           []byte
	phpVersion     *version.Version
	errHandlerFunc func(*errors.Error)

	// p: current position being lexed/checked.
	// pe: length in bytes of input.
	p, pe, cs int
	// ts: start position of the current token.
	// te: end position of the current token.
	ts, te, act int
	stack       []int
	top         int

	heredocLabel []byte
	tokenPool    *token.Pool
	positionPool *position.Pool
	newLines     posbuilder.NewLines

	// retainFreeFloating attaches whitespace and comments to the next token.
	// Lint parsing leaves this false.
	retainFreeFloating bool
}

func NewLexer(data []byte, config conf.Config) *Lexer {
	lex := &Lexer{
		data:           data,
		phpVersion:     config.Version,
		errHandlerFunc: config.ErrorHandlerFunc,

		pe:    len(data),
		stack: make([]int, 0),

		tokenPool:          token.NewPool(token.DefaultBlockSize),
		positionPool:       position.NewPool(position.DefaultBlockSize),
		newLines:           posbuilder.NewNewLines(),
		retainFreeFloating: config.Fidelity,
	}

	initLexer(lex)

	return lex
}

func (lex *Lexer) setTokenPosition(token *token.Token) {
	pos := lex.positionPool.Get()

	sl, slb := lex.newLines.GetLine(lex.ts)
	el, elb := lex.newLines.GetLine(lex.te - 1)

	pos.StartLine = sl
	pos.EndLine = el
	pos.StartPos = lex.ts
	pos.EndPos = lex.te
	pos.StartCol = lex.ts - slb
	pos.EndCol = lex.te - elb

	token.Position = pos
}

func (lex *Lexer) addFreeFloatingToken(t *token.Token, id token.ID, ps, pe int) {
	if !lex.retainFreeFloating {
		return
	}

	skippedTkn := lex.tokenPool.Get()
	skippedTkn.ID = id
	skippedTkn.Value = lex.data[ps:pe]

	lex.setTokenPosition(skippedTkn)

	if t.FreeFloating == nil {
		t.FreeFloating = make([]*token.Token, 0, 2)
	}

	t.FreeFloating = append(t.FreeFloating, skippedTkn)
}

func (lex *Lexer) isNotStringVar() bool {
	p := lex.p
	if lex.data[p-1] == '\\' && lex.data[p-2] != '\\' {
		return true
	}

	if len(lex.data) < p+1 {
		return true
	}

	if lex.data[p] == '$' && (lex.data[p+1] == '{' || isValidVarNameStart(lex.data[p+1])) {
		return false
	}

	if lex.data[p] == '{' && lex.data[p+1] == '$' {
		return false
	}

	return true
}

func (lex *Lexer) isNotStringEnd(s byte) bool {
	p := lex.p
	if lex.data[p-1] == '\\' && lex.data[p-2] != '\\' {
		return true
	}

	return !(lex.data[p] == s)
}

// versionAtLeast reports whether the target language profile is at least
// the given major.minor version. A nil version is treated as the newest
// profile so the unified scanner accepts the full superset.
func (lex *Lexer) versionAtLeast(major, minor uint64) bool {
	return lex.phpVersion.AtLeast(major, minor)
}

func (lex *Lexer) isHeredocEnd(p int) bool {
	if !lex.versionAtLeast(7, 3) {
		return lex.isHeredocEndBefore73(p)
	}

	return lex.isHeredocEndSince73(p)
}

func (lex *Lexer) isHeredocEndBefore73(p int) bool {
	if lex.data[p-1] != '\r' && lex.data[p-1] != '\n' {
		return false
	}

	l := len(lex.heredocLabel)
	if len(lex.data) < p+l {
		return false
	}

	if len(lex.data) > p+l && lex.data[p+l] != ';' && lex.data[p+l] != '\r' && lex.data[p+l] != '\n' {
		return false
	}

	if len(lex.data) > p+l+1 && lex.data[p+l] == ';' && lex.data[p+l+1] != '\r' && lex.data[p+l+1] != '\n' {
		return false
	}

	return bytes.Equal(lex.heredocLabel, lex.data[p:p+l])
}

func (lex *Lexer) isHeredocEndSince73(p int) bool {
	if lex.data[p-1] != '\r' && lex.data[p-1] != '\n' {
		return false
	}

	if p == len(lex.data) {
		return false
	}

	for lex.data[p] == ' ' || lex.data[p] == '\t' {
		p++
	}

	l := len(lex.heredocLabel)
	if len(lex.data) < p+l {
		return false
	}

	if len(lex.data) > p+l && isValidVarName(lex.data[p+l]) {
		return false
	}

	if bytes.Equal(lex.heredocLabel, lex.data[p:p+l]) {
		lex.p = p
		return true
	}

	return false
}

func (lex *Lexer) isNotHeredocEnd(p int) bool {
	return !lex.isHeredocEnd(p)
}

func (lex *Lexer) growCallStack() {
	if lex.top == len(lex.stack) {
		lex.stack = append(lex.stack, 0)
	}
}

func (lex *Lexer) isNotPhpCloseToken() bool {
	if lex.p+1 == len(lex.data) {
		return true
	}

	return lex.data[lex.p] != '?' || lex.data[lex.p+1] != '>'
}

func (lex *Lexer) isNotNewLine() bool {
	if lex.data[lex.p] == '\n' && lex.data[lex.p-1] == '\r' {
		return true
	}

	return lex.data[lex.p-1] != '\n' && lex.data[lex.p-1] != '\r'
}

func (lex *Lexer) call(state int, fnext int) {
	lex.growCallStack()

	lex.stack[lex.top] = state
	lex.top++

	lex.p++
	lex.cs = fnext
}

func (lex *Lexer) ret(n int) {
	lex.top = lex.top - n
	if lex.top < 0 {
		lex.top = 0
	}
	lex.cs = lex.stack[lex.top]
	lex.p++
}

func (lex *Lexer) ungetStr(s string) {
	n := len(s)
	if n == 0 || lex.te-lex.ts < n {
		return
	}
	tail := lex.data[lex.te-n : lex.te]
	for i := 0; i < n; i++ {
		if tail[i] != s[i] {
			return
		}
	}
	lex.ungetCnt(n)
}

func (lex *Lexer) ungetCnt(n int) {
	lex.p = lex.p - n
	lex.te = lex.te - n
}

func (lex *Lexer) ungetWhile(s byte) {
	for lex.te > 0 && lex.te < len(lex.data) && lex.p > 0 && lex.data[lex.te] != s {
		lex.te--
		lex.p--
	}

	lex.te++
	lex.p++
}

func (lex *Lexer) error(msg string) {
	if lex.errHandlerFunc == nil {
		return
	}

	sl, slb := lex.newLines.GetLine(lex.ts)
	el, elb := lex.newLines.GetLine(lex.te - 1)
	pos := position.NewPosition(
		sl,
		el,
		lex.ts,
		lex.te,
		lex.ts-slb,
		lex.te-elb,
	)

	lex.errHandlerFunc(errors.NewError(msg, pos))
}

func isValidVarNameStart(r byte) bool {
	return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || r == '_' || r >= 0x80
}

func isValidVarName(r byte) bool {
	return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '_' ||
		r >= 0x80
}
