package numberswithsteps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSteps(t *testing.T) {
	cases := []struct {
		input1   int
		input2   int
		expected []int
	}{
		{input1: 0, input2: 1, expected: []int{}},
		{input1: 6, input2: 1, expected: []int{6, 5, 4, 3, 2, 1}},
		{input1: 6, input2: 2, expected: []int{6, 4, 2}},
	}

	for _, c := range cases {
		out := Steps(c.input1, c.input2)
		assert.Equal(t, c.expected, out)
	}
}
