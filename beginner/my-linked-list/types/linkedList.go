package types

import (
	"fmt"
)

// Node is a node in linked list
type Node struct {
	Value string
	next  *Node
}

// Stringer
func (node Node) String() string {
	if node.next == nil {
		return fmt.Sprintf("%v -×", node.Value)
	}
	return fmt.Sprintf("%v -> %v", node.Value, node.next)
}

// LinkedList это черт его знает что
type LinkedList struct {
	Head *Node
	tail *Node
}

// Length - это длина списка
func (ll *LinkedList) Length() int {
	if ll.Head == nil {
		return 0
	}

	length := 1

	for current := ll.Head; current.next != nil; current = current.next {
		length++
	}

	return length
}

// PushFront добавляет узел в голову списка
func (ll *LinkedList) PushFront(node *Node) {
	if ll.Head == nil {
		ll.Head = node
		ll.tail = node
	} else {
		old := ll.Head
		ll.Head = node
		ll.Head.next = old
	}
}

// PushBack добавляет узел в конец списка
func (ll *LinkedList) PushBack(node *Node) {
	if ll.tail == nil {
		ll.Head = node
		ll.tail = node
	} else {
		old := ll.tail
		old.next = node
		ll.tail = node
	}
}

// New это .ctor
func New() LinkedList {
	return LinkedList{Head: nil}
}
