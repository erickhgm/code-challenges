package countdown

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDown(t *testing.T) {
	cases := []struct {
		input    int
		expected []int
	}{
		{input: 1, expected: []int{1, 0}},
		{input: 4, expected: []int{4, 3, 2, 1, 0}},
	}

	for _, c := range cases {
		out := Down(c.input)
		assert.Equal(t, c.expected, out)
	}
}
