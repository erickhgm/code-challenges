package binarytree

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
	//////////////////////
	//        50        //
	//    30      70    //
	//  20  40  60  80  //
	//////////////////////
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

	assert.Equal(t, true, root.Has(50))
	assert.Equal(t, true, root.Has(30))
	assert.Equal(t, true, root.Has(20))
	assert.Equal(t, true, root.Has(40))
	assert.Equal(t, true, root.Has(70))
	assert.Equal(t, true, root.Has(60))
	assert.Equal(t, true, root.Has(80))

	assert.Equal(t, false, root.Has(10))
	assert.Equal(t, false, root.Has(100))

	root.PrintInOrder()
	root.PrintPreOrder()
	root.PrintPostOrder()
}

func TestBinaryTreeInvert(t *testing.T) {
	//////////////////////
	//        50        //
	//    30      70    //
	//  20  40  60  80  //
	//////////////////////
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

	root.PrintInOrder()
	root.Invert()
	root.PrintInOrder()
}
