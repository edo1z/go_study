package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		123.123, -234.234,
	},
	"Google": Vertex{
		555.555, -122.122,
	},
}

func main() {
	for i, v := range m {
		fmt.Printf("%s => %f\n", i, v)
	}
}
