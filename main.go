package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("start")
	var wg sync.WaitGroup
	for _, a := range []string{"1", "2"} {
		wg.Add(1)
		go func(str string) {
			fmt.Println("I'm " + str)
			time.Sleep(1 * time.Second)
			wg.Done()
		}(a)
	}
	wg.Wait()
	fmt.Println("end")
}
