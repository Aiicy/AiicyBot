package main

import (
	"fmt"
	"log"

	"github.com/wcharczuk/go-chart"
	"image/png"
	"os"
)

func main() {
	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
				YValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
			},
		},
	}
	collector := &chart.ImageWriter{}
	graph.Render(chart.PNG, collector)

	Image, err := collector.Image()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Final Image: %dx%d\n", Image.Bounds().Size().X, Image.Bounds().Size().Y)
	
	file1, err := os.Create("test.png")
    if err != nil {
        panic(err)
    }
    defer file1.Close()

	err = png.Encode(file1,Image)
	if err != nil {
		log.Fatal(err)
	}
}
