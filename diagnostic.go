package phplint

import (
	"fmt"

	phpposition "github.com/shyim/go-phplint/internal/php/pkg/position"
)

// Position identifies a byte position in source code.
type Position struct {
	Offset int
	Line   int
	Column int
}

// Phase identifies the stage that produced a diagnostic.
type Phase uint8

const (
	PhaseLex Phase = iota + 1
	PhaseParse
	PhaseCompile
)

func (p Phase) String() string {
	switch p {
	case PhaseLex:
		return "lex"
	case PhaseParse:
		return "parse"
	case PhaseCompile:
		return "compile"
	default:
		return "unknown"
	}
}

// Diagnostic describes one source-level lint failure.
type Diagnostic struct {
	Filename string
	Message  string
	Phase    Phase
	Start    Position
	End      Position
}

func (d Diagnostic) String() string {
	return fmt.Sprintf("%s:%d:%d: %s", d.Filename, d.Start.Line, d.Start.Column, d.Message)
}

func positionFromInternal(pos *phpposition.Position) (Position, Position) {
	if pos == nil {
		unknown := Position{Line: 1, Column: 1}
		return unknown, unknown
	}

	return Position{
			Offset: pos.StartPos,
			Line:   pos.StartLine,
			Column: pos.StartCol + 1,
		}, Position{
			Offset: pos.EndPos,
			Line:   pos.EndLine,
			Column: pos.EndCol + 1,
		}
}
