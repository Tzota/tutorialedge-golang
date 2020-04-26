package main

import "fmt"

var name string

func init() {
	fmt.Println("This will get called on main initialization")
	name = "Elliot"
}

func main() {
	fmt.Println("My Wonderful Go Program")
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Init Counter: %d\n", initCounter)
}

var initCounter int

func init() {
	fmt.Println("Called First in Order of Declaration")
	initCounter++
}

func init() {
	fmt.Println("Called second in order of declaration")
	initCounter++
}
