package reversestring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{input: "abcd", expected: "dcba"},
		{input: "aabbccdd", expected: "ddccbbaa"},
	}

	for _, c := range cases {
		out := Reverse(c.input)
		assert.Equal(t, c.expected, out)
	}
}
