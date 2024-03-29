package stack

import (
	"reflect"
	"testing"
)

func assertPanic(t *testing.T, fn func()) {
	t.Helper()
	defer func() {
		if err := recover(); err == nil {
			t.Error(err)
		}
	}()
	fn()
}

func TestStack(t *testing.T) {
	t.Run("New()", func(t *testing.T) {
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
		s.Push(1)
		s.Push(2)
		s.Push(3)
		got := s.Len()
		exp := 3
		if exp != got {
			t.Errorf("Expected %v got %v\n", exp, got)
		}
	})

	t.Run("Peek()", func(t *testing.T) {
		t.Run("Common", func(t *testing.T) {
			s := New[string]()
			s.Push("boo")
			s.Push("moo")
			exp := "moo"
			got := s.Peek()
			if exp != got {
				t.Errorf("Expected %v got %v\n", exp, got)
			}
		})

		t.Run("Empty", func(t *testing.T) {
			s := New[string]()
			assertPanic(t, func() { s.Peek() })
		})
	})

	t.Run("Pop()", func(t *testing.T) {
		s := New[string]()
		s.Push("boo")
		s.Push("moo")
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

	t.Run("Reset()", func(t *testing.T) {
		s := New[int]()
		s.Push(1)
		s.Push(2)
		s.Push(3)
		s.Push(4)
		s.Push(5)
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

	t.Run("Dump()", func(t *testing.T) {
		s := New[int]()
		s.Push(1)
		s.Push(2)
		s.Push(3)
		s.Push(4)
		s.Push(5)
		exp := []int{1, 2, 3, 4, 5}
		got := s.Dump()
		if !reflect.DeepEqual(exp, got) {
			t.Errorf("Expected %v got %v\n", exp, got)
		}
	})

	t.Run("String()", func(t *testing.T) {
		s := New[int]()
		s.Push(11)
		s.Push(22)
		s.Push(33)
		s.Push(44)
		exp := "[11 22 33 44]"
		got := s.String()
		if exp != got {
			t.Errorf("Expected %v got %v\n", exp, got)
		}
	})
}
