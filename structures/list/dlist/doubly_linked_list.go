package dlist

// Node of the doubly linked list.
// TODO comparable may not be equal of using structs.
type Node[T comparable] struct {
	next, prev *Node[T]
	list       *DoublyLinkedList[T]
	Value      T
}

// Next is the next iterator for the node. Retrieves the next node.
func (n *Node[T]) Next() *Node[T] {
	if nxt := n.next; n.list != nil && nxt != &n.list.root {
		return nxt
	}
	return nil
}

// Prev is the previous iterator for the node. Retrieves the previous node.
func (n *Node[T]) Prev() *Node[T] {
	if prv := n.prev; n.list != nil && prv != &n.list.root {
		return prv
	}
	return nil
}

// DoublyLinkedList represents the doubly linked list.
//
// NOTE This list is created only for educational purposes and not recommenced
// to use in production. If you need the linked lists structures it is better
// to use https://pkg.go.dev/container/list
type DoublyLinkedList[T comparable] struct {
	root   Node[T]
	length int
}

// New create a new doubly linked list.
func New[T comparable]() *DoublyLinkedList[T] {
	li := &DoublyLinkedList[T]{}
	li.root.next = &li.root
	li.root.prev = &li.root
	return li
}

// PushBack pushes the value to the end of the list.
func (li *DoublyLinkedList[T]) PushBack(val T) *Node[T] {
	node := &Node[T]{
		Value: val,
		list:  li,
	}
	return li.insertValue(node, li.root.prev)
}

// PushFront pushes the value to the beginning of the list.
func (li *DoublyLinkedList[T]) PushFront(val T) *Node[T] {
	node := &Node[T]{
		Value: val,
		list:  li,
	}
	return li.insertValue(node, &li.root)
}

func (li *DoublyLinkedList[T]) insertValue(node, at *Node[T]) *Node[T] {
	node.prev = at
	node.next = at.next
	node.prev.next = node
	node.next.prev = node
	li.length++
	return node
}

// Front gets the node from the front of the list.
func (li *DoublyLinkedList[T]) Front() *Node[T] {
	if li.length == 0 {
		return nil
	}
	return li.root.next
}

// Back gets the node from the back of the list.
func (li *DoublyLinkedList[T]) Back() *Node[T] {
	if li.length == 0 {
		return nil
	}
	return li.root.prev
}

// FindByValue searches by value the first found element in
// forwards direction. If not found the nil is returned.
func (li *DoublyLinkedList[T]) FindByValue(cmp T) *Node[T] {
	for node := li.Front(); node != nil; node = node.Next() {
		if node.Value == cmp {
			return node
		}
	}
	return nil
}

// Remove deletes node from the list.
func (li *DoublyLinkedList[T]) Remove(node *Node[T]) *Node[T] {
	if node == nil {
		return nil
	}
	node.prev.next = node.next
	node.next.prev = node.prev
	node.next = nil
	node.prev = nil
	node.list = nil
	li.length--
	return node
}

// RemoveFront remove the node from the front of the list.
func (li *DoublyLinkedList[T]) RemoveFront() *Node[T] {
	node := li.root.next
	if node == &li.root {
		return nil
	}
	return li.Remove(node)

}

// RemoveBack remove the node from the front of the list.
func (li *DoublyLinkedList[T]) RemoveBack() *Node[T] {
	node := li.root.prev
	if node == &li.root {
		return nil
	}
	return li.Remove(node)
}

// Len returns the length of the list.
func (li *DoublyLinkedList[T]) Len() int {
	return li.length
}

// Empty checks that list is empty.
func (li *DoublyLinkedList[T]) Empty() bool {
	return li.length == 0
}
