package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

func newFile(fn string) *os.File {
	fp, err := os.OpenFile(fn, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return fp
}

func main() {
	fp := newFile("./data/data.txt")
	defer fp.Close()
	reader := bufio.NewReader(fp)

	str, err := Readln(reader)
	for err == nil {
		fmt.Println(str)
		str, err = Readln(reader)
	}
}
