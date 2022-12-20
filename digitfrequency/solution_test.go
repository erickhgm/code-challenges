package digitfrequency

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqualsFrequency(t *testing.T) {
	cases := []struct {
		input1   string
		input2   string
		expected bool
	}{
		{input1: "123", input2: "123", expected: true},
		{input1: "4557", input2: "745", expected: false},
	}

	for _, c := range cases {
		out := EqualsFrequency(c.input1, c.input2)
		assert.Equal(t, c.expected, out)
	}
}
