package set

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestSet(t *testing.T) {
	t.Run("New()", func(t *testing.T) {
		set := New[int]()
		if set == nil {
			t.Error("Cannot create set")
		}
		if !set.Empty() {
			t.Error("Newly created set should be empty, got length", set.Len())
		}
	})

	t.Run("Add()", func(t *testing.T) {
		set := New[int]()
		exp := 0
		got := set.Len()
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}
		set.Add(11)
		exp = 1
		got = set.Len()
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}
		set.Add(22)
		exp = 2
		got = set.Len()
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}
		set.Add(22)
		exp = 2
		got = set.Len()
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}
	})

	t.Run("Delete()", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}
		set := New[Person]()
		p1 := Person{"Bill", 33}
		p2 := Person{"Jennifer", 19}
		set.Add(p1)
		set.Add(p2)
		got := set.Len()
		exp := 2
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}

		set.Delete(p2)
		got = set.Len()
		exp = 1
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}
		if _, ok := set[p2]; ok {
			t.Errorf("Expected %v got %v", false, ok)
		}
	})

	t.Run("Has()", func(t *testing.T) {
		set := New[string]()
		set.Add("Foo")
		set.Add("Bar")
		got := set.Has("Bar")
		exp := true
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}
		got = set.Has("Jar")
		exp = false
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}
	})

	t.Run("Clear()", func(t *testing.T) {
		set := New[string]()
		set.Add("Foo")
		set.Add("Bar")
		got := set.Len()
		exp := 2
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}

		set.Clear()
		got = set.Len()
		exp = 0
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}
	})

	t.Run("All()", func(t *testing.T) {
		set := New[string]()
		set.Add("Foo")
		set.Add("Bar")
		exp := []string{"Bar", "Foo"}
		got := set.All()
		sort.Slice(got, func(i, j int) bool {
			return got[i] < got[j]
		})
		if !reflect.DeepEqual(got, exp) {
			t.Errorf("Expected %v got %v", exp, got)
		}
	})

	t.Run("String()", func(t *testing.T) {
		set := New[int]()
		exp := ""
		got := set.String()
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}
		set.Add(11)
		set.Add(22)
		set.Add(33)
		set.Add(44)
		set.Add(55)
		gotLen := len(strings.Fields(set.String()))
		expLen := 5
		if expLen != gotLen {
			t.Errorf("Expected %v got %v", expLen, gotLen)
		}
	})
}
