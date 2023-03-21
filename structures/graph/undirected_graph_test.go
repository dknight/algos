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

func TestUndirectedGraph(t *testing.T) {
	t.Run("NewUndirectedGraph()", func(t *testing.T) {
		graph := NewUndirectedGraph[int]()
		if graph == nil {
			t.Error("Cannot create an undirected graph")
		}
		if !graph.Empty() {
			t.Errorf("Newly create graph should be empty, got length %d",
				graph.Len())
		}
	})

	t.Run("AddEdge()", func(t *testing.T) {
		graph := NewUndirectedGraph[int]()
		graph.AddEdge(11, 22)
		graph.AddEdge(11, 33)
		exp := 3
		got := graph.Len()
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}

	})

	t.Run("DFS()", func(t *testing.T) {
		graph := NewUndirectedGraph[string]()
		graph.AddEdge("A", "B")
		graph.AddEdge("A", "D")
		graph.AddEdge("B", "C")
		graph.AddEdge("B", "D")
		graph.AddEdge("C", "D")
		graph.AddEdge("C", "G")
		graph.AddEdge("D", "E")
		graph.AddEdge("E", "D")
		graph.AddEdge("E", "F")
		graph.AddEdge("F", "G")
		graph.AddEdge("G", "H")
		graph.AddEdge("H", "I")
		graph.AddEdge("H", "J")
		graph.AddEdge("J", "K")

		exp := []string{"A", "B", "C", "G", "H", "I", "J", "K", "D", "E", "F"}
		got := graph.DFS("A")
		if !reflect.DeepEqual(exp, got) {
			t.Errorf("Expected %v got %v", exp, got)
		}
		// fmt.Println(got)

		graph = NewUndirectedGraph[string]()
		assertPanic(t, func() {
			got = graph.DFS("A")
		})
	})

	t.Run("BFS()", func(t *testing.T) {
		graph := NewUndirectedGraph[string]()
		graph.AddEdge("A", "B")
		graph.AddEdge("A", "D")
		graph.AddEdge("B", "C")
		graph.AddEdge("B", "D")
		graph.AddEdge("C", "D")
		graph.AddEdge("C", "G")
		graph.AddEdge("D", "E")
		graph.AddEdge("E", "D")
		graph.AddEdge("E", "F")
		graph.AddEdge("F", "G")
		graph.AddEdge("G", "H")
		graph.AddEdge("H", "I")
		graph.AddEdge("H", "J")
		graph.AddEdge("J", "K")

		exp := []string{"A", "D", "B", "E", "C", "F", "G", "H", "J", "I", "K"}
		got := graph.BFS("A")
		if !reflect.DeepEqual(exp, got) {
			t.Errorf("Expected %v got %v", exp, got)
		}
		// fmt.Println(got)

		graph = NewUndirectedGraph[string]()
		assertPanic(t, func() {
			got = graph.BFS("A")
		})
	})
}
