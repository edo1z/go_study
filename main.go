package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	go signalTest()
	fmt.Println("start")
	for {
		continue
	}
}

func signalTest() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	s := <-c
	fmt.Println()
	fmt.Println("got signal: ", s)
	os.Exit(1)
}
