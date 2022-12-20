package queue

type queueable interface {
	string | int64 | float64
}

type queue[T queueable] struct {
	elmts []T
}

func NewQueue[T queueable]() queue[T] {
	elmts := []T{}
	return queue[T]{elmts}
}

func (q *queue[T]) Add(elmt T) {
	q.elmts = append(q.elmts, elmt)
}

func (q *queue[T]) Remove() (elmt *T) {
	if len(q.elmts) == 0 {
		return nil
	}
	temp := q.elmts[0]
	q.elmts = q.elmts[1:]
	return &temp
}

func (q *queue[T]) Peek() (elmt *T) {
	if len(q.elmts) == 0 {
		return nil
	}
	return &q.elmts[0]
}

func (q *queue[T]) IsEmpty() bool {
	return len(q.elmts) == 0
}
