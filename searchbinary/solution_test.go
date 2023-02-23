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

func TestBinarySearch(t *testing.T) {
	cases := []struct {
		input1   []int
		input2   int
		expected int
	}{
		{input1: []int{1, 2, 3}, input2: 1, expected: 0},
		{input1: []int{1, 2, 3}, input2: 2, expected: 1},
		{input1: []int{1, 2, 3}, input2: 3, expected: 2},
		{input1: []int{1, 2, 3}, input2: 4, expected: -1},

		{input1: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, input2: 1, expected: 0},
		{input1: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, input2: 10, expected: 9},
		{input1: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, input2: 20, expected: 19},
		{input1: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, input2: 77, expected: -1},
	}

	for _, c := range cases {
		out := binarySearch(c.input1, c.input2)
		assert.Equal(t, c.expected, out)
	}
}
