package template

import (
	"testing"
)

func TestMyFunc(t *testing.T) {
	cases := []struct {
		input []int
	}{
		//{input: []int{2, 3}},
		{input: []int{2, 3}},
	}

	for _, c := range cases {
		MyFunc(c.input)
	}
}
