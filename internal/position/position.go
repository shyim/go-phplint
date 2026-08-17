package position

// Tests that were written before the addition of StartCol and EndCol can set this to false,
// Which will not check if StartCol or EndCol are equal in the Equal function.
var CheckColEquality = true

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

func (p Position) Equal(other Position) bool {
	if p.StartLine != other.StartLine {
		return false
	}

	if p.StartPos != other.StartPos {
		return false
	}

	if p.EndLine != other.EndLine {
		return false
	}

	if p.EndPos != other.EndPos {
		return false
	}

	if CheckColEquality {
		if p.StartCol != other.StartCol {
			return false
		}

		if p.EndCol != other.EndCol {
			return false
		}
	}

	return true
}
