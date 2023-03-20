package llist

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	t.Run("New()", func(t *testing.T) {
		li := New[int]()
		if li == nil {
			t.Error("List should not be nil")
		}
	})

	t.Run("Len()", func(t *testing.T) {
		li := New[int]()
		exp := 0
		got := li.Len()
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}
	})

	t.Run("Append()", func(t *testing.T) {
		li := New[int]()
		li.Append(1)
		exp := 1
		got := li.Tail().Value
		length := li.Len()
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}
		if length != 1 {
			t.Errorf("Expected %v got %v", 1, length)
		}

		li.Append(2)
		exp = 2
		got = li.Tail().Value
		length = li.Len()
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}
		if length != 2 {
			t.Errorf("Expected %v got %v", 2, length)
		}
	})

	t.Run("Push()", func(t *testing.T) {
		li := New[int]()

		li.Push(1)
		exp := 1
		got := li.Head().Value
		length := li.Len()
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}
		if length != 1 {
			t.Errorf("Expected %v got %v", 1, length)
		}

		li.Push(2)
		exp = 2
		got = li.Head().Value
		length = li.Len()
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}
		if length != 2 {
			t.Errorf("Expected %v got %v", 2, length)
		}
	})

	t.Run("Head()", func(t *testing.T) {
		t.Run("Empty", func(t *testing.T) {
			li := New[int]()
			got := li.Head()
			if got != nil {
				t.Errorf("Expected %v got %v", nil, got)
			}
		})

		t.Run("Common", func(t *testing.T) {
			li := New[int]()
			li.Append(1)
			li.Append(2)
			li.Append(3)
			exp := 1
			got := li.Head()
			if exp != got.Value {
				t.Errorf("Expected %v got %v", exp, got.Value)
			}
		})
	})

	t.Run("Tail()", func(t *testing.T) {
		t.Run("Empty()", func(t *testing.T) {
			li := New[int]()
			got := li.Tail()
			if got != nil {
				t.Errorf("Expected %v got %v", nil, got)
			}
		})

		t.Run("Common", func(t *testing.T) {
			li := New[int]()
			li.Append(1)
			li.Append(2)
			li.Append(3)
			exp := 3
			got := li.Tail()
			if exp != got.Value {
				t.Errorf("Expected %v got %v", exp, got.Value)
			}
		})
	})

	t.Run("Empty()", func(t *testing.T) {
		li := New[int]()
		exp := true
		got := li.Empty()
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}
		li.Append(1)
		exp = false
		got = li.Empty()
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}
	})

	t.Run("RemoveByValue()", func(t *testing.T) {
		t.Run("empty", func(t *testing.T) {
			li := New[int]()
			got := li.RemoveByValue(2)
			if got != nil {
				t.Errorf("Expected %v got %v", nil, got)
			}
		})

		t.Run("Common", func(t *testing.T) {
			li := New[int]()
			li.Append(1)
			li.Append(2)
			li.Append(3)
			got := li.RemoveByValue(2)
			if got == nil {
				t.Errorf("Expected %v got %v", got, nil)
			}
			if li.Len() != 2 {
				t.Errorf("Expected %v got %v", 2, li.Len())
			}

			got = li.RemoveByValue(4)
			if got != nil {
				t.Errorf("Expected %v got %v", nil, got)
			}

			elem1 := li.Head()
			elem2 := elem1.Next()
			if elem1.Value != 1 {
				t.Errorf("Expected %v got %v", 1, elem1.Value)
			}
			if elem2.Value != 3 {
				t.Errorf("Expected %v got %v", 3, elem2.Value)
			}
			li.RemoveByValue(1)
			li.RemoveByValue(3)
			if li.Len() != 0 {
				t.Errorf("Expected %v got %v", 0, li.Len())
			}
		})
	})

	t.Run("Remove()", func(t *testing.T) {
		t.Run("Empty", func(t *testing.T) {
			li := New[int]()
			got := li.Remove(nil)
			if got != nil {
				t.Errorf("Expected %v got %v", nil, got)
			}
		})

		t.Run("Common", func(t *testing.T) {
			li := New[int]()
			li.Append(11)
			li.Append(22)
			li.Append(33)
			node := li.FindByValue(22)
			removed := li.Remove(node)
			if removed != node {
				t.Errorf("Expected %v got %v", removed, node)
			}
			if li.Len() != 2 {
				t.Errorf("Expected %v got %v", 2, li.Len())
			}

			node = li.FindByValue(11)
			removed = li.Remove(node)
			if removed != node {
				t.Errorf("Expected %v got %v", removed, node)
			}
			if li.Len() != 1 {
				t.Errorf("Expected %v got %v", 1, li.Len())
			}
			exp := 33
			got := li.Head().Value
			if got != exp {
				t.Errorf("Expected %v got %v", exp, got)
			}
			got = li.Tail().Value
			if got != exp {
				t.Errorf("Expected %v got %v", exp, got)
			}
		})
	})

	t.Run("FindByValue()", func(t *testing.T) {
		t.Run("Common", func(t *testing.T) {
			li := New[int]()
			li.Append(1)
			li.Append(2)
			li.Append(3)
			found := li.FindByValue(2)
			got := found.Value
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
			people := []Person{
				Person{
					Name: "Toby",
					Age:  12,
				},
				Person{
					Name: "Julia",
					Age:  27,
				},
				Person{
					Name: "Joe",
					Age:  34,
				},
			}
			li := New[Person]()
			for _, p := range people {
				li.Append(p)
			}
			exp := len(people)
			got := li.Len()
			if got != exp {
				t.Errorf("Expected %v got %v", exp, got)
			}

			li.RemoveByValue(people[1])
			exp = 2
			got = li.Len()
			if got != exp {
				t.Errorf("Expected %v got %v", exp, got)
			}

			li.RemoveByValue(Person{Name: "Joe", Age: 34})
			exp = 1
			got = li.Len()
			if got != exp {
				t.Errorf("Expected %v got %v", exp, got)
			}
		})
	})

	t.Run("InsertAfter()", func(t *testing.T) {
		t.Run("Empty", func(t *testing.T) {
			li := New[int]()
			node := &Node[int]{11, nil}
			inserted := li.InsertAfter(node, nil)
			exp := 11
			got := inserted.Value
			if got != exp {
				t.Errorf("Expected %v got %v", exp, got)
			}
		})

		t.Run("Common", func(t *testing.T) {
			nodes := []*Node[int]{
				&Node[int]{11, nil},
				&Node[int]{22, nil},
				&Node[int]{33, nil},
				&Node[int]{44, nil},
			}
			li := New[int]()
			got := li.InsertAfter(nodes[0], nil)
			exp := li.Head()
			if got != exp {
				t.Errorf("Expected %v got %v", exp, got)
			}

			got = li.InsertAfter(nodes[1], exp)
			exp = exp.Next()
			if got != exp {
				t.Errorf("Expected %v got %v", exp, got)
			}

			got = li.InsertAfter(nodes[2], exp)
			exp = exp.Next()
			if got != exp {
				t.Errorf("Expected %v got %v", exp, got)
			}

			got = li.InsertAfter(nodes[3], nodes[1])
			exp = nodes[1].Next()
			if got != exp {
				t.Errorf("Expected %v got %v", exp, got)
			}
		})
	})

	t.Run("InsertWithValue()", func(t *testing.T) {
		li := New[int]()
		node := li.InsertWithValue(11, li.head)
		exp := 11
		got := node.Value
		if got != exp {
			t.Errorf("Expected %v got %v", exp, got)
		}

		node = li.InsertWithValue(22, node)
		exp = 22
		got = node.Value
		if got != exp {
			t.Errorf("Expected %v got %v", exp, got)
		}
		got = li.Tail().Value
		if got != exp {
			t.Errorf("Expected %v got %v", exp, got)
		}
		got = li.Head().Value
		exp = 11
		if got != exp {
			t.Errorf("Expected %v got %v", exp, got)
		}

		node = li.InsertWithValue(33, li.head)
		exp = 33
		got = li.Head().Next().Value
		if got != exp {
			t.Errorf("Expected %v got %v", exp, got)
		}
		exp = 33
		got = li.Head().Next().Value
		if got != exp {
			t.Errorf("Expected %v got %v", exp, got)
		}
		// for n := li.Head(); n != nil; n = n.Next() {
		// 	fmt.Println(n)
		// }
	})
}

func TestNode(t *testing.T) {
	t.Run("Next()", func(t *testing.T) {
		ranks := []string{
			"private", "corporal", "sergeant", "lieutenant",
		}
		li := New[string]()
		for i := 0; i < len(ranks); i++ {
			li.Append(ranks[i])
		}
		for i, node := 0, li.Head(); node != nil; i, node = i+1, node.Next() {
			exp := ranks[i]
			got := node.Value
			if exp != got {
				t.Errorf("Expected %v got %v", exp, got)
			}
		}
	})
}
