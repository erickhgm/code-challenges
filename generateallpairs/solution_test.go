package generateallpairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPairs(t *testing.T) {
	cases := []struct {
		input    int
		expected []Pair
	}{
		{input: 0, expected: []Pair{{0, 0}}},
		{input: 1, expected: []Pair{{0, 0}, {0, 1}, {1, 0}, {1, 1}}},
		{input: 2, expected: []Pair{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 1}, {1, 2}, {2, 0}, {2, 1}, {2, 2}}},
	}

	for _, c := range cases {
		out := GetPairs(c.input)
		assert.Equal(t, c.expected, out)
	}
}
