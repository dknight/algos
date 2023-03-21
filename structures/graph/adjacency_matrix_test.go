package graph

import (
	"fmt"
	"testing"
)

func TestAdjancencyMatrix(t *testing.T) {
	t.Run("AdjancencyMatrix()", func(t *testing.T) {
		gr := NewUndirectedGraph[string]()
		gr.AddEdge("A", "B")
		gr.AddEdge("A", "C")
		gr.AddEdge("C", "D")
		am := gr.AdjacencyMatrix()
		got := am.String()
		// FIXME sort issue
		// exp := "A B C D\n----------\nA 0 1 1 0\nB 0 0 0 0\nC 0 0 0 1\nD 0 0 0 0\n"
		// if exp != got {
		// 	t.Errorf("Expected %v got %v", exp, got)
		// }
		fmt.Println(got)
	})
}
