package graph

import (
	"fmt"

	"github.com/dknight/algos/structures/list/llist"
	"github.com/dknight/algos/structures/queue/queue"
	"github.com/dknight/algos/structures/set"
	"github.com/dknight/algos/structures/stack"
)

// Graph represents the graph built with adjacency list.
type Graph[T comparable] struct {
	adjList  AdjacencyList[T]
	vertices set.Set[T]
}

// NewGraph creates a new graph and returns the pointer
// to it.
func NewGraph[T comparable]() *Graph[T] {
	return &Graph[T]{
		adjList:  AdjacencyList[T]{},
		vertices: set.New[T](),
	}
}

// AddEdge adds 2 vertices and an edge between them. Returns true if the
// edge is create successfully, false on failure.
func (g *Graph[T]) AddEdge(u, v T) {
	g.addVertices(u, v)
	g.adjList[u].Push(v)
}

// DFS (Deep-First Search) is an algorithm for searching all the vertices of
// a graph or tree data structure. Traversal means visiting all the nodes
// of a graph. Returns slices of vertices in the order or visiting.
func (g *Graph[T]) DFS(c T) []T {
	if !g.Vertices().Contains(c) {
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
		li := g.adjList[c]
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
func (g *Graph[T]) BFS(c T) []T {
	if !g.Vertices().Contains(c) {
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
		li := g.adjList[c]
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

// Len gets number of vertices in the graph.
func (g *Graph[T]) Len() int {
	return g.Vertices().Len()
}

// Empty checks is graph is empty, has not any vertices).
func (g *Graph[T]) Empty() bool {
	return g.Vertices().Len() == 0
}

// addVertices adds vertices to the graph, also create adjacency list for each
// vertex. Later adjacency list can be empty, if vertex does not connect to any
// other vertices.
func (g *Graph[T]) addVertices(vs ...T) {
	for _, v := range vs {
		if _, ok := g.adjList[v]; !ok {
			g.adjList[v] = llist.New[T]()
		}
		g.Vertices().Add(v)
	}
}

// AdjancencyMatrix gets the adjacency matrix of the graph.
func (g *Graph[T]) AdjacencyMatrix() AdjacencyMatrix[T] {
	matrix := AdjacencyMatrix[T]{}
	for u := range g.Vertices() {
		matrix[u] = map[T]int{}
		for v := range g.Vertices() {
			matrix[u][v] = 0
			for w := g.adjList[u].Head(); w != nil; w = w.Next() {
				matrix[u][w.Value] = 1
			}
		}
	}
	return matrix
}

// IsBipartite checks is graph is bipartite.
func (g *Graph[T]) IsBipartite() bool {
	que := queue.New[T]()
	black := -1
	white := 0

	// "Paint" the graph
	colors := map[T]int{}
	for _, v := range g.Vertices().Keys() {
		colors[v] = black
	}

	start := g.Vertices().Keys()[0]
	colors[start] = white
	que.Enqueue(start)

	for !que.Empty() {
		v := que.Dequeue()
		li := g.adjList[v]
		for node := li.Head(); node != nil; node = node.Next() {
			u := node.Value
			if colors[u] == colors[v] {
				return false
			}
			if colors[u] == black {
				colors[u] = 1 - colors[v]
				que.Enqueue(u)
			}
		}
	}
	return true
}

// AdjacencyList gets adjacency list of the graph.
func (g *Graph[T]) AdjacencyList() AdjacencyList[T] {
	return g.adjList
}

// Vertices gets all vertices of the graph.
func (g *Graph[T]) Vertices() set.Set[T] {
	return g.vertices
}
