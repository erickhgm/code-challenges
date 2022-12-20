package getoddnumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOdd(t *testing.T) {
	cases := []struct {
		input    []int
		expected []int
	}{
		{input: []int{1, 2, 3}, expected: []int{1, 3}},
		{input: []int{4, 5, 6, 7, 8, 9}, expected: []int{5, 7, 9}},
	}

	for _, c := range cases {
		out := Odd(c.input)
		assert.Equal(t, c.expected, out)
	}
}
