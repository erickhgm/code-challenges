package treetraversal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryTreeRight(t *testing.T) {
	root := NewBinaryTree(1)
	root.Add(2)
	root.Add(3)
	root.Add(4)
	root.Add(5)
	root.Add(6)
	root.Add(7)

	assert.Equal(t, root.value, 1)
	assert.Equal(t, root.right.value, 2)
	assert.Equal(t, root.right.right.value, 3)
	assert.Equal(t, root.right.right.right.value, 4)
	assert.Equal(t, root.right.right.right.right.value, 5)
	assert.Equal(t, root.right.right.right.right.right.value, 6)
	assert.Equal(t, root.right.right.right.right.right.right.value, 7)
}

func TestBinaryTreeLeft(t *testing.T) {
	root := NewBinaryTree(7)
	root.Add(6)
	root.Add(5)
	root.Add(4)
	root.Add(3)
	root.Add(2)
	root.Add(1)

	assert.Equal(t, root.value, 7)
	assert.Equal(t, root.left.value, 6)
	assert.Equal(t, root.left.left.value, 5)
	assert.Equal(t, root.left.left.left.value, 4)
	assert.Equal(t, root.left.left.left.left.value, 3)
	assert.Equal(t, root.left.left.left.left.left.value, 2)
	assert.Equal(t, root.left.left.left.left.left.left.value, 1)
}

func TestBinaryTree(t *testing.T) {
	root := NewBinaryTree(50)
	root.Add(30)
	root.Add(20)
	root.Add(40)
	root.Add(70)
	root.Add(60)
	root.Add(80)

	assert.Equal(t, root.value, 50)
	assert.Equal(t, root.left.value, 30)
	assert.Equal(t, root.left.left.value, 20)
	assert.Equal(t, root.left.right.value, 40)
	assert.Equal(t, root.right.value, 70)
	assert.Equal(t, root.right.left.value, 60)
	assert.Equal(t, root.right.right.value, 80)
}

func TestBinaryTreeGetLevelOrder(t *testing.T) {
	// Breath First traversal (BF traversal) - By row
	//        50
	//      /    \
	//    30      70
	//   /  \    /  \
	// 20   40  60   80
	//
	// Output: 50 -> 30 -> 70 -> 20 -> 40 -> 60 -> 80

	root := NewBinaryTree(50)
	root.Add(30)
	root.Add(20)
	root.Add(40)
	root.Add(70)
	root.Add(60)
	root.Add(80)

	assert.Equal(t, root.value, 50)
	assert.Equal(t, root.left.value, 30)
	assert.Equal(t, root.left.left.value, 20)
	assert.Equal(t, root.left.right.value, 40)
	assert.Equal(t, root.right.value, 70)
	assert.Equal(t, root.right.left.value, 60)
	assert.Equal(t, root.right.right.value, 80)

	out := root.GetLevelOrder()
	assert.Equal(t, []int{50, 30, 70, 20, 40, 60, 80}, out)
}

func TestBinaryTreePreOrder(t *testing.T) {
	// Pre Order Depth First traversal
	//        50
	//      /    \
	//    30      70
	//   /  \    /  \
	// 20   40  60   80
	//
	// Output: 50 -> 30 -> 20 -> 40 -> 70 -> 60 -> 80

	root := NewBinaryTree(50)
	root.Add(30)
	root.Add(20)
	root.Add(40)
	root.Add(70)
	root.Add(60)
	root.Add(80)

	assert.Equal(t, root.value, 50)
	assert.Equal(t, root.left.value, 30)
	assert.Equal(t, root.left.left.value, 20)
	assert.Equal(t, root.left.right.value, 40)
	assert.Equal(t, root.right.value, 70)
	assert.Equal(t, root.right.left.value, 60)
	assert.Equal(t, root.right.right.value, 80)

	out := root.GetPreOrder()
	assert.Equal(t, []int{50, 30, 20, 40, 70, 60, 80}, out)
}

func TestBinaryTreeInOrder(t *testing.T) {
	//  In Order Depth First traversal - By column
	//        50
	//      /    \
	//    30      70
	//   /  \    /  \
	// 20   40  60   80
	//
	// Output: 20 -> 30 -> 40 -> 50 -> 60 -> 70 -> 80

	root := NewBinaryTree(50)
	root.Add(30)
	root.Add(20)
	root.Add(40)
	root.Add(70)
	root.Add(60)
	root.Add(80)

	assert.Equal(t, root.value, 50)
	assert.Equal(t, root.left.value, 30)
	assert.Equal(t, root.left.left.value, 20)
	assert.Equal(t, root.left.right.value, 40)
	assert.Equal(t, root.right.value, 70)
	assert.Equal(t, root.right.left.value, 60)
	assert.Equal(t, root.right.right.value, 80)

	out := root.GetInOrder()
	assert.Equal(t, []int{20, 30, 40, 50, 60, 70, 80}, out)
}

func TestBinaryTreePostOrder(t *testing.T) {
	//  Post Order Depth First traversal
	//        50
	//      /    \
	//    30      70
	//   /  \    /  \
	// 20   40  60   80
	//
	// Output: 20 -> 40 -> 30 -> 50 -> 60 -> 80 -> 70

	root := NewBinaryTree(50)
	root.Add(30)
	root.Add(20)
	root.Add(40)
	root.Add(70)
	root.Add(60)
	root.Add(80)

	assert.Equal(t, root.value, 50)
	assert.Equal(t, root.left.value, 30)
	assert.Equal(t, root.left.left.value, 20)
	assert.Equal(t, root.left.right.value, 40)
	assert.Equal(t, root.right.value, 70)
	assert.Equal(t, root.right.left.value, 60)
	assert.Equal(t, root.right.right.value, 80)

	out := root.GetPostOrder()
	assert.Equal(t, []int{20, 40, 30, 60, 80, 70, 50}, out)
}
