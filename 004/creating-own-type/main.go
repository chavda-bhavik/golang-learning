package main

import (
	"fmt"
)

var a string

type bhavik string

var z bhavik

func main() {
	a = "Temp text"
	z = "Hello Bhavik"
	fmt.Println(z)
	fmt.Printf("%T\n", z)

	// conversions
	a = string(z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
