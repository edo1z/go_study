package main

import (
	_ "bufio"
	_ "fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
}
