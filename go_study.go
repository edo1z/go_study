package main

import (
	"fmt"
)

func main() {
	defer fmt.Println(". Thank you.")
	defer fmt.Print(" world")
	fmt.Print("Hello")
}
