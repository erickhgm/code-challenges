package formattrainroute

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {
	cases := []struct {
		input    []string
		expected string
	}{
		{input: []string{"Luton"}, expected: "Train is calling at Luton"},
		{input: []string{"Luton", "Harpenden"}, expected: "Train is calling at Luton and Harpenden"},
		{input: []string{"Luton", "Harpenden", "London"}, expected: "Train is calling at Luton, Harpenden and London"},
		{input: []string{"Luton", "Harpenden", "London", "NY"}, expected: "Train is calling at Luton, Harpenden, London and NY"},
	}

	for _, c := range cases {
		out := Format(c.input)
		assert.Equal(t, c.expected, out)
	}
}
