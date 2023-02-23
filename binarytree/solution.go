package binarytree

import "fmt"

// type node struct {
// 	value int
// 	left  *node
// 	right *node
// }

// func NewBinaryTree(v int) *node {
// 	return &node{value: v}
// }

// func (n *node) Add(v int) {
// 	if n.value == v {
// 		return
// 	}

// 	if n.value > v {
// 		if n.left == nil {
// 			n.left = &node{value: v}
// 		} else {
// 			n.left.Add(v)
// 		}
// 	} else if n.value < v {
// 		if n.right == nil {
// 			n.right = &node{value: v}
// 		} else {
// 			n.right.Add(v)
// 		}
// 	}
// }

type Node struct {
	value int
	left  *Node
	right *Node
}

func NewBinaryTree(value int) *Node {
	return &Node{value: value, left: nil, right: nil}
}

func (n *Node) Add(value int) {
	if value == n.value {
		return
	}
	if value < n.value {
		if n.left != nil {
			n.left.Add(value)
		} else {
			n.left = &Node{value: value}
		}
	} else {
		if n.right != nil {
			n.right.Add(value)
		} else {
			n.right = &Node{value: value}
		}
	}
}

func (n *Node) Has(value int) bool {
	if value == n.value {
		return true
	}
	if value < n.value {
		if n.left == nil {
			return false
		} else {
			return n.left.Has(value)
		}
	} else {
		if n.right == nil {
			return false
		} else {
			return n.right.Has(value)
		}
	}
}

func (n *Node) Invert() {
	if n == nil {
		return
	}
	n.right, n.left = n.left, n.right
	n.left.Invert()
	n.right.Invert()
}

// left -> root -> right
func (n *Node) PrintInOrder() {
	if n.left != nil {
		n.left.PrintInOrder()
	}
	fmt.Println(n.value)
	if n.right != nil {
		n.right.PrintInOrder()
	}
}

// root -> left -> right
func (n *Node) PrintPreOrder() {
	fmt.Println(n.value)
	if n.left != nil {
		n.left.PrintPreOrder()
	}
	if n.right != nil {
		n.right.PrintPreOrder()
	}
}

// left -> right -> root
func (n *Node) PrintPostOrder() {
	if n.left != nil {
		n.left.PrintPostOrder()
	}
	if n.right != nil {
		n.right.PrintPostOrder()
	}
	fmt.Println(n.value)
}
