package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	// #region integer overflow
	var myint int8
	for i := 0; i < 129; i++ {
		myint += 1
	}
	fmt.Println(myint) // prints out -127
	// #endregion

	// #region Conversion of Types
	var men uint8 = 5
	var women int16 = 6

	var people int
	// this throws a compile error
	// people = men + women
	// this handles converting to a standard format
	// and is legal within your go programs
	people = int(men) + int(women)
	fmt.Println(people)
	// #endregion

	// #region Floating Point Numbers
	var maxFloat32 float32 = 16777216
	fmt.Println(maxFloat32 == maxFloat32+10) // you would typically expect this to return false
	// it returns true
	fmt.Println(maxFloat32 + 10)      // 16777216
	fmt.Println(maxFloat32 + 2000000) // 16777216

	// converting from int to float
	var myint3 int
	myfloat := float64(myint3)

	// converting from float to int
	var myfloat2 float64
	myint2 := int(myfloat2)
	fmt.Println(myfloat, myint2)
	// #endregion
}
