package common

type Queue[T Numeric] struct {
	values []T
}

func NewQueue[T Numeric]() Queue[T] {
	q := Queue[T]{}
	q.values = make([]T, 0)

	return q
}

func (q *Queue[T]) Enqueue(key T) {
	q.values = append(q.values, key)
}

func (q *Queue[T]) Dequeue() T {
	if len(q.values) > 1 {
		q.values = q.values[1:]
	} else {
		q.values = []T{}
	}

	return q.values[0]
}
