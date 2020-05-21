package main

import "fmt"

// Declare variable with keyword "y" and assign it value 20
var y = 20

// if you didn't assign value to the variable it'll take zero value
// zero values are as below for TYPES
// 0 for int, 0.0 for float
// '' for string, false for boolean,
// nil for pointers, functions, interfaces, slices, channels and maps

// Declare variable with keyword "z" of type String and assign zero value which is "" empty
var z string

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T\n", z)
}
