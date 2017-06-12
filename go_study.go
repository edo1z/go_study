package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	count := make(map[string]int)
	words := strings.Split(s, " ")
	for _, w := range words {
		count[w] += 1
	}
	return count
}

func main() {
	wc.Test(WordCount)
}
