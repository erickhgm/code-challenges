package selectionsort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	cases := []struct {
		input    []int
		expected []int
	}{
		{input: []int{}, expected: []int{}},
		{input: []int{7}, expected: []int{7}},
		{input: []int{9, 3}, expected: []int{3, 9}},
		{input: []int{5, 1, 4, 2}, expected: []int{1, 2, 4, 5}},
		{input: []int{5, 2, 1, 8, 4, 7, 6, 3}, expected: []int{1, 2, 3, 4, 5, 6, 7, 8}},
	}

	for _, c := range cases {
		out := Sort(c.input)
		assert.Equal(t, c.expected, out)
	}
}
