// Package queue presents the queue abstract data structure and basic
// operations with it.
package queue

// Queue represents the queue structure.
type Queue[T any] []T

// New creates new queue instance and returns pointer to it.
func New[T any]() *Queue[T] {
	return &Queue[T]{}
}

// Enqueue puts an element to the end of the queue and returns enqueued
// element.
func (q *Queue[T]) Enqueue(v T) T {
	*q = append(*q, v)
	return v
}

// Dequeue puts an element to the end of the queue and returns dequeued
// element. Before dequeing it is practice to check is queue is empty or not.
//
//	queue := queue.New[any]()
//	if !queue.Empty() {
//		_ = queue.Dequeue()
//	}
func (q *Queue[T]) Dequeue() T {
	if q.Empty() {
		panic("Empty queue")
	}
	ret := (*q)[0]
	*q = (*q)[1:]
	return ret
}

// Head returns the head (first element) of the queue.
func (q *Queue[T]) Head() T {
	return (*q)[0]
}

// Tail returns the tail (last element) of the queue.
func (q *Queue[T]) Tail() T {
	return (*q)[len(*q)-1]
}

// Len returns the size (length) of the queue.
func (q *Queue[T]) Len() int {
	return len(*q)
}

// Empty checks that queue is empty.
func (q *Queue[T]) Empty() bool {
	return len(*q) == 0
}
