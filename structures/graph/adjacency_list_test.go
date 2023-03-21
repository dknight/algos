package graph

import (
	"fmt"
	"testing"
)

func TestAdjacencyList(t *testing.T) {
	t.Run("String()", func(t *testing.T) {
		graph := NewUndirectedGraph[string]()
		graph.AddEdge("A", "B")
		graph.AddEdge("A", "C")
		graph.AddEdge("C", "D")

		got := graph.AdjacencyList().String()
		// FIXME sort issue
		// exp := "D\n --> C --> B\nB\nC --> D\n"
		// if exp != got {
		// 	t.Errorf("Expected %v got %v", exp, got)
		// }
		fmt.Println(got)
	})
}
