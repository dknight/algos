package list

import (
	"reflect"
)

// Node represents the node in linked in.
type Node[T any] struct {
	value T
	next  *Node[T]
}

// Next is the nodes next iterator retrieves the next node in linked list.
func (n *Node[T]) Next() *Node[T] {
	return n.next
}

// Value gets the value of the node.
func (n *Node[T]) Value() T {
	return n.value
}

// LinkedList is a single linked list.
type LinkedList[T any] struct {
	head, tail *Node[T]
	length     int
}

// NewLinkedList create a new linked list.
func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// Len returns the length of the linked list.
func (li *LinkedList[T]) Len() int {
	return li.length
}

// Push adds the value to the beginning of the list.
func (li *LinkedList[T]) Push(v T) *Node[T] {
	node := &Node[T]{value: v}
	if li.head == nil {
		li.head = node
		li.tail = node
	} else {
		li.tail.next = node
		li.tail = node
	}
	li.length++
	return node
}

// Removes the first found node by value fron linked list.
func (li *LinkedList[T]) Remove(v T) *Node[T] {
	var prev *Node[T]
	node := li.head
	if node != nil && reflect.DeepEqual(node.Value(), v) {
		li.head = node.next
		li.length--
		return node
	}

	for node != nil && !reflect.DeepEqual(node.Value(), v) {
		prev = node
		node = node.next
	}

	if node == nil {
		return nil
	}

	prev.next = node.next
	li.length--
	return node
}

// Head returns the head (firstly added node) of the linked list.
func (li *LinkedList[T]) Head() *Node[T] {
	return li.head
}

// Tail returns the head (lastly added node) of the linked list.
func (li *LinkedList[T]) Tail() *Node[T] {
	return li.tail
}

// Empty check if the linked list is empty.
func (li *LinkedList[T]) Empty() bool {
	return li.length == 0
}
