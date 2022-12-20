package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	cases := []struct {
		input1   []string
		input2   string
		expected int
	}{
		{input1: []string{"A", "B", "C"}, input2: "A", expected: 0},
		{input1: []string{"A", "B", "C"}, input2: "B", expected: 1},
		{input1: []string{"A", "B", "C"}, input2: "H", expected: -1},

		{input1: []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "O", "P"}, input2: "A", expected: 0},
		{input1: []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "O", "P"}, input2: "P", expected: 14},

		{input1: []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "O", "P"}, input2: "B", expected: 1},
		{input1: []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "O", "P"}, input2: "O", expected: 13},
		{input1: []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "O", "P"}, input2: "Z", expected: -1},
	}

	for _, c := range cases {
		out := Search(c.input1, c.input2)
		assert.Equal(t, c.expected, out)
	}
}
