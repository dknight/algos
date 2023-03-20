// Package set represents set data structure implementation. Set has only one
// value in the collection.
package set

import (
	"fmt"
	"strings"
)

// Set is a set data structure.
type Set[T comparable] map[T]bool

// New creates a new set returns pointer to it.
// NOTE Set is not ordered data structure.
func New[T comparable]() Set[T] {
	return make(Set[T], 0)
}

// Adds value to the set and returns pointer to newly added value.
func (s Set[T]) Add(v T) *T {
	s[v] = true
	return &v
}

// Remove removes the value from the set and returns a pointer to deleted
// element.
func (s Set[T]) Remove(v T) *T {
	delete(s, v)
	return &v
}

// Contains checks that given element is present in the set.
func (s Set[T]) Contains(v T) bool {
	_, ok := s[v]
	return ok
}

// Clear whole set. Remove all values from the set.
func (s Set[T]) Clear() {
	for k := range s {
		delete(s, k)
	}
}

// All values of the set.
func (s Set[T]) All() []T {
	ret := make([]T, 0, len(s))
	for k := range s {
		ret = append(ret, k)
	}
	return ret
}

// Len returns the length of the set.
func (s Set[T]) Len() int {
	return len(s)
}

// Empty checks that set is empty.
func (s Set[T]) Empty() bool {
	return len(s) == 0
}

// String represents set as a string.
func (s Set[T]) String() string {
	var b strings.Builder
	for _, v := range s.All() {
		fmt.Fprintf(&b, " %v", v)
	}
	str := b.String()
	return strings.Trim(str, " ")
}
