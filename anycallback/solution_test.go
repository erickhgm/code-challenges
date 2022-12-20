package anycallback

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnyCallbackGreaterThanThree(t *testing.T) {
	cases := []struct {
		input    []int
		expected bool
	}{
		{input: []int{1, 2, 3, 4}, expected: true},
		{input: []int{1, 2, 3}, expected: false},
	}

	callback := func(n []int) bool {
		for _, v := range n {
			if v > 3 {
				return true
			}
		}
		return false
	}
	for _, c := range cases {
		out := AnyCallback(c.input, callback)
		assert.Equal(t, c.expected, out)
	}
}

func TestAnyCallbackLessThanThree(t *testing.T) {
	cases := []struct {
		input    []int
		expected bool
	}{
		{input: []int{1, 2, 3, 4}, expected: true},
		{input: []int{3, 4, 5, 6}, expected: false},
	}

	callback := func(n []int) bool {
		for _, v := range n {
			if v < 3 {
				return true
			}
		}
		return false
	}
	for _, c := range cases {
		out := AnyCallback(c.input, callback)
		assert.Equal(t, c.expected, out)
	}
}
