package decapitalizeconst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecapitalize(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{input: "FOO", expected: "foo"},
		{input: "FOO_BAR", expected: "fooBar"},
		{input: "__FOO_BAR_BAZ", expected: "fooBarBaz"},
	}

	for _, c := range cases {
		out := Decapitalize(c.input)
		assert.Equal(t, c.expected, out)
	}
}
