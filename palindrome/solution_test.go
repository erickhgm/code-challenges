package template

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
		{input: "abcdefg", expected: false},
	}

	for _, c := range cases {
		out := IsPalindrome(c.input)
		assert.Equal(t, c.expected, out)
	}
}

func TestNextPalindrome(t *testing.T) {
	cases := []struct {
		input    int
		expected int
	}{
		//{input: 1, expected: 1},
		{input: 11, expected: 22},

		{input: 120, expected: 121},     // odd
		{input: 123, expected: 131},     // odd
		{input: 193, expected: 202},     // odd
		{input: 19312, expected: 19391}, // odd
		{input: 10312, expected: 10401}, // odd
		{input: 10912, expected: 11011}, // odd

		{input: 10, expected: 11},     // even
		{input: 13, expected: 22},     // even
		{input: 19, expected: 22},     // even
		{input: 1912, expected: 1991}, // even
		{input: 1012, expected: 1111}, // even
		{input: 1912, expected: 1991}, // even
	}

	for _, c := range cases {
		out := NextPalindrome(c.input)
		assert.Equal(t, c.expected, out)
	}
}
