package main

import (
	"fmt"
)

func main() {
	a := "string..."
	b := 29
	c := 0.54
	d := uint32(b)

	fmt.Printf("a is of type %T\n", a)
	fmt.Printf("b is of type %T\n", b)
	fmt.Printf("c is of type %T\n", c)
	fmt.Printf("d is of type %T\n", d)
}
