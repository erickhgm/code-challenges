package product

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {
	cases := []struct {
		input    []int
		expected int
	}{
		{input: []int{1, 2}, expected: 2},
		{input: []int{1, 2, 4}, expected: 8},
		{input: []int{2, 4, 10}, expected: 80},
	}

	for _, c := range cases {
		out := Product(c.input)
		assert.Equal(t, c.expected, out)
	}
}
