package combinequeue

type queueable interface {
	string | int | any
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

func combine[T1 queueable, T2 queueable](a queue[T1], b queue[T2]) queue[any] {
	elmts := []any{}

	for a.Peek() != nil || b.Peek() != nil {
		if v := a.Remove(); v != nil {
			elmts = append(elmts, *v)
		}
		if v := b.Remove(); v != nil {
			elmts = append(elmts, *v)
		}
	}
	return queue[any]{elmts}
}
