package treelevel

type node[T any] struct {
	value    T
	children []node[T]
}

func NewTree[T any](v T) *node[T] {
	return &node[T]{value: v}
}

func (n *node[T]) Add(v *node[T]) {
	n.children = append(n.children, *v)
}

func (n *node[T]) GetLevelWidth() []int {
	if n == nil {
		return []int{}
	}
	var widths []int
	children := []node[T]{*n}

	for len(children) != 0 {
		var width int
		var temp []node[T]
		for _, v := range children {
			width++
			temp = append(temp, v.children...)
		}
		widths = append(widths, width)
		children = temp
	}
	return widths
}
