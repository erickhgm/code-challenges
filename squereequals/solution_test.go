package squereequals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyFunc(t *testing.T) {
	cases := []struct {
		input1   []int
		input2   []int
		expected bool
	}{
		{input1: []int{2}, input2: []int{4}, expected: true},
		{input1: []int{1, 2, 3}, input2: []int{9, 1, 4}, expected: true},
		{input1: []int{1, 2, 3}, input2: []int{9, 1, 7}, expected: false}, // (does not have square of 3)
		{input1: []int{1, 2, 3}, input2: []int{9, 1}, expected: false},    // (does not have square of 2)
		{input1: []int{1, 2, 1}, input2: []int{4, 1, 4}, expected: false}, // (frequency must be the same)
	}

	for _, c := range cases {
		out := Square(c.input1, c.input2)
		assert.Equal(t, c.expected, out)
	}
}
