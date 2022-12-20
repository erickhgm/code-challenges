package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueueString(t *testing.T) {
	queue := NewQueue[string]()

	out := queue.IsEmpty()
	assert.True(t, out)

	queue.Add("A")
	out = queue.IsEmpty()
	assert.False(t, out)

	queue.Add("B")
	queue.Add("C")

	elmt := queue.Peek()
	assert.Equal(t, "A", *elmt)
	elmt = queue.Peek()
	assert.Equal(t, "A", *elmt)

	elmt = queue.Remove()
	assert.Equal(t, "A", *elmt)

	elmt = queue.Remove()
	assert.Equal(t, "B", *elmt)

	elmt = queue.Remove()
	assert.Equal(t, "C", *elmt)

	elmt = queue.Remove()
	assert.Nil(t, elmt)
}
