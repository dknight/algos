// Package deque represents deque abstract data structure and basic
// operations on it. Deque is similar to queue, but elements can be inserted
// and removed from the start and from the end.
//
// Actually deque is the generalized structure of all types of queues.
package deque

// Deque represents the deque structure.
type Deque[T any] []T

// New creates new deque instance and returns pointer to it.
func New[T any]() *Deque[T] {
	return &Deque[T]{}
}

// PushBack puts an element to the end of the deque and returns pushed element.
func (d *Deque[T]) PushBack(v T) T {
	*d = append(*d, v)
	return v
}

// PopBack pops out an element from the end of the deque and returns popped
// element. Before pop it is a good practice to check is deque is empty or
// not.
//
//	deque := deque.New[any]()
//	if !deque.Empty() {
//		_ = deque.PopBack()
//	}
//
func (d *Deque[T]) PopBack() T {
	if d.Empty() {
		panic("Empty deque")
	}
	ret := (*d)[len(*d)-1]
	*d = (*d)[:len(*d)-1]
	return ret
}

// PushFront puts an element to the beginning of the deque and returns pushed
// element.
func (d *Deque[T]) PushFront(v T) T {
	*d = append([]T{v}, *d...)
	return v
}

// PopFront pops out an element from the front of the deque and returns popped
// element. Before pop it is a good practice to check is deque is empty or
// not.
//
//	deque := deque.New[any]()
//	if !deque.Empty() {
//		_ = deque.PopFront()
//	}
//
func (d *Deque[T]) PopFront() T {
	if d.Empty() {
		panic("Empty deque")
	}
	ret := (*d)[0]
	*d = (*d)[1:]
	return ret
}

// Front returns the front (first) element of the deque.
func (d *Deque[T]) Front() T {
	return (*d)[0]
}

// Back returns the back (last) element of the deque.
func (d *Deque[T]) Back() T {
	return (*d)[len(*d)-1]
}

// Len returns the size (length) of the deque
func (d *Deque[T]) Len() int {
	return len(*d)
}

// Empty checks that deque is empty.
func (d *Deque[T]) Empty() bool {
	return len(*d) == 0
}
