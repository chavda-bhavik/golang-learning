package main

import (
	"fmt"
	"runtime"
)

// Difference between uint and int are
// unit stores value from 0 to x
// int stores value from -x to +x

var a int8 = 0
var b uint8 = 0
var c float32 = 55555555555555.0

func main() {
	x := 32
	y := -555
	z := 55.55

	fmt.Printf("%T\t%v\n", x, x)
	fmt.Printf("%T\t%v\n", y, y)
	fmt.Printf("%T\t%v\n", z, z)

	// a = 128 // throws an error
	fmt.Printf("%T\t%v\n", a, a)
	fmt.Printf("%T\t%v\n", b, b)
	fmt.Printf("%T\t%v\n", c, c)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.MemProfile)
	fmt.Println(runtime.Stack)
	fmt.Println(runtime.GOROOT)
}
