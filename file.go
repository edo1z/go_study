package main

import (
	"log"
	"os"
)

func getFileName() string {
	if len(os.Args) < 2 {
		log.Fatal("Please input file name")
	}
	return os.Args[1]
}
