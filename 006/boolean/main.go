package main

import "fmt"

var x bool

func main() {
	// printing zero value assign to the boolean which is false
	fmt.Println(x)
	x = true
	fmt.Println(x) // Printing changed value

	// using comparision operators that produces boolean result
	a := 3
	b := 5
	fmt.Println(a == b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)
	fmt.Println(a > b)
	fmt.Println(a < b)
}
