package factorial

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFactorial(t *testing.T) {
	cases := []struct {
		input    int64
		expected int64
	}{
		{input: 1, expected: 1},
		{input: 2, expected: 2},
		{input: 3, expected: 6},
		{input: 4, expected: 24},
		{input: 5, expected: 120},
		{input: 6, expected: 720},
		{input: 10, expected: 3628800},
	}

	for _, c := range cases {
		out := Factorial(c.input)
		assert.Equal(t, c.expected, out)
	}
}

func TestFactorialRecursive(t *testing.T) {
	cases := []struct {
		input    int64
		expected int64
	}{
		{input: 1, expected: 1},
		{input: 2, expected: 2},
		{input: 3, expected: 6},
		{input: 4, expected: 24},
		{input: 5, expected: 120},
		{input: 6, expected: 720},
		{input: 10, expected: 3628800},
	}

	for _, c := range cases {
		out := FactorialRecursive(c.input)
		assert.Equal(t, c.expected, out)
	}
}

func TestFactorialPerformance(t *testing.T) {
	input := int64(1000)
	expected := 3628800

	before := time.Now()
	out := Factorial(input)
	t.Log("time:", time.Since(before))
	t.Log("out:", out)

	assert.Equal(t, expected, 1)

	t.Log("***********")

	before = time.Now()
	out = FactorialRecursive(input)
	t.Log("time:", time.Since(before))
	t.Log("out:", out)

	assert.Equal(t, expected, 1)
}

func TestPrintPermutations(t *testing.T) {
	cases := []struct {
		input []int
	}{
		{input: []int{1, 2}},
		{input: []int{1, 2, 3}},
	}

	for _, c := range cases {
		PrintPermutations(c.input, len(c.input))
	}
}
