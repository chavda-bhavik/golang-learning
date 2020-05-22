package main

import (
	"fmt"
)

var y = 42
var a int

type bhavik string

var z bhavik

func main() {
	a = 42
	z = "Hello Bhavik"
	fmt.Println(z)
	fmt.Printf("%T\n", z)
}
