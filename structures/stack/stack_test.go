package stack

import (
	"sync"
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("New()", func(t *testing.T) {
		t.Parallel()
		s := New[any]()
		if s == nil {
			t.Error("Cannot create stack")
		}
	})
	t.Run("Len()", func(t *testing.T) {
		s := New[any]()
		exp := 0
		got := s.Len()
		if exp != got {
			t.Errorf("Expected %v got %v\n", exp, got)
		}
	})
	t.Run("Push()", func(t *testing.T) {
		s := New[int]()
		s.Push(1, 2, 3)
		got := s.Len()
		exp := 3
		if exp != got {
			t.Errorf("Expected %v got %v\n", exp, got)
		}
	})
	t.Run("Peek()", func(t *testing.T) {
		s := New[string]()
		s.Push("boo", "moo")
		exp := "moo"
		got := s.Peek()
		if exp != got {
			t.Errorf("Expected %v got %v\n", exp, got)
		}

		// sEmpt := New[string]()
		// got = sEmpt.Peek()
	})
	t.Run("Pop()", func(t *testing.T) {
		s := New[string]()
		s.Push("boo", "moo")
		got := s.Pop()
		exp := "moo"
		if exp != got {
			t.Errorf("Expected %v got %v\n", exp, got)
		}

		exp = "boo"
		got = s.Pop()
		if exp != got {
			t.Errorf("Expected %v got %v\n", exp, got)
		}

		// got = s.Pop()
	})
	t.Run("Empty()", func(t *testing.T) {
		s := New[string]()
		exp := true
		got := s.Empty()
		if exp != got {
			t.Errorf("Expected %v got %v\n", exp, got)
		}

		s.Push("xxx")
		exp = false
		got = s.Empty()
		if exp != got {
			t.Errorf("Expected %v got %v\n", exp, got)
		}
	})
}

func TestStackAsync(t *testing.T) {
	t.Run("Async Push", func(t *testing.T) {
		n := 10
		s := New[int]()
		for i := 0; i < n; i++ {
			go func(i int) {
				s.Push(i)

			}(i)
		}
	})
	t.Run("Async Pop/Len", func(t *testing.T) {
		s := New[int]()
		var wg sync.WaitGroup
		wg.Add(s.Len())

		for i := 0; i < s.Len(); i++ {
			go func(i int) {
				s.Push(i)

			}(i)
		}

		for i := 0; i < s.Len(); i++ {
			go func() {
				s.Pop()
				wg.Done()
			}()
		}

		go func() {
			wg.Wait()
			exp := 0
			got := s.Len()
			if exp != got {
				t.Errorf("Expected %v got %v\n", exp, got)
			}
		}()
	})
}
