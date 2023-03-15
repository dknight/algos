// Package stack represents the stack data structure and basic operations
// with it. This data structure is non-thread safe.
//
// NOTE That implementing stack using slices without structure performs really
// faster. If you need really fast performance you shouldn't not to use this
// data structure. Just use raw slices with certain types.
// Comparison can be view in `../../algorithms/stocks` problem.
package stack

import (
	"errors"
	"fmt"
)

var (
	// ErrStackIsEmpty represents the error when the stack is empty.
	ErrStackIsEmpty = errors.New("Stack is empty")
)

// Stack is a LIFO data structure.
type Stack[T any] []T

// New creates a pointer to the new stack.
func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Len returns the length of the stack.
func (s *Stack[T]) Len() int {
	return len(*s)
}

// Push adds values to the stack.
func (s *Stack[T]) Push(vals ...T) {
	for _, v := range vals {
		*s = append(*s, v)
	}
}

// Peek returns the last element of the stack.
func (s *Stack[T]) Peek() T {
	if len(*s) == 0 {
		panic(ErrStackIsEmpty)
	}
	return (*s)[len(*s)-1]
}

// Pop removes the last element from stack and returns it. If the stack is
// empty the nil is returned.
func (s *Stack[T]) Pop() T {
	if len(*s) == 0 {
		panic(ErrStackIsEmpty)
	}
	ret := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return ret
}

// Empty checks is the stack is empty or not.
func (s *Stack[T]) Empty() bool {
	return len(*s) == 0
}

// Reset remove all elements from the stack.
func (s *Stack[T]) Reset() {
	*s = *New[T]()
}

// Dump returns all elements from the stack.
func (s *Stack[T]) Dump() []T {
	return []T(*s)
}

// String represents stack as a string.
func (s *Stack[T]) String() string {
	return fmt.Sprintf("%+v", *s)
}
