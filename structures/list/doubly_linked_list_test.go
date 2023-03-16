package list

import (
	"testing"
)

func TestList(t *testing.T) {
	t.Run("TestListNew()", func(t *testing.T) {
		li := NewDoublyLikedList[int]()
		if li == nil {
			t.Error("List should not be nil")
		}
	})
	t.Run("TestListLen()", func(t *testing.T) {
		li := NewDoublyLikedList[int]()
		if li.Len() != 0 {
			t.Error("List should not be nil")
		}
	})

	t.Run("TestListBack_Empty()", func(t *testing.T) {
		li := NewDoublyLikedList[int]()
		got := li.Back()
		if got != nil {
			t.Errorf("Expected %v got %v", nil, got)
		}
	})
	t.Run("TestListBack()", func(t *testing.T) {
		li := NewDoublyLikedList[int]()
		li.PushBack(111)
		expBack := 111
		gotBack := li.Back()
		if gotBack.Value() != expBack {
			t.Errorf("Expected %v got %v", expBack, gotBack.Value())
		}
		gotBack = li.PushBack(222)
		expBack = 222
		if gotBack.Value() != expBack {
			t.Errorf("Expected %v got %v", expBack, gotBack.Value())
		}

		gotFront := li.Front()
		expFront := 111
		if gotFront.Value() != expFront {
			t.Errorf("Expected %v got %v", expFront, gotFront.Value())
		}
	})

	t.Run("TestListFront_Empty()", func(t *testing.T) {
		li := NewDoublyLikedList[int]()
		got := li.Front()
		if got != nil {
			t.Errorf("Expected %v got %v", nil, got)
		}
	})
	t.Run("TestListFront()", func(t *testing.T) {
		li := NewDoublyLikedList[int]()
		li.PushFront(111)
		expFront := 111
		expBack := 111
		gotFront := li.Front()
		gotBack := li.Back()
		if gotFront.Value() != expFront {
			t.Errorf("Expected %v got %v", expFront, gotFront.Value())
		}
		if gotBack.Value() != expBack {
			t.Errorf("Expected %v got %v", expBack, gotBack.Value())
		}

		gotFront = li.PushFront(222)
		expBack = 111
		expFront = 222
		gotFront = li.Front()
		gotBack = li.Back()
		if gotFront.Value() != expFront {
			t.Errorf("Expected %v got %v", expFront, gotFront.Value())
		}
		if gotBack.Value() != expBack {
			t.Errorf("Expected %v got %v", expBack, gotBack.Value())
		}
	})

	t.Run("NodeNext()", func(t *testing.T) {
		ranks := []string{
			"private", "corporal", "sergeant", "lieutenant",
		}
		li := NewDoublyLikedList[string]()
		for i := 0; i < len(ranks); i++ {
			li.PushBack(ranks[i])
		}
		for i, node := 0, li.Front(); node != nil; i, node = i+1, node.Next() {
			exp := ranks[i]
			got := node.Value()
			if exp != got {
				t.Errorf("Expected %v got %v", exp, got)
			}
		}
		for i, node := li.Len()-1, li.Back(); node != nil; i, node = i-1, node.Prev() {
			exp := ranks[i]
			got := node.Value()
			if exp != got {
				t.Errorf("Expected %v got %v", exp, got)
			}
		}
	})
	t.Run("NodePrev()", func(t *testing.T) {
		ranks := []string{
			"private", "corporal", "sergeant", "lieutenant",
		}
		li := NewDoublyLikedList[string]()
		for i := 0; i < len(ranks); i++ {
			li.PushFront(ranks[i])
		}
		for i, nod := 0, li.Back(); nod != nil; i, nod = i+1, nod.Prev() {
			exp := ranks[i]
			got := nod.Value()
			if exp != got {
				t.Errorf("Expected %v got %v", exp, got)
			}
		}
		for i, nod := li.Len()-1, li.Front(); nod != nil; i, nod = i-1, nod.Next() {
			exp := ranks[i]
			got := nod.Value()
			if exp != got {
				t.Errorf("Expected %v got %v", exp, got)
			}
		}
	})

	t.Run("FindByValue()", func(t *testing.T) {
		ranks := []string{
			"private", "corporal", "sergeant", "lieutenant",
		}
		li := NewDoublyLikedList[string]()
		for i := 0; i < len(ranks); i++ {
			li.PushFront(ranks[i])
		}
		exp := ranks[2]
		found := li.FindByValue(ranks[2])
		got := found.Value()
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}
	})
	t.Run("FindByValue_Empty()", func(t *testing.T) {
		li := NewDoublyLikedList[string]()
		got := li.FindByValue("Bambarbia")
		if got != nil {
			t.Errorf("Expected %v got %v", nil, got)
		}
	})

	t.Run("Remove()", func(t *testing.T) {
		li := NewDoublyLikedList[string]()
		li.PushBack("Hello")
		li.PushBack("World")
		if li.Len() != 2 {
			t.Errorf("Expected %v got %v", 2, li.Len())
		}

		found := li.FindByValue("Hello")
		ok := li.Remove(found)
		if li.Len() != 1 || !ok {
			t.Errorf("Expected %v got %v", 1, li.Len())
		}
		ok = li.Remove(nil)
		if ok {
			t.Errorf("Expected %v got %v", false, ok)
		}
	})
}
