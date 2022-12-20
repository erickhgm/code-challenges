package caesarcipher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCipher(t *testing.T) {
	cases := []struct {
		input1   string
		input2   int
		expected string
	}{
		{input1: "abc", input2: 1, expected: "bcd"},
		{input1: "z", input2: 1, expected: "a"},
		{input1: "abcdefghijklmnopqrstuvwxyz", input2: 1, expected: "bcdefghijklmnopqrstuvwxyza"},
		{input1: "abcdefghijklmnopqrstuvwxyz", input2: 7, expected: "hijklmnopqrstuvwxyzabcdefg"},
		{input1: "abcdefghijklmnopqrstuvwxyz", input2: 50, expected: "yzabcdefghijklmnopqrstuvwx"},
	}

	for _, c := range cases {
		out := Cipher(c.input1, c.input2)
		assert.Equal(t, c.expected, out)
	}
}
