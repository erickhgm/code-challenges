package maxoccurrentchar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyFunc(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{input: "abcccccccd", expected: "c"},
		{input: "apple 1231111", expected: "1"},
	}

	for _, c := range cases {
		out := MaxChar(c.input)
		assert.Equal(t, c.expected, out)
	}
}
