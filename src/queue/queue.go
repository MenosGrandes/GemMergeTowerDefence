package queue

import (
	"errors"
	"fmt"
)

type Queue[T any] struct {
	Elements []T
	Size     int
}

func (q *Queue[T]) Enqueue(elem T) {
	// if q.GetLength() == q.Size {
	// 	fmt.Println("Overflow")
	// 	return
	// }
	q.Elements = append(q.Elements, elem)
}

func (q *Queue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		fmt.Println("UnderFlow")
		var zero T
		return zero, errors.New("UnderFlow queue")
	}
	element := q.Elements[0]
	if q.GetLength() == 1 {
		q.Elements = nil
		return element, nil
	}
	q.Elements = q.Elements[1:]
	return element, nil // Slice off the element once it is dequeued.
}

func (q *Queue[T]) GetLength() int {
	return len(q.Elements)
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.Elements) == 0
}

func (q *Queue[T]) Peek() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, errors.New("empty queue")
	}
	return q.Elements[0], nil
}

func NewEmptyQueue[T any]() Queue[T] {
	return Queue[T]{}
}
