package main

import (
	"fmt"
)

func main() {
	s := "Hello World\n"
	rows := `Hello this " \ \s, yeah we can write anything insise backtick`
	fmt.Println(s)
	fmt.Println(rows)

	bs := []byte(s)
	fmt.Printf("%T\n", bs)
	fmt.Println(bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U\t", s[i])
	}
	fmt.Println("")
	for i, v := range s {
		fmt.Printf("at index %d we have %#x\n", i, v)
	}

	fmt.Printf("%#x\n", "bhavik")
	fmt.Printf("%#U", 'a')

	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	fmt.Println(sample)
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}
	fmt.Printf("%x\n", sample)
}
