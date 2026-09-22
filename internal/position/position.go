package position

// Position represents node position
type Position struct {
	StartLine int
	EndLine   int
	StartCol  int
	EndCol    int
	StartPos  int
	EndPos    int
}

// NewPosition Position constructor
func NewPosition(
	StartLine int,
	EndLine int,
	StartPos int,
	EndPos int,
	StartCol int,
	EndCol int,
) *Position {
	return &Position{
		StartLine: StartLine,
		EndLine:   EndLine,
		StartPos:  StartPos,
		EndPos:    EndPos,
		StartCol:  StartCol,
		EndCol:    EndCol,
	}
}
