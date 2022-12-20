package tolerantpalindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		input    string
		expected bool
	}{
		{input: "abba", expected: true},
		{input: "abb#a", expected: true},
		{input: "abcdefg", expected: false},
		{input: "abc#cba", expected: true},
	}

	for _, c := range cases {
		out := IsPalindrome(c.input)
		assert.Equal(t, c.expected, out)
	}
}
