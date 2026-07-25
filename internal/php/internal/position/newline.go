package position

// NewLines wraps a slice of ints, each int is the beginning position of a line.
type NewLines struct {
	data []int
}

func NewNewLines() NewLines {
	return NewLines{make([]int, 0, 128)}
}

func (nl *NewLines) Append(p int) {
	if len(nl.data) == 0 || nl.data[len(nl.data)-1] < p {
		nl.data = append(nl.data, p)
	}
}

// GetLine returns the line number, and beginning of the line.
func (nl *NewLines) GetLine(p int) (line int, lineStart int) {
	line = len(nl.data) + 1
	lineStart = 0
	if len(nl.data) > 0 {
		lineStart = nl.data[len(nl.data)-1]
	}

	for i := len(nl.data) - 1; i >= 0; i-- {
		if p < nl.data[i] {
			line = i + 1

			if i-1 >= 0 {
				lineStart = nl.data[i-1]
			} else {
				lineStart = 0
			}
		} else {
			break
		}
	}

	return line, lineStart
}
