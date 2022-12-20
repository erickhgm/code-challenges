package binarytree

type node struct {
	value int
	left  *node
	right *node
}

func NewBinaryTree(v int) *node {
	return &node{value: v}
}

func (n *node) Add(v int) {
	if n.value == v {
		return
	}

	if n.value > v {
		if n.left == nil {
			n.left = &node{value: v}
		} else {
			n.left.Add(v)
		}
	} else if n.value < v {
		if n.right == nil {
			n.right = &node{value: v}
		} else {
			n.right.Add(v)
		}
	}
}
