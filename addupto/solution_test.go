package addupto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpTo(t *testing.T) {
	cases := []struct {
		input    int
		expected int
	}{
		{input: 1, expected: 1},
		{input: 2, expected: 3},
		{input: 3, expected: 6},
	}

	for _, c := range cases {
		out := UpTo(c.input)
		assert.Equal(t, c.expected, out)
	}
}
