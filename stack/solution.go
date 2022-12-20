package stack

type stack[T any] struct {
	elmts []T
}

func NewStack[T any]() stack[T] {
	elmts := []T{}
	return stack[T]{elmts}
}

func (q *stack[T]) Add(elmt T) {
	q.elmts = append(q.elmts, elmt)
}

func (q *stack[T]) Remove() (elmt *T) {
	if len(q.elmts) == 0 {
		return nil
	}
	temp := q.elmts[len(q.elmts)-1]
	q.elmts = q.elmts[:len(q.elmts)-1]
	return &temp
}

func (q *stack[T]) Peek() (elmt *T) {
	if len(q.elmts) == 0 {
		return nil
	}
	return &q.elmts[len(q.elmts)-1]
}

func (q *stack[T]) IsEmpty() bool {
	return len(q.elmts) == 0
}

func (q *stack[T]) Size() int {
	return len(q.elmts)
}
