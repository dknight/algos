package main

import (
	"container/list"

	"github.com/dknight/algos/stack"
)

func SimpleStocksSpan(quotes []int) []int {
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

func StackStocksSpan(quotes []int) []int {
	n := len(quotes)
	spans := make([]int, n)
	s := stack.New[int](n)

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

func StackSlicesStocksSpan(quotes []int) []int {
	n := len(quotes)
	spans := make([]int, n)
	s := make([]int, n)
	length := 0

	spans[0] = 1
	s[0] = 0
	length++

	for i := 1; i < n; i++ {
		for length != 0 && quotes[s[length-1]] <= quotes[i] {
			s = s[:len(s)-1]
			length--
		}
		if length == 0 {
			spans[i] = i + 1
		} else {
			spans[i] = i - s[length-1]
		}
		s[length] = i
		length++
	}
	return spans
}

func ListStocksSpan(quotes []int) []int {
	n := len(quotes)
	spans := make([]int, n)
	l := list.New()

	spans[0] = 1
	l.PushBack(0)

	for i := 1; i < n; i++ {
		for l.Len() != 0 && quotes[l.Back().Value.(int)] <= quotes[i] {
			elem := l.Back()
			l.Remove(elem)
		}
		if l.Len() == 0 {
			spans[i] = i + 1
		} else {
			spans[i] = i - l.Back().Value.(int)
		}
		l.PushBack(i)
	}
	return spans
}
