package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	primes[3] = 9
	fmt.Println(primes)

	fmt.Println(len(a))
	fmt.Println(len(primes))
}
