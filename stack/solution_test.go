package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueueString(t *testing.T) {
	stack := NewStack[string]()

	out := stack.IsEmpty()
	assert.True(t, out)

	stack.Add("A")
	out = stack.IsEmpty()
	assert.False(t, out)

	stack.Add("B")
	stack.Add("C")

	size := stack.Size()
	assert.Equal(t, 3, size)

	elmt := stack.Peek()
	assert.Equal(t, "C", *elmt)
	elmt = stack.Peek()
	assert.Equal(t, "C", *elmt)

	elmt = stack.Remove()
	assert.Equal(t, "C", *elmt)

	elmt = stack.Remove()
	assert.Equal(t, "B", *elmt)

	elmt = stack.Remove()
	assert.Equal(t, "A", *elmt)

	elmt = stack.Remove()
	assert.Nil(t, elmt)
}
