package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	cases := []struct {
		input    int
		expected []string
	}{
		{input: 5, expected: []string{"1", "2", "Fizz", "4", "Buzz"}},
		{input: 16, expected: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz", "16"}},
	}

	for _, c := range cases {
		out := FizzBuzz(c.input)
		assert.Equal(t, c.expected, out)
	}
}
