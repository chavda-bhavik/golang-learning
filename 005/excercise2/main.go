package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	// values are not assigned to the variables
	// so that they will have zero values
	// zero values are as following
	// 0 for int, 0.0 for float
	// false for bool, '' for string
	// nil/null for pointer, function, map, etc

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	s := fmt.Sprintf("%v\t%v\t%v\n", x, y, z)
	fmt.Println(s)
}
