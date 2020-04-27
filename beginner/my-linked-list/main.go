package main

import (
	"fmt"
)

// Node is a node in linked list
type Node struct {
	value string
	next  *Node
}

// Stringer
func (node Node) String() string {
	if node.next == nil {
		return fmt.Sprintf("%v -×", node.value)
	}
	return fmt.Sprintf("%v -> %v", node.value, node.next)
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (ll *LinkedList) Length() int {
	if ll.head == nil {
		return 0
	}

	length := 1

	for current := ll.head; current.next != nil; current = current.next {
		length++
	}

	return length
}

func (ll *LinkedList) PushFront(node *Node) {
	if ll.head == nil {
		ll.head = node
		ll.tail = node
	} else {
		old := ll.head
		ll.head = node
		ll.head.next = old
	}
}

func (ll *LinkedList) PushBack(node *Node) {
	if ll.tail == nil {
		ll.head = node
		ll.tail = node
	} else {
		old := ll.tail
		old.next = node
		ll.tail = node
	}
}

func New() LinkedList {
	return LinkedList{head: nil}
}

func main() {
	ll := New()
	ll.PushFront(&Node{value: "2"})
	ll.PushFront(&Node{value: "1"})
	ll.PushBack(&Node{value: "3"})
	fmt.Printf("%v, total length: %v\n", ll.head, ll.Length())
}
