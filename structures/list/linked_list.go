package list

// Node represents the node in linked in.
// TODO comparable may not be equal of using structs.
// NOTE Maybe it is better to use reflect.DeepEqual() to be sure?
type Node[T comparable] struct {
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
//
// NOTE This list is created only for educational purposes and not recommenced
// to use in production. If you need the linked lists structures it is better
// to use https://pkg.go.dev/container/list
type LinkedList[T comparable] struct {
	head, tail *Node[T]
	length     int
}

// NewLinkedList create a new linked list.
func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// Append adds the value to the end of the list.
func (li *LinkedList[T]) Append(v T) *Node[T] {
	return li.InsertWithValue(v, li.tail)
}

// InsertWithValue creates a node with given value and inserts it after the
// given node.
func (li *LinkedList[T]) InsertWithValue(v T, after *Node[T]) *Node[T] {
	node := &Node[T]{value: v}
	return li.InsertAfter(node, after)
}

// InsertAfter a node after a given node.
func (li *LinkedList[T]) InsertAfter(node, after *Node[T]) *Node[T] {
	li.length++
	if after == nil {
		node.next = li.head
		li.head = node
		if li.tail == nil {
			li.tail = node
		}
		return node
	}
	if after != li.tail {
		node.next = after.next
	} else {
		li.tail = node
	}
	after.next = node
	return node
}

// RemoveByValue the first found node by value fron linked list.
func (li *LinkedList[T]) RemoveByValue(v T) *Node[T] {
	node := li.FindByValue(v)
	if node == nil {
		return nil
	}
	return li.Remove(node)
}

// Remove removes the node from the linked list.
func (li *LinkedList[T]) Remove(n *Node[T]) *Node[T] {
	var prev *Node[T]
	curr := li.head
	if curr != nil && curr == n {
		li.head = curr.next
		li.length--
		return curr
	}

	for curr != nil && n != curr {
		prev = curr
		curr = curr.next
	}

	if curr == nil {
		return nil
	}
	prev.next = curr.next
	li.length--
	// Need to figure about "island of pointer" in Go.
	// Prevent memory leak?
	// curr.next = nil
	return curr
}

// FindByValue search the node by it's value.
func (li *LinkedList[T]) FindByValue(v T) *Node[T] {
	for node := li.head; node != nil; node = node.Next() {
		if node.Value() == v {
			return node
		}
	}
	return nil
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

// Len returns the length of the linked list.
func (li *LinkedList[T]) Len() int {
	return li.length
}
