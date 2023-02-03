package template

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyFunc(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{input: "A", expected: 1},
		{input: "B", expected: -1},
	}

	for _, c := range cases {
		out := MyFunc(c.input)
		assert.Equal(t, c.expected, out)
	}
}
