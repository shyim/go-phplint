package position_test

import (
	"strconv"
	"testing"

	"github.com/shyim/go-phplint/internal/php/internal/position"
	"gotest.tools/assert"
)

func TestNewLine(t *testing.T) {
	cases := []struct {
		lines        []int
		pos          int
		expectedLine int
		expectedBol  int
	}{
		{
			lines:        []int{},
			pos:          0,
			expectedLine: 1,
			expectedBol:  0,
		},
		{
			lines:        []int{3, 10, 28},
			pos:          4,
			expectedLine: 2,
			expectedBol:  3,
		},
		{
			lines:        []int{7},
			pos:          7,
			expectedLine: 2,
			expectedBol:  7,
		},
		{
			lines:        []int{6, 8, 16, 18, 19, 21, 22},
			pos:          8,
			expectedLine: 3,
			expectedBol:  8,
		},
	}

	for i, tt := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			nl := position.NewNewLines()
			for _, l := range tt.lines {
				nl.Append(l)
			}

			line, bol := nl.GetLine(tt.pos)
			assert.Equal(t, line, tt.expectedLine)
			assert.Equal(t, bol, tt.expectedBol)
		})
	}
}
