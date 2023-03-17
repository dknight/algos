package list

import "testing"

func TestLinkedList(t *testing.T) {
	t.Run("New()", func(t *testing.T) {
		li := NewLinkedList[int]()
		if li == nil {
			t.Error("List should not be nil")
		}
	})

	t.Run("Len()", func(t *testing.T) {
		li := NewLinkedList[int]()
		exp := 0
		got := li.Len()
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}
	})

	t.Run("Push()", func(t *testing.T) {
		t.Skip()
		li := NewLinkedList[int]()
		li.Push(1)
		li.Push(2)
		exp := 2
		got := li.Len()
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}
	})

	t.Run("Head_Empty()", func(t *testing.T) {
		li := NewLinkedList[int]()
		got := li.Head()
		if got != nil {
			t.Errorf("Expected %v got %v", nil, got)
		}
	})

	t.Run("Head()", func(t *testing.T) {
		li := NewLinkedList[int]()
		li.Push(1)
		li.Push(2)
		li.Push(3)
		exp := 1
		got := li.Head()
		if exp != got.Value() {
			t.Errorf("Expected %v got %v", exp, got.Value())
		}
	})

	t.Run("Tail_Empty()", func(t *testing.T) {
		li := NewLinkedList[int]()
		got := li.Tail()
		if got != nil {
			t.Errorf("Expected %v got %v", nil, got)
		}
	})

	t.Run("Tail()", func(t *testing.T) {
		li := NewLinkedList[int]()
		li.Push(1)
		li.Push(2)
		li.Push(3)
		exp := 3
		got := li.Tail()
		if exp != got.Value() {
			t.Errorf("Expected %v got %v", exp, got.Value())
		}
	})

	t.Run("Node_Next()", func(t *testing.T) {
		ranks := []string{
			"private", "corporal", "sergeant", "lieutenant",
		}
		li := NewLinkedList[string]()
		for i := 0; i < len(ranks); i++ {
			li.Push(ranks[i])
		}
		for i, node := 0, li.Head(); node != nil; i, node = i+1, node.Next() {
			exp := ranks[i]
			got := node.Value()
			if exp != got {
				t.Errorf("Expected %v got %v", exp, got)
			}
		}
	})

	t.Run("Empty()", func(t *testing.T) {
		li := NewLinkedList[int]()
		exp := true
		got := li.Empty()
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}
		li.Push(1)
		exp = false
		got = li.Empty()
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}
	})

	t.Run("Remove()_Empty", func(t *testing.T) {
		li := NewLinkedList[int]()
		got := li.Remove(2)
		if got != nil {
			t.Errorf("Expected %v got %v", nil, got)
		}
	})

	t.Run("Remove()", func(t *testing.T) {
		li := NewLinkedList[int]()
		li.Push(1)
		li.Push(2)
		li.Push(3)
		got := li.Remove(2)
		if got == nil {
			t.Errorf("Expected %v got %v", got, nil)
		}
		if li.Len() != 2 {
			t.Errorf("Expected %v got %v", 2, li.Len())
		}

		got = li.Remove(4)
		if got != nil {
			t.Errorf("Expected %v got %v", nil, got)
		}

		elem1 := li.Head()
		elem2 := elem1.Next()
		if elem1.Value() != 1 {
			t.Errorf("Expected %v got %v", 1, elem1.Value())
		}
		if elem2.Value() != 3 {
			t.Errorf("Expected %v got %v", 3, elem2.Value())
		}
		li.Remove(1)
		li.Remove(3)
		if li.Len() != 0 {
			t.Errorf("Expected %v got %v", 0, li.Len())
		}
	})

	t.Run("SearchByValue()", func(t *testing.T) {
		li := NewLinkedList[int]()
		li.Push(1)
		li.Push(2)
		li.Push(3)
		found := li.FindByValue(2)
		got := found.Value()
		exp := 2
		if got != exp {
			t.Errorf("Expected %v got %v", exp, got)
		}

		found = li.FindByValue(4)
		if found != nil {
			t.Errorf("Expected %v got %v", nil, got)
		}
	})

	t.Run("structs", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}
		p := Person{"Dima", 37}
		li := NewLinkedList[Person]()
		li.Push(p)
		li.Remove(p)
		exp := 0
		got := li.Len()
		if got != exp {
			t.Errorf("Expected %v got %v", exp, got)
		}
	})
}
