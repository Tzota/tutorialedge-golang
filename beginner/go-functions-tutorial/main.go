package main

import "fmt"

func myfunction(firstName string, lastName string) (string, error) {
	fullname := firstName + " " + lastName
	return fullname, nil
}

func addOne() func() int {
	var x int
	// we define and return an
	// anonymous function which in turn
	// returns an integer value
	return func() int {
		// this anonymous function
		// has access to the x variable
		// defined in the parent function
		x++
		return x + 1
	}
}

// Add sdf sdf
func Add(x int, y int) int {
	return x + y
}

func main() {
	// we can assign the results to multiple variables
	// by defining their names in a comma separated list
	// like so:
	fullName, err := myfunction("Elliot", "Forbes")
	if err != nil {
		fmt.Println("Handle Error Case")
	}
	fmt.Println(fullName)

	myFunc := addOne()
	fmt.Println(myFunc()) // 2
	fmt.Println(myFunc()) // 3
	fmt.Println(myFunc()) // 4
	fmt.Println(myFunc()) // 5
}
