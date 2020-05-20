package main

import (
	"fmt"
)

func main() {

	// Control flows
	fmt.Println("Just printing anything")
	foo()

	for i := 0; i < 50; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()

	// handling return and error
	n, _ := fmt.Println("testing return and error")
	fmt.Println(n)
}

func foo() {
	fmt.Println("Printing from foo")
}
func bar() {
	fmt.Println("and here we go.")
}
