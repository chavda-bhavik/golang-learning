package main

import (
	"fmt"
)

const (
	_ = iota
	b = 1 << iota
	c = 1 << iota
	d
	e
	f
)

func main() {
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
