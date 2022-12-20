package permutationpalindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindromePermutation(t *testing.T) {
	cases := []struct {
		input    string
		expected bool
	}{
		{input: "gikig", expected: true},
		{input: "ookvk", expected: true},
		{input: "sows", expected: false},
		{input: "tami", expected: false},
	}

	for _, c := range cases {
		out := IsPalindromePermutation(c.input)
		assert.Equal(t, c.expected, out)
	}
}
