package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	num := rand.Intn(10)
	pi := math.Pi
	fmt.Printf("Number is %d\n", num)
	fmt.Printf("PI is %f\n", pi)
}
