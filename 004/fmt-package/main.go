package main

import "fmt"

var x int = 10

func main() {
	fmt.Println(x)
	// formating output
	fmt.Printf("Original value %v\n", x)
	fmt.Printf("Type of value %T\n", x)
	fmt.Printf("Binary formated value %b\n", x)
	fmt.Printf("Base 10 value %d\n", x)
	fmt.Printf("Base 8 value %o\n", x)
	fmt.Printf("Base 16, with lower-case letters %x\n", x)
	fmt.Printf("Base 16, with upper-case letters %X\n", x)
	fmt.Printf("Unicode format %U\n", x)
	fmt.Printf("%v\t%T\t%d\n", x, x, x)
}
