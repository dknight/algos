// Package set represents set data structure implementation. Set has only one
// value in the collection.
package set

import (
	"fmt"
	"sort"
	"strings"
)

// Set is a set data structure.
// Similar package can be found in https://pkg.go.dev/golang.org/x/exp/maps
type Set[T comparable] map[T]bool

// New creates a new set returns pointer to it. The values in the set are in
// an indeterminate order.
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

// Keys returns slice of set's keys. Alias of Values().
func (s Set[T]) Keys() []T {
	return s.Values()
}

// Values returns all values of the set. Values are in indeterminate order.
func (s Set[T]) Values() []T {
	ret := make([]T, 0, len(s))
	for k := range s {
		ret = append(ret, k)
	}
	return ret
}

// OrderedKeys returns the keys in order, if possible to compare them.
// If keys cannot be compare, program will panic.
func (s Set[T]) OrderedKeys() []T {
	keys := s.Keys()
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] != keys[j]
	})
	return keys
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
	for _, v := range s.Values() {
		fmt.Fprintf(&b, " %v", v)
	}
	str := b.String()
	return strings.Trim(str, " ")
}
