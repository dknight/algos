package stocks

import (
	"github.com/dknight/algos/structures/stack"
)

// SpansSlices stocks problem resolver based on slices (arrays).
func SpansSlices(quotes []int) []int {
	n := len(quotes)
	spans := make([]int, n)

	for i := 0; i < n; i++ {
		k := 1
		end := false
		for i-k >= 0 && !end {
			if quotes[i-k] <= quotes[i] {
				k++
			} else {
				end = true
			}
		}
		spans[i] = k
	}
	return spans
}

// SpansStack stocks problem resolver based on stack (arrays).
//
// NOTE: Working with raw primitives is much faster rather what using
// structures.
func SpansStack(quotes []int) []int {
	n := len(quotes)
	spans := make([]int, n)
	s := stack.New[int]()

	spans[0] = 1
	s.Push(0)

	for i := 1; i < n; i++ {
		for !s.Empty() && quotes[s.Peek()] <= quotes[i] {
			s.Pop()
		}
		if s.Empty() {
			spans[i] = i + 1
		} else {
			spans[i] = i - s.Peek()
		}
		s.Push(i)
	}
	return spans
}
