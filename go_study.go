package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")

	file_path := GetFileName()
	content := ReadAllOfFile(file_path)
	fmt.Println(content)

	content = ReadAllOfFile2(file_path)
	fmt.Println(content)

	ReadFile(file_path)
}
