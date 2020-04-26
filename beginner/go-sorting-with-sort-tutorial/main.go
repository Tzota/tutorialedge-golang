package main

import (
	"fmt"
	"sort"
)

type Programmer struct {
	Age int
}

type byAge []Programmer

func (p byAge) Len() int {
	return len(p)
}

func (p byAge) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p byAge) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}

func main() {
	fmt.Println("Go Sorting Tutorial")

	// our unsorted int array
	myInts := []int{1, 3, 2, 6, 3, 4}
	fmt.Println(myInts)

	// we can use the sort.Ints
	sort.Ints(myInts)
	fmt.Println(myInts)

	programmers := []Programmer{
		Programmer{Age: 30},
		Programmer{Age: 20},
		Programmer{Age: 50},
		Programmer{Age: 1000},
	}

	sort.Sort(byAge(programmers))

	fmt.Println(programmers)
}
