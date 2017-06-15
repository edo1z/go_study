package main

import (
	"fmt"
)

func main() {
	a := 1
	b := "hoge"
	c := []int{1, 2, 3}
	d := b

	fmt.Println(&a)
	fmt.Printf("%p\n", &a)
	fmt.Println(*&a)
	fmt.Println()

	fmt.Println(b)
	fmt.Println(&b)
	fmt.Println(&d)
	fmt.Println()

	fmt.Println(c)

	hoge(&a)
	hoge(&a)
	hoge(&c[2])
}

func hoge(i *int) {
	fmt.Println("\nhoge start")
	fmt.Printf("%T\n", i)
	fmt.Println(*i)
	fmt.Println(i)
	fmt.Println(&i)
	*i++
	fmt.Println("hoge end\n")
}
