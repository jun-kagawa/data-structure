package datastructure

import (
	"errors"
)

type ArrayQueue[T any] struct {
	array []T
	j     int
	n     int
}

var ErrQueueEmpty = errors.New("queue is empty")

func NewArrayQueue[T any]() *ArrayQueue[T] {
	return &ArrayQueue[T]{
		array: make([]T, 6),
	}
}

func (q *ArrayQueue[T]) Size() int {
	return q.n
}

func (q *ArrayQueue[T]) Add(x T) bool {
	if q.n+1 >= len(q.array) {
		q.resize()
	}
	q.array[(q.j+q.n)%len(q.array)] = x
	q.n++
	return true
}

func (q *ArrayQueue[T]) Remove() (T, error) {
	if q.n == 0 {
		var zero T
		return zero, ErrQueueEmpty
	}
	x := q.array[q.j]
	clear(q.array[q.j : q.j+1])
	q.j = (q.j + 1) % len(q.array)
	q.n--
	if len(q.array) >= 3*q.n {
		q.resize()
	}
	return x, nil
}

func (q *ArrayQueue[T]) resize() {
	array := make([]T, max(1, 2*q.n))
	for k := 0; k < q.n; k++ {
		array[k] = q.array[(q.j+k)%len(q.array)]
	}
	q.array = array
	q.j = 0
}
