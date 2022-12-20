package capitalizesentence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCap(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{input: "this is a house", expected: "This Is A House"},
		{input: "flower", expected: "Flower"},
	}

	for _, c := range cases {
		out := Cap(c.input)
		assert.Equal(t, c.expected, out)
	}
}
