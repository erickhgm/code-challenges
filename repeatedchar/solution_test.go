package repeatedchar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyFunc(t *testing.T) {
	cases := []struct {
		input    string
		expected bool
	}{
		{input: "abc", expected: false},
		{input: "aabc", expected: true},
		{input: "aabcc", expected: true},
	}

	for _, c := range cases {
		out := Repeated(c.input)
		assert.Equal(t, c.expected, out)
	}
}
