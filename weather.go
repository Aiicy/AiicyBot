package main

import (
	yahoo "github.com/sndnvaps/yahoo_weather_api"
	"fmt"
)

//location == "jiangmen,guangdong,china"
func GetWeatherForecastInfo(location string) (string, string) {
	f, icon_url  := yahoo.GetForecastlData(location)
	outstr := fmt.Sprintf("Weather Forecast of %s\n", location)
	outstr = outstr + fmt.Sprintf(" Day Date ,Low(℃),High(℃), Text\n")
	for _ , v  := range f {
		outstr = outstr + fmt.Sprintf("%s %s, %.2f, %.2f ,%s\n", v.Day, v.Date, v.Low, v.High, v.Text)
	}
	return outstr, icon_url
}

