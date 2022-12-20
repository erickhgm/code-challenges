package sumzero

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumZero(t *testing.T) {
	cases := []struct {
		input    []int
		expected []int
	}{
		{input: []int{}, expected: []int{}},
		{input: []int{1, 2}, expected: []int{}},
		{input: []int{1, 2, 4, 7}, expected: []int{}},
		{input: []int{-4, -3, -2, 1, 2, 3, 5}, expected: []int{-3, 3}},
		{input: []int{-4, -3, -2, 1, 2, 5}, expected: []int{-2, 2}},
	}

	for _, c := range cases {
		out := SumZero(c.input)
		assert.Equal(t, c.expected, out)
	}
}
