package list

import (
	"testing"
)

func TestList(t *testing.T) {
	t.Run("NewDoublyLikedList()", func(t *testing.T) {
		li := NewDoublyLikedList[int]()
		if li == nil {
			t.Error("List should not be nil")
		}
	})

	t.Run("DoublyLikedList_Len()", func(t *testing.T) {
		li := NewDoublyLikedList[int]()
		if li.Len() != 0 {
			t.Error("List should not be nil")
		}
	})

	t.Run("DoublyLikedList_Empty()", func(t *testing.T) {
		li := NewDoublyLikedList[int]()
		exp := true
		got := li.Empty()
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}
		li.PushBack(1)
		exp = false
		got = li.Empty()
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}
	})

	t.Run("DoublyLikedList_Back() _Empty()", func(t *testing.T) {
		li := NewDoublyLikedList[int]()
		got := li.Back()
		if got != nil {
			t.Errorf("Expected %v got %v", nil, got)
		}
	})

	t.Run("DoublyLikedList_Back()", func(t *testing.T) {
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

	t.Run("DoublyLinkedList_Front_Empty()", func(t *testing.T) {
		li := NewDoublyLikedList[int]()
		got := li.Front()
		if got != nil {
			t.Errorf("Expected %v got %v", nil, got)
		}
	})

	t.Run("DoublyLinkedList_Front()", func(t *testing.T) {
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

	t.Run("DoublyLinkedNode_Next()", func(t *testing.T) {
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

	t.Run("DoublyLinkedNode_Prev()", func(t *testing.T) {
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

	t.Run("DoublyLikedList_FindByValue()", func(t *testing.T) {
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

	t.Run("DoublyLikedList_FindByValue_Empty()", func(t *testing.T) {
		li := NewDoublyLikedList[string]()
		got := li.FindByValue("Bambarbia")
		if got != nil {
			t.Errorf("Expected %v got %v", nil, got)
		}
	})

	t.Run("DoublyLikedList_Remove()", func(t *testing.T) {
		li := NewDoublyLikedList[string]()
		li.PushBack("Hello")
		li.PushBack("World")
		if li.Len() != 2 {
			t.Errorf("Expected %v got %v", 2, li.Len())
		}

		found := li.FindByValue("Hello")
		node := li.Remove(found)
		if li.Len() != 1 && node == nil {
			t.Errorf("Expected %v got %v", 1, li.Len())
		}
		node = li.Remove(nil)
		if node != nil {
			t.Errorf("Expected %v got %v", false, node)
		}
	})

	t.Run("DoublyLikedList_RemoveFront()", func(t *testing.T) {
		li := NewDoublyLikedList[string]()
		li.PushBack("Hello")
		li.PushBack("Fluffy")
		li.PushBack("Kitty")

		node := li.RemoveFront()
		exp := "Hello"
		got := node.Value()
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}

		node = li.RemoveFront()
		exp = "Fluffy"
		got = node.Value()
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}

		node = li.RemoveFront()
		exp = "Kitty"
		got = node.Value()
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}

		node = li.RemoveFront()
		if node != nil {
			t.Errorf("Expected %v got %v", nil, node)
		}

		if li.Len() != 0 {
			t.Errorf("Expected %v got %v", 0, li.Len())
		}
	})

	t.Run("DoublyLikedList_RemoveBack()", func(t *testing.T) {
		li := NewDoublyLikedList[string]()
		li.PushBack("Hello")
		li.PushBack("Fluffy")
		li.PushBack("Kitty")

		node := li.RemoveBack()
		exp := "Kitty"
		got := node.Value()
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}

		node = li.RemoveBack()
		exp = "Fluffy"
		got = node.Value()
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}

		node = li.RemoveBack()
		exp = "Hello"
		got = node.Value()
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}

		node = li.RemoveBack()
		if node != nil {
			t.Errorf("Expected %v got %v", nil, node)
		}

		if li.Len() != 0 {
			t.Errorf("Expected %v got %v", 0, li.Len())
		}
	})
}
