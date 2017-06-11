package main

import (
	"fmt"
)

var c, python, php bool

func hoge() {
	c = true
}

func main() {
	var i int
	hoge()
	fmt.Println(i, c, python, php)
}
