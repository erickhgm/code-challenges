package pairwithaverage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasAverage(t *testing.T) {
	cases := []struct {
		input1   []int
		input2   float32
		expected bool
	}{
		{input1: []int{}, input2: 1.0, expected: false},
		{input1: []int{3, 4, 7, 9}, input2: 6.5, expected: true},
		{input1: []int{3, 4, 7, 9}, input2: 3.0, expected: false},
		{input1: []int{3, 5, 9, 7, 11, 14}, input2: 8.0, expected: true},
	}

	for _, c := range cases {
		out := HasAverage(c.input1, c.input2)
		assert.Equal(t, c.expected, out)
	}
}
