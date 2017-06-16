package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	target := "./data/"
	chkTargetDir(target)
}

func chkTargetDir(target string) {
	info, err := os.Stat(target)
	if err != nil {
		fmt.Println("target is invalid.")
	} else if !info.IsDir() {
		fmt.Println("target is not a valid directory")
	}
	println(strings.TrimSuffix(target, "/"))
}
