package surroundedletter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyFunc(t *testing.T) {
	cases := []struct {
		input    string
		expected bool
	}{
		{input: "a", expected: false},
		{input: "+a+", expected: true},
		{input: "a+", expected: false},
		{input: "+a", expected: false},
		{input: "+a+b+", expected: true},
		{input: "+a++b+", expected: true},
		{input: "+ab+", expected: false},
		{input: "a+b+", expected: false},
		{input: "+a+b", expected: false},
		{input: "+a+b+++c++d+e++", expected: true},
		{input: "+a+d++de++e+", expected: false},
	}

	for _, c := range cases {
		out := Surrounded(c.input)
		assert.Equal(t, c.expected, out)
	}
}
