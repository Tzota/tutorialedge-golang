package main

import "fmt"

func main() {
	type Person struct {
		name string
		age  int
	}

	// our Team struct
	type Team struct {
		name    string
		players [3]Person
	}

	// declaring an empty array of strings
	// var days []string

	// declaring an array with elements
	days := [...]string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}

	fmt.Println(days[0]) // prints 'monday'
	fmt.Println(days[5]) // prints 'saturday'

	weekdays := days[0:5]
	fmt.Println(weekdays)

	youtubeSubscribers := map[string]int{
		"TutorialEdge":     2240,
		"MKBHD":            6580350,
		"Fun Fun Function": 171220,
	}

	fmt.Println(youtubeSubscribers["MKBHD"]) // prints out 6580350

	elliot := Person{name: "Elliot", age: 24}
	fmt.Println(elliot.age)

	// declaring an empty 'Team'
	var myTeam Team
	fmt.Println(myTeam)

	players := [...]Person{Person{name: "Forrest"}, Person{name: "Gordon"}, Person{"tzota", 15}}
	// declaring a team with players
	celtic := Team{name: "Celtic FC", players: players}
	fmt.Println(celtic)
}
