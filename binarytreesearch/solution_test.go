package binarytreesearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryFourNodes(t *testing.T) {
	root := NewBinaryTree(10)
	root.Add(5)
	root.Add(15)
	root.Add(17)

	assert.Equal(t, root.value, 10)
	assert.Equal(t, root.left.value, 5)
	assert.Equal(t, root.right.value, 15)
	assert.Equal(t, root.right.right.value, 17)

	assert.True(t, root.Contains(10))
	assert.True(t, root.Contains(5))
	assert.True(t, root.Contains(15))
	assert.True(t, root.Contains(17))
	assert.False(t, root.Contains(100))
}

func TestBinarySevenNodes(t *testing.T) {
	root := NewBinaryTree(10)
	root.Add(5)
	root.Add(15)
	root.Add(0)
	root.Add(20)
	root.Add(-5)
	root.Add(3)

	assert.Equal(t, root.value, 10)
	assert.Equal(t, root.left.value, 5)
	assert.Equal(t, root.right.value, 15)
	assert.Equal(t, root.left.left.value, 0)
	assert.Equal(t, root.right.right.value, 20)
	assert.Equal(t, root.left.left.left.value, -5)
	assert.Equal(t, root.left.left.right.value, 3)

	assert.True(t, root.Contains(10))
	assert.True(t, root.Contains(5))
	assert.True(t, root.Contains(15))
	assert.True(t, root.Contains(0))
	assert.True(t, root.Contains(20))
	assert.True(t, root.Contains(-5))
	assert.True(t, root.Contains(3))
	assert.False(t, root.Contains(100))
}
