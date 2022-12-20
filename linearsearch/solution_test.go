package linearsearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexOf(t *testing.T) {
	cases := []struct {
		input1   []string
		input2   string
		expected int
	}{
		{input1: []string{"A", "B", "C"}, input2: "A", expected: 0},
		{input1: []string{"A", "B", "C"}, input2: "B", expected: 1},
		{input1: []string{"A", "B", "C"}, input2: "D", expected: -1},
	}

	for _, c := range cases {
		out := IndexOf(c.input1, c.input2)
		assert.Equal(t, c.expected, out)
	}
}
