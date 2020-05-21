package main

import "fmt"

// Declare + Assign = Instantiate
// Declare variable with "y" keyword and assign it value 30
var y = 30

func main() {
	// Short declarative operator
	// := is a operator, coz it's operating
	// x & 20 r operands, coz we are working on them
	x := 20

	fmt.Println("You're doing great")
	fmt.Println(x)

	foo()
}

func foo() {
	fmt.Println(y)

	// cann't access x, coz it's out of the scope
	// fmt.Println(x)
}
