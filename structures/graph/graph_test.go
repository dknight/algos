package graph

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

func TestGraph(t *testing.T) {
	graph := NewGraph[string]()
	graph.AddEdge("A", "B")
	graph.AddEdge("A", "D")
	graph.AddEdge("B", "A")
	graph.AddEdge("B", "C")
	graph.AddEdge("B", "D")
	graph.AddEdge("C", "B")
	graph.AddEdge("C", "D")
	graph.AddEdge("C", "G")
	graph.AddEdge("D", "A")
	graph.AddEdge("D", "B")
	graph.AddEdge("D", "C")
	graph.AddEdge("D", "E")
	graph.AddEdge("E", "D")
	graph.AddEdge("E", "F")
	graph.AddEdge("F", "E")
	graph.AddEdge("F", "G")
	graph.AddEdge("G", "C")
	graph.AddEdge("G", "F")
	graph.AddEdge("G", "H")
	graph.AddEdge("H", "G")
	graph.AddEdge("H", "I")
	graph.AddEdge("H", "J")
	graph.AddEdge("I", "H")
	graph.AddEdge("J", "H")
	graph.AddEdge("J", "K")
	graph.AddEdge("K", "J")

	t.Run("NewGraph()", func(t *testing.T) {
		graph := NewGraph[int]()
		if graph == nil {
			t.Error("Cannot create an graph")
		}
		if !graph.Empty() {
			t.Errorf("Newly create graph should be empty, got length %d",
				graph.Len())
		}
	})

	t.Run("AddEdge()", func(t *testing.T) {
		graph := NewGraph[int]()
		graph.AddEdge(11, 22)
		graph.AddEdge(11, 33)
		exp := 3
		got := graph.Len()
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}
	})

	t.Run("DFS()", func(t *testing.T) {
		exp := []string{"A", "B", "C", "G", "F", "E", "H", "I", "J", "K", "D"}
		got := graph.DFS("A")
		if !reflect.DeepEqual(exp, got) {
			t.Errorf("Expected %v got %v", exp, got)
		}

		graphEmpty := NewGraph[string]()
		assertPanic(t, func() {
			got = graphEmpty.DFS("A")
		})
	})

	t.Run("BFS()", func(t *testing.T) {
		exp := []string{"A", "D", "B", "E", "C", "F", "G", "H", "J", "I", "K"}
		got := graph.BFS("A")
		if !reflect.DeepEqual(exp, got) {
			t.Errorf("Expected %v got %v", exp, got)
		}

		graphEmpty := NewGraph[string]()
		assertPanic(t, func() {
			got = graphEmpty.BFS("A")
		})
	})
	t.Run("IsBipartite()", func(t *testing.T) {
		got := graph.IsBipartite()
		exp := false
		if got != exp {
			t.Errorf("Expected %v got %v", exp, got)
		}

		bipart := NewGraph[string]()
		bipart.AddEdge("A", "B")
		bipart.AddEdge("A", "C")
		bipart.AddEdge("B", "D")
		bipart.AddEdge("C", "D")
		bipart.AddEdge("D", "E")
		bipart.AddEdge("E", "F")
		bipart.AddEdge("E", "G")
		bipart.AddEdge("E", "H")
		bipart.AddEdge("G", "I")
		got = bipart.IsBipartite()
		exp = true
		if got != exp {
			t.Errorf("Expected %v got %v", exp, got)
		}
	})

}
