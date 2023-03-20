package deque

import (
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

func TestDeque(t *testing.T) {
	t.Run("New()", func(t *testing.T) {
		d := New[int]()
		if d == nil {
			t.Error("Newly create queue should not be nil")
		}
		if !d.Empty() {
			t.Errorf("Newly create queue should be empty, got length %d",
				d.Len())
		}
	})

	t.Run("PushBack()", func(t *testing.T) {
		q := New[int]()
		q.PushBack(11)
		q.PushBack(22)
		q.PushBack(33)
		exp := 3
		got := q.Len()
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}
	})

	t.Run("PopBack()", func(t *testing.T) {
		d := New[int]()
		d.PushBack(11)
		d.PushBack(22)
		d.PushBack(33)
		exp := 3
		got := d.Len()
		if exp != got {
			t.Errorf("Expected %v got %v", 2, got)
		}

		dElem := d.PopBack()
		if dElem != 33 {
			t.Errorf("Expected %v got %v", 33, dElem)
		}
		dElem = d.PopBack()
		if dElem != 22 {
			t.Errorf("Expected %v got %v", 22, dElem)
		}
		dElem = d.PopBack()
		if dElem != 11 {
			t.Errorf("Expected %v got %v", 11, dElem)
		}
		assertPanic(t, func() {
			dElem = d.PopBack()
		})
	})

	t.Run("PushBack()", func(t *testing.T) {
		q := New[int]()
		q.PushFront(11)
		q.PushFront(22)
		q.PushFront(33)
		exp := 3
		got := q.Len()
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}
	})

	t.Run("PopFront()", func(t *testing.T) {
		d := New[int]()
		d.PushFront(11)
		d.PushFront(22)
		d.PushFront(33)
		exp := 3
		got := d.Len()
		if exp != got {
			t.Errorf("Expected %v got %v", 2, got)
		}

		dElem := d.PopFront()
		if dElem != 33 {
			t.Errorf("Expected %v got %v", 33, dElem)
		}
		dElem = d.PopFront()
		if dElem != 22 {
			t.Errorf("Expected %v got %v", 22, dElem)
		}
		dElem = d.PopFront()
		if dElem != 11 {
			t.Errorf("Expected %v got %v", 11, dElem)
		}
		assertPanic(t, func() {
			dElem = d.PopFront()
		})
	})

	t.Run("Front()", func(t *testing.T) {
		d := New[int]()
		d.PushFront(11)
		d.PushFront(22)
		d.PushFront(33)
		d.PushBack(44)
		exp := 33
		got := d.Front()
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}
	})

	t.Run("Back()", func(t *testing.T) {
		d := New[int]()
		d.PushFront(11)
		d.PushFront(22)
		d.PushFront(33)
		exp := 11
		got := d.Back()
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}
		d.PushBack(44)
		exp = 44
		got = d.Back()
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}
	})
}
