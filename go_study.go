package main

import (
	"fmt"
)

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[1:2]
	printSlice(s)

	s = s[3:4]
	printSlice(s)

	s = s[1:]
	printSlice(s)

	fmt.Println("?")
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
