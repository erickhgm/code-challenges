package countupdown

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorial(t *testing.T) {
	cases := []struct {
		input    int
		expected []int
	}{
		{input: 1, expected: []int{0, 1, 0}},
		{input: 2, expected: []int{0, 1, 2, 1, 0}},
		{input: 3, expected: []int{0, 1, 2, 3, 2, 1, 0}},
		{input: 4, expected: []int{0, 1, 2, 3, 4, 3, 2, 1, 0}},
		{input: 5, expected: []int{0, 1, 2, 3, 4, 5, 4, 3, 2, 1, 0}},
	}

	for _, c := range cases {
		out := UpAndDown(c.input)
		assert.Equal(t, c.expected, out)
	}
}
