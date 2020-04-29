package main

import (
	"fmt"

	t "example.com/kindauser/kindamodule/types"
)

func main() {
	ll := t.New()
	ll.PushFront(&t.Node{Value: "2"})
	ll.PushFront(&t.Node{Value: "1"})
	ll.PushBack(&t.Node{Value: "3"})
	fmt.Printf("%v, total length: %v\n", ll.Head, ll.Length())
}
