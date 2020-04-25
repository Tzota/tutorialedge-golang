package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func byteToString(b []byte) string {
	return string(b[:])
}

type myByteArray []byte

func (bytes myByteArray) toString() string {
	return string(bytes[:])
}

func execute() {

	// here we perform the pwd command.
	// we can store the output of this in our out variable
	// and catch any errors in err
	out, err := exec.Command("ls", "-ltr").Output()

	// if there is an error with our execution
	// handle it here
	if err != nil {
		fmt.Printf("%s", err)
	}
	// as the out variable defined above is of type []byte we need to convert
	// this to a string or else we will see garbage printed out in our console
	// this is how we convert it to a string
	fmt.Println("Command Successfully Executed")
	fmt.Println(myByteArray(out).toString())

	// let's try the pwd command herer
	out, err = exec.Command("pwd").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println("Command Successfully Executed")
	fmt.Println(byteToString(out))
}

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		execute()
	}
}
