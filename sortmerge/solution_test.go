package template

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	cases := []struct {
		input    []int
		expected []int
	}{
		{input: []int{1, 3, 2, 4}, expected: []int{1, 2, 3, 4}},
		{input: []int{1, 3, 2, 4, 5}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{1, 3, 5, 2, 4, 6}, expected: []int{1, 2, 3, 4, 5, 6}},
		{input: []int{99, 1, 3, 5, 70, 2, 4, 6, 0}, expected: []int{0, 1, 2, 3, 4, 5, 6, 70, 99}},
	}

	for _, c := range cases {
		out := MergeSort(c.input)
		assert.Equal(t, c.expected, out)
	}
}

func TestMerge(t *testing.T) {
	cases := []struct {
		input    [][]int
		expected []int
	}{
		{input: [][]int{{1, 3}, {2, 4}}, expected: []int{1, 2, 3, 4}},
		{input: [][]int{{1, 3, 5}, {2, 4, 6}}, expected: []int{1, 2, 3, 4, 5, 6}},
	}

	for _, c := range cases {
		out := merge(c.input[0], c.input[1])
		assert.Equal(t, c.expected, out)
	}
}
