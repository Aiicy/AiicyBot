package main

import (
	"fmt"
	"log"

	"github.com/sndnvaps/go-yql-finance"
	"github.com/wcharczuk/go-chart"
	"image/png"
	"os"
	"time"
)

func main() {

	var date_duration [2]string = [2]string{"2016-03-01", "2016-03-30"}
	days := yql.GetHistoricalData("AAPL", "daily", date_duration)
	var day_t_x []time.Time
	var day_c_y []float64
	var day_c_y1 []float64

	for _, day := range days {
		day_t_x = append(day_t_x, day.Date)
		day_c_y = append(day_c_y, day.Low)
		day_c_y1 = append(day_c_y1, day.High)
	}

	// for debug
	//fmt.Println(day_c_y)
	//fmt.Println("\n")
	//fmt.Println(day_c_y1)

	graph := chart.Chart{
		Width:  1920,
		Height: 1080,
		XAxis: chart.XAxis{
			Name:      "Stock info of AAPL",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
		YAxis: chart.YAxis{
			Name:      "Low price",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
		YAxisSecondary: chart.YAxis{
			Name:      "High price",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(), //enables / displays the secondary y-axis
		},
		Series: []chart.Series{
			chart.TimeSeries{
				Name:    "The Low price of AAPL at March",
				XValues: day_t_x, // []time.Time{ ...},
				YValues: day_c_y, //[]float64{...},
			},
			chart.TimeSeries{
				Name:    "The High price of AAPL at March",
				YAxis:   chart.YAxisSecondary,
				XValues: day_t_x,  // []time.Time{ ...},
				YValues: day_c_y1, //[]float64{...},
			},
		},
	}

	//note we have to do this as a separate step because we need a reference to graph
	graph.Elements = []chart.Renderable{
		chart.Legend(&graph),
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

	err = png.Encode(file1, Image)
	if err != nil {
		log.Fatal(err)
	}
}
