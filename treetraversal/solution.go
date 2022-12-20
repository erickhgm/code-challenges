package treetraversal

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
			n.left = NewBinaryTree(v)
		} else {
			n.left.Add(v)
		}
	} else if n.value < v {
		if n.right == nil {
			n.right = NewBinaryTree(v)
		} else {
			n.right.Add(v)
		}
	}
}

func (n *node) GetLevelOrder() []int {
	var result []int

	queue := []node{*n}

	for len(queue) != 0 {

		temp := queue[0]
		queue = queue[1:]
		result = append(result, temp.value)

		if temp.left != nil {
			queue = append(queue, *temp.left)
		}
		if temp.right != nil {
			queue = append(queue, *temp.right)
		}
	}
	return result
}

func (n *node) GetPreOrder() []int {
	var result []int
	var preOrder func(*node)

	preOrder = func(root *node) {
		if root != nil {
			result = append(result, root.value)
			preOrder(root.left)
			preOrder(root.right)
		}
	}
	preOrder(n)

	return result
}

func (n *node) GetInOrder() []int {
	var result []int
	var inOrder func(*node)

	inOrder = func(root *node) {
		if root != nil {
			inOrder(root.left)
			result = append(result, root.value)
			inOrder(root.right)
		}
	}
	inOrder(n)

	return result
}

func (n *node) GetPostOrder() []int {
	var result []int
	var postOrder func(*node)

	postOrder = func(root *node) {
		if root != nil {
			postOrder(root.left)
			postOrder(root.right)
			result = append(result, root.value)
		}
	}
	postOrder(n)

	return result
}
