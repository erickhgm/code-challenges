package fibonacci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalc(t *testing.T) {
	cases := []struct {
		input    int
		expected []int
	}{
		{input: 0, expected: []int{0}},
		{input: 1, expected: []int{0, 1}},
		{input: 2, expected: []int{0, 1, 1}},
		{input: 3, expected: []int{0, 1, 1, 2}},
		{input: 4, expected: []int{0, 1, 1, 2, 3}},
		{input: 5, expected: []int{0, 1, 1, 2, 3, 5}},
		{input: 6, expected: []int{0, 1, 1, 2, 3, 5, 8}},
		{input: 7, expected: []int{0, 1, 1, 2, 3, 5, 8, 13}},
		{input: 8, expected: []int{0, 1, 1, 2, 3, 5, 8, 13, 21}},
		{input: 9, expected: []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}},
		{input: 10, expected: []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}},
	}

	for _, c := range cases {
		out := Calc(c.input)
		assert.Equal(t, c.expected, out)
	}
}

func TestFib(t *testing.T) {
	cases := []struct {
		input    int
		expected int
	}{
		{input: 0, expected: 0},   // 0
		{input: 1, expected: 1},   // 0, 1
		{input: 2, expected: 1},   // 0, 1, 1
		{input: 3, expected: 2},   // 0, 1, 1, 2
		{input: 4, expected: 3},   // 0, 1, 1, 2, 3
		{input: 5, expected: 5},   // 0, 1, 1, 2, 3, 5
		{input: 6, expected: 8},   // 0, 1, 1, 2, 3, 5, 8
		{input: 7, expected: 13},  // 0, 1, 1, 2, 3, 5, 8, 13
		{input: 8, expected: 21},  // 0, 1, 1, 2, 3, 5, 8, 13, 21
		{input: 9, expected: 34},  // 0, 1, 1, 2, 3, 5, 8, 13, 21, 34
		{input: 10, expected: 55}, // 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55
	}

	for _, c := range cases {
		out := Fib(c.input)
		assert.Equal(t, c.expected, out)
	}
}
