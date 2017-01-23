package main

import (
	"fmt"
	"log"

	"github.com/wcharczuk/go-chart"
	"github.com/sndnvaps/go-yql-finance"
	"image/png"
	"os"
	"time"
)

func main() {
	
	var date_duration [2]string = [2]string{"2016-01-01", "2016-06-01"}
	days := yql.GetHistoricalData("GOOG", "daily", date_duration)
	var day_t_x []time.Time
	var day_c_y []float64
	for _, day := range days {
		day_t_x = append(day_t_x, day.Date)
		day_c_y = append(day_c_y, day.Close)
	}

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Style: chart.Style{
				Show: true,
			},
		},
		Series: []chart.Series{
			chart.TimeSeries{
				XValues: day_t_x ,// []time.Time{ ...},
				YValues:  day_c_y , //[]float64{...},
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

