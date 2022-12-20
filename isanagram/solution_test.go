package isanagram

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAnagram(t *testing.T) {
	cases := []struct {
		input1   string
		input2   string
		expected bool
	}{
		{input1: "rail safety", input2: "fairy tales", expected: true},
		{input1: "RAIL! SAFETY!", input2: "fairy tales", expected: true},
		{input1: "Hi there", input2: "Bye there", expected: false},
	}

	for _, c := range cases {
		out := IsAnagram(c.input1, c.input2)
		assert.Equal(t, c.expected, out)
	}
}
