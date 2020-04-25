package main

import (
	"fmt"
)

func myFunc(a interface{}) {
	fmt.Println(a)
}

// #region

type Guitarist interface {
	// PlayGuitar prints out "Playing Guitar"
	// to the terminal
	PlayGuitar()
}

type BaseGuitarist struct {
	Name string
}

type AcousticGuitarist struct {
	Name string
}

func (b BaseGuitarist) PlayGuitar() {
	fmt.Printf("%s plays the Bass Guitar\n", b.Name)
}

func (b AcousticGuitarist) PlayGuitar() {
	fmt.Printf("%s plays the Acoustic Guitar\n", b.Name)
}

// #endregion

// #region

type Employee interface {
	// Name() string
	Language() string
	// Age() int
	// Random() (string, error)
}

type Engineer struct {
	Name string
}

// func (e *Engineer) Name() string {
// 	return e.Name
// }

func (e *Engineer) Language() string {
	return e.Name + " programs in Go"
}

// func (e *Engineer) Age() int {
// 	return 15
// }

// func (e *Engineer) Random() (string, error) {
// 	return "not a random", nil
// }

// #endregion

func main() {
	// #region
	var my_age int
	my_age = 25

	myFunc(my_age)
	myFunc("test")
	// #endregion

	// #region
	var player BaseGuitarist = BaseGuitarist{Name: "Paul"}
	player.PlayGuitar()

	var player2 AcousticGuitarist
	player2.Name = "Ringo"
	player2.PlayGuitar()

	var guitarists []Guitarist
	guitarists = append(guitarists, player)
	guitarists = append(guitarists, player2)

	for _, val := range guitarists {
		val.PlayGuitar()
	}
	// #endregion

	// #region

	// This will throw an error
	var programmers []Employee
	elliot := Engineer{Name: "Elliot"}
	// Engineer does not implement the Employee interface
	// you'll need to implement Age() and Random()
	programmers = append(programmers, &elliot)

	for index, val := range programmers {
		fmt.Println(index, val.Language())
	}

	// #endregion
}
