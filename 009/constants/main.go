package main

import "fmt"

// Declaring single constant
const a = 32

// Declaring Untyped bunch of constant
const (
	x = 43
	y = 99
)

// Declaring typed bunch of constant
const (
	m int     = 44
	n float64 = 33
)

func main() {
	fmt.Printf("%T\t%v\n", a, a)
	fmt.Printf("%T\t%v\n", x, x)
	fmt.Printf("%T\t%v\n", y, y)
	fmt.Printf("%T\t%v\n", m, m)
	fmt.Printf("%T\t%v\n", n, n)
}
