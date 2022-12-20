package findrectangle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoordinates(t *testing.T) {
	cases := []struct {
		input    [][]int
		expected []int
	}{
		{
			input: [][]int{
				{1, 1, 0, 0, 0, 1},
				{1, 1, 0, 0, 0, 1},
				{1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1},
			},
			expected: []int{0, 2, 1, 4},
		},
		{
			input: [][]int{
				{1, 1, 0, 0, 0, 1},
				{1, 1, 0, 0, 0, 1},
				{1, 1, 0, 0, 0, 1},
				{1, 1, 1, 1, 1, 1},
			},
			expected: []int{0, 2, 2, 4},
		},
	}

	for _, c := range cases {
		out := Coordinates(c.input)
		assert.Equal(t, c.expected, out)
	}
}
