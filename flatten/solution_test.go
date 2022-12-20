package flatten

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyFunc(t *testing.T) {
	cases := []struct {
		input    []any
		expected []any
	}{
		{input: []any{1}, expected: []any{1}},
		{input: []any{1, []any{2}}, expected: []any{1, 2}},
		{input: []any{1, []any{2, []any{3}}}, expected: []any{1, 2, 3}},
		{input: []any{1, []any{2, []any{3, []any{4}}}}, expected: []any{1, 2, 3, 4}},
	}

	for _, c := range cases {
		out := Flatten(c.input)
		assert.Equal(t, c.expected, out)
	}
}
