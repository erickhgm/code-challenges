package binarytreesearch

import "golang.org/x/exp/constraints"

type node[T constraints.Ordered] struct {
	value T
	left  *node[T]
	right *node[T]
}

func NewBinaryTree[T constraints.Ordered](v T) *node[T] {
	return &node[T]{value: v}
}

func (n *node[T]) Add(v T) {
	if n.value == v {
		return
	}

	if n.value > v {
		if n.left == nil {
			n.left = &node[T]{value: v}
		} else {
			n.left.Add(v)
		}
	} else if n.value < v {
		if n.right == nil {
			n.right = &node[T]{value: v}
		} else {
			n.right.Add(v)
		}
	}
}

func (n *node[T]) Contains(v T) bool {
	if n.value == v {
		return true
	}
	if n.value > v {
		if n.left != nil {
			return n.left.Contains(v)
		} else {
			return false
		}
	} else if n.value < v {
		if n.right != nil {
			return n.right.Contains(v)
		} else {
			return false
		}
	}
	return false
}
