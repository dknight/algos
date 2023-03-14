// Package stack represents the stack data structure and basic operations
// with it. This data structure is thread-safe.
package stack

import (
	"errors"
	"fmt"
	"sync"
)

var (
	// ErrStackIsEmpty represents the error when the stack is empty.
	ErrStackIsEmpty = errors.New("Stack is empty")
)

// Stack is a LIFO data structure.
type Stack[T any] struct {
	items []T
	mutex sync.Mutex
}

// New creates a pointer to the new stack.
func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Len returns the length of the stack.
// Please note that len() is non-thred safe read operation and also should be
// locked.
func (s *Stack[T]) Len() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return len(s.items)
}

// Push adds values to the stack.
func (s *Stack[T]) Push(vals ...T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, v := range vals {
		s.items = append(s.items, v)
	}
}

// Peek returns the last element of the stack.
func (s *Stack[T]) Peek() T {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if len(s.items) == 0 {
		panic(ErrStackIsEmpty)
	}
	return s.items[len(s.items)-1]
}

// Pop removes the last element from stack and returns it. If the stack is
// empty the nil is returned.
func (s *Stack[T]) Pop() T {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if len(s.items) == 0 {
		panic(ErrStackIsEmpty)
	}
	ret := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return ret
}

// Empty checks is the stack is empty or not.
func (s *Stack[T]) Empty() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return len(s.items) == 0
}

func (s *Stack[T]) String() string {
	return fmt.Sprintf("%+v", s.items)
}
