package queue

import "testing"

func assertPanic(t *testing.T, fn func()) {
	t.Helper()
	defer func() {
		if err := recover(); err == nil {
			t.Error(err)
		}
	}()
	fn()
}

func TestQueue(t *testing.T) {
	t.Run("New()", func(t *testing.T) {
		q := New[int]()
		if q == nil {
			t.Error("Newly create queue should not be nil")
		}
		if !q.Empty() {
			t.Errorf("Newly create queue should be empty, got length %d",
				q.Len())
		}
	})

	t.Run("Enqueue()", func(t *testing.T) {
		q := New[int]()
		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)
		exp := 3
		got := q.Len()
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}
	})

	t.Run("Dequeue()", func(t *testing.T) {
		q := New[int]()
		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)
		q.Dequeue()
		dElem := q.Dequeue()
		exp := 1
		got := q.Len()
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}
		if dElem != 2 {
			t.Errorf("Expected %v got %v", 2, dElem)
		}

		dElem = q.Dequeue()
		if dElem != 3 {
			t.Errorf("Expected %v got %v", 3, dElem)
		}
		assertPanic(t, func() {
			dElem = q.Dequeue()
		})
	})

	t.Run("Head()", func(t *testing.T) {
		q := New[int]()
		q.Enqueue(11)
		q.Enqueue(22)
		q.Enqueue(33)
		exp := 11
		got := q.Head()
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}
	})

	t.Run("Tail()", func(t *testing.T) {
		q := New[int]()
		q.Enqueue(11)
		q.Enqueue(22)
		q.Enqueue(33)
		exp := 33
		got := q.Tail()
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}
	})
}
