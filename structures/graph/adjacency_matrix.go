package graph

import (
	"fmt"
	"math"
	"strings"
)

// AdjacencyMatrix represents adjacency matrix of the graph.
//
// TODO FIXME: In Go 1.20 try to sort generic type. At this moment vertices
// can be in undetermined order.
type AdjacencyMatrix[T comparable] map[T]map[T]int

// String beautifully prints adjacency matrix.
func (matrix AdjacencyMatrix[T]) String() string {
	var head strings.Builder
	var body strings.Builder
	format := ""
	valLen := 0

	// Collect keys
	keys := []T{}
	for k, _ := range matrix {
		keys = append(keys, k)
		valLen = int(math.Max(
			float64(valLen),
			float64(len(fmt.Sprintf("%v", k))+1),
		))
	}

	head.WriteString(strings.Repeat(" ", valLen))

	// FIXME sort keys?
	// sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	format = fmt.Sprintf("%%%dv", valLen)
	for _, k := range keys {
		body.WriteString(fmt.Sprintf(format, k))
		head.WriteString(fmt.Sprintf(format, k))
		for _, i := range keys {
			body.WriteString(fmt.Sprintf(format, matrix[k][i]))
		}
		body.WriteString("\n")
	}
	head.WriteString("\n")
	head.WriteString(strings.Repeat("-", len(matrix)*valLen+valLen))
	head.WriteString("\n")
	return head.String() + body.String()
}
