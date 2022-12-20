package power

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPow(t *testing.T) {
	cases := []struct {
		input1   int
		input2   int
		expected int
	}{
		{input1: 2, input2: 1, expected: 1},
		{input1: 2, input2: 2, expected: 4},
		{input1: 3, input2: 3, expected: 27},
		{input1: 2, input2: 10, expected: 1024},
	}

	for _, c := range cases {
		out := Pow(c.input1, c.input2)
		assert.Equal(t, c.expected, out)
	}
}
