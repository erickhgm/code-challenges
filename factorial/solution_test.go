package factorial

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorial(t *testing.T) {
	cases := []struct {
		input    int
		expected int
	}{
		{input: 1, expected: 1},
		{input: 2, expected: 2},
		{input: 3, expected: 6},
		{input: 4, expected: 24},
		{input: 5, expected: 120},
		{input: 10, expected: 3628800},
	}

	for _, c := range cases {
		out := Factorial(c.input)
		assert.Equal(t, c.expected, out)
	}
}
