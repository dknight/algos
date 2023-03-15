package stack

import (
	"reflect"
	"testing"
)

func assertPanic(t *testing.T, fn func()) {
	t.Helper()
	defer func() { _ = recover() }()
	fn()
	t.Errorf("should have panicked")
}

func TestStack(t *testing.T) {
	t.Run("TestStackNew()", func(t *testing.T) {
		s := New[any]()
		if s == nil {
			t.Error("Cannot create stack")
		}
	})
	t.Run("TestStackLen()", func(t *testing.T) {
		s := New[any]()
		exp := 0
		got := s.Len()
		if exp != got {
			t.Errorf("Expected %v got %v\n", exp, got)
		}
	})
	t.Run("TestStackPush()", func(t *testing.T) {
		s := New[int]()
		s.Push(1, 2, 3)
		got := s.Len()
		exp := 3
		if exp != got {
			t.Errorf("Expected %v got %v\n", exp, got)
		}
	})
	t.Run("TestStackPeek()", func(t *testing.T) {
		s := New[string]()
		s.Push("boo", "moo")
		exp := "moo"
		got := s.Peek()
		if exp != got {
			t.Errorf("Expected %v got %v\n", exp, got)
		}
	})
	t.Run("TestStackEmptyPeek()", func(t *testing.T) {
		s := New[string]()
		assertPanic(t, func() { s.Peek() })
	})
	t.Run("TestStackPop()", func(t *testing.T) {
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

		assertPanic(t, func() { s.Pop() })
	})
	t.Run("TestStackEmpty()", func(t *testing.T) {
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
	t.Run("TestStackReset()", func(t *testing.T) {
		s := New[int]()
		s.Push(1, 2, 3, 4, 5)
		exp := 5
		got := s.Len()
		if exp != got {
			t.Errorf("Expected %v got %v\n", exp, got)
		}
		s.Reset()
		exp = 0
		got = s.Len()
		if exp != got {
			t.Errorf("Expected %v got %v\n", exp, got)
		}
	})
	t.Run("TestStackDump()", func(t *testing.T) {
		s := New[int]()
		s.Push(1, 2, 3, 4, 5)
		exp := []int{1, 2, 3, 4, 5}
		got := s.Dump()
		if !reflect.DeepEqual(exp, got) {
			t.Errorf("Expected %v got %v\n", exp, got)
		}
	})
	t.Run("TestStackString()", func(t *testing.T) {
		s := New[int]()
		s.Push(11, 22, 33, 44)
		exp := "[11 22 33 44]"
		got := s.String()
		if exp != got {
			t.Errorf("Expected %v got %v\n", exp, got)
		}
	})
}
