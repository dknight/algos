package graph

import (
	"fmt"
	"strings"

	"github.com/dknight/algos/structures/list/llist"
)

const arrSep = " -> "

// AdjacencyList represents the adjacency list of the undirected graph.
type AdjacencyList[T comparable] map[T]*llist.LinkedList[T]

// String beautifully prints adjacency list.
func (al AdjacencyList[T]) String() string {
	var body strings.Builder
	for k, li := range al {
		elems := []string{}
		elems = append(elems, fmt.Sprintf("%v", k))
		for node := li.Head(); node != nil; node = node.Next() {
			elems = append(elems, fmt.Sprintf("%v", node.Value))
		}
		body.WriteString(strings.Join(elems, arrSep))
		body.WriteString("\n")
	}
	return body.String()
}
