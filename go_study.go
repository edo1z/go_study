package main

import (
	"log"
	"os"
)

func main() {
	err := os.MkdirAll("data/hoge", 0755)
	if err != nil {
		log.Fatal("mkdir error")
	}

	defer os.RemoveAll("data/hoge")

	f, err := os.Create("data/hoge/hoge.txt")
	if err != nil {
		log.Fatal("create file error")
	}

	defer f.Close()
}
