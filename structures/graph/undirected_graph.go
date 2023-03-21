package graph

import (
	"fmt"

	"github.com/dknight/algos/structures/list/llist"
	"github.com/dknight/algos/structures/queue/queue"
	"github.com/dknight/algos/structures/set"
	"github.com/dknight/algos/structures/stack"
)

// UndirectedGraph represents the undirected graph built with adjacency list.
type UndirectedGraph[T comparable] struct {
	AdjacencyList map[T]*llist.LinkedList[T]
	Vertices      set.Set[T]
}

// NewUndirectedGraph creates a new undirected graph and returns the pointer
// to it.
func NewUndirectedGraph[T comparable]() *UndirectedGraph[T] {
	return &UndirectedGraph[T]{
		AdjacencyList: map[T]*llist.LinkedList[T]{},
		Vertices:      set.New[T](),
	}
}

// AddEdge adds 2 vertices and an edge between them. Returns true if the
// edge is create successfully, false on failure.
func (g *UndirectedGraph[T]) AddEdge(u, v T) {
	g.addVertices(u, v)
	g.AdjacencyList[u].Push(v)
}

// DFS (Deep-First Search) is an algorithm for searching all the vertices of
// a graph or tree data structure. Traversal means visiting all the nodes
// of a graph. Returns slices of vertices in the order or visiting.
func (g *UndirectedGraph[T]) DFS(c T) []T {
	if !g.Vertices.Contains(c) {
		err := fmt.Sprintf("Vertex %v isn't present in graph", c)
		panic(err)
	}
	stk := stack.New[T]()
	visited := set.New[T]()
	instack := set.New[T]()
	path := make([]T, 0)

	stk.Push(c)
	instack.Add(c)

	for !stk.Empty() {
		c = stk.Pop()
		instack.Remove(c)
		visited.Add(c)
		li := g.AdjacencyList[c]
		path = append(path, c)
		for node := li.Head(); node != nil; node = node.Next() {
			v := node.Value
			if !visited.Contains(v) && !instack.Contains(v) {
				stk.Push(v)
				instack.Add(v)
			}
		}
	}
	return path
}

// BFS (Breadth-First Search) is an algorithm for searching all the vertices of
// a graph or tree data structure. Traversal means visiting all the nodes
// of a graph. Returns slices of vertices in the order or visiting.
func (g *UndirectedGraph[T]) BFS(c T) []T {
	if !g.Vertices.Contains(c) {
		err := fmt.Sprintf("Vertex %v isn't present in graph", c)
		panic(err)
	}
	que := queue.New[T]()
	visited := set.New[T]()
	inqueue := set.New[T]()
	path := make([]T, 0)

	que.Enqueue(c)
	inqueue.Add(c)

	for !que.Empty() {
		c = que.Dequeue()
		inqueue.Remove(c)
		visited.Add(c)
		li := g.AdjacencyList[c]
		path = append(path, c)
		for node := li.Head(); node != nil; node = node.Next() {
			v := node.Value
			if !visited.Contains(v) && !inqueue.Contains(v) {
				que.Enqueue(v)
				inqueue.Add(v)
			}
		}
	}
	return path
}

// Len gets number of vertices in the undirected graph.
func (g *UndirectedGraph[T]) Len() int {
	return g.Vertices.Len()
}

// Empty checks is graph is empty, has not any vertices).
func (g *UndirectedGraph[T]) Empty() bool {
	return g.Vertices.Len() == 0
}

// addVertices adds vertices to the graph, also create adjacency list for each
// vertex. Later adjacency list can be empty, if vertex does not connect to any
// other vertices.
func (g *UndirectedGraph[T]) addVertices(vs ...T) {
	for _, v := range vs {
		if _, ok := g.AdjacencyList[v]; !ok {
			g.AdjacencyList[v] = llist.New[T]()
		}
		g.Vertices.Add(v)
	}
}
