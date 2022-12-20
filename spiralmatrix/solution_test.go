package spiralmatrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	cases := []struct {
		input    int
		expected [][]int
	}{
		{input: 2, expected: [][]int{{1, 2}, {4, 3}}},
		{input: 3, expected: [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}},
		{input: 4, expected: [][]int{{1, 2, 3, 4}, {12, 13, 14, 5}, {11, 16, 15, 6}, {10, 9, 8, 7}}},
	}

	for _, c := range cases {
		out := Generate(c.input)
		assert.Equal(t, c.expected, out)
	}
}
