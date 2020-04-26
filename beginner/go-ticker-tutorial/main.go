package main

import (
	"fmt"
	"time"
)

func backgroundTask() {
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		fmt.Println("Tock")
	}
}

func main() {
	fmt.Println("Go Tickers Tutorial")
	// this creates a new ticker which will
	// `tick` every 1 second.
	ticker := time.NewTicker(1 * time.Second)

	go backgroundTask()

	// for every `tick` that our `ticker`
	// emits, we print `tock`
	for c := range ticker.C {
		fmt.Printf("tock %v\n", c)
	}
}
