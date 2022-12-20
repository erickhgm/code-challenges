package duplicatedarguments

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDup(t *testing.T) {
	cases := []struct {
		input    []string
		expected []string
	}{
		{input: []string{}, expected: []string{}},
		{input: []string{"a", "b", "c"}, expected: []string{}},
		{input: []string{"a", "b", "c", "a"}, expected: []string{"a"}},
		{input: []string{"a", "e", "a", "e", "d", "a"}, expected: []string{"a", "e"}},
	}

	for _, c := range cases {
		out := Dup(c.input)
		assert.Equal(t, c.expected, out)
	}
}
