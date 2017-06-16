package main

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
)

func task(a, b int) {
	fmt.Println("I am running task.")
	fmt.Printf("%d + %d = %d", a, b, a+b)
}

func main() {
	gocron.Clear()
	s := gocron.NewScheduler()
	s.Every(1).Minute().Do(task, 10, 5)
	s.Start()
}
