package insertionsort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyFunc(t *testing.T) {
	cases := []struct {
		input    []int
		expected []int
	}{
		{input: []int{}, expected: []int{}},
		{input: []int{7}, expected: []int{7}},
		{input: []int{9, 3}, expected: []int{3, 9}},
		{input: []int{5, 1, 4, 2}, expected: []int{1, 2, 4, 5}},
		{input: []int{5, 2, 1, 8, 4, 7, 6, 3}, expected: []int{1, 2, 3, 4, 5, 6, 7, 8}},
		{input: []int{17, 4, 12, 19, 80, 75, 16}, expected: []int{4, 12, 16, 17, 19, 75, 80}},
		{input: []int{11, 40, 40, 50, 43, 10, 30, 42, 20, 6, 19, 32, 20, 41, 23, 27}, expected: []int{6, 10, 11, 19, 20, 20, 23, 27, 30, 32, 40, 40, 41, 42, 43, 50}},
	}

	for _, c := range cases {
		out := Sort(c.input)
		assert.Equal(t, c.expected, out)
	}
}
