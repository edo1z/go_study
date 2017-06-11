package main

import (
	"fmt"
)

func main() {
	fmt.Println(add(12, 15))
}

func add(x, y int) int {
	return x + y
}
