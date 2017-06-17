package main

import (
	"bytes"
	"fmt"
	"github.com/wcharczuk/go-chart"
	"io"
	"os"
)

func main() {
	var x []float64
	var y []float64
	for i := 0.0; i < 100; i++ {
		x = append(x, i-50.0)
		y = append(y, (i-50.0)*(i-50.0))
	}
	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: x,
				YValues: y,
			},
		},
	}
	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.PNG, buffer)
	if err != nil {
		fmt.Println("render error")
	}
	out, err := os.Create("./data/graph.png")
	if err != nil {
		fmt.Println("create file error")
	}
	defer out.Close()
	io.Copy(out, buffer)
}
