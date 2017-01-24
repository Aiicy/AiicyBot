package main

import (
	"github.com/sndnvaps/yahoo_weather_api"
	"fmt"
)

/*
select woeid from geo.places where text="jiangmen,guangdong,china"


select * from weather.forecast where woeid in  (select woeid from geo.places where text="jiangmen,guangdong,china")

https://query.yahooapis.com/v1/public/yql?q=select%20*%20from%20weather.forecast%20where%20woeid%20in%20%20(select%20woeid%20from%20geo.places%20where%20text%3D%22jiangmen%2Cguangdong%2Cchina%22)&format=json&diagnostics=true&callback=


{"query":{"count":1,"created":"2017-01-24T15:33:02Z","lang":"zh-CN","diagnostics":{"publiclyCallable":"true","url":[{"execution-start-time":"2","execution-stop-time":"8","execution-time":"6","content":"http://unifiedgeo.geotech.yahoo.com:4080/geo/v1/geocode?q=jiangmen%2Cguangdong%2Cchina&start=0&size=10&optionalfields=woe.ancestors&minconfidence=0.0001"},{"execution-start-time":"11","execution-stop-time":"16","execution-time":"5","content":"http://weather-ydn-yql.media.yahoo.com:4080/v3/public/weather/rss?w=2161847"}],"javascript":{"execution-start-time":"0","execution-stop-time":"10","execution-time":"9","instructions-used":"16110","table-name":"geo.places"},"user-time":"17","service-time":"11","build-version":"2.0.84"},"results":{"channel":{"units":{"distance":"mi","pressure":"in","speed":"mph","temperature":"F"},"title":"Yahoo! Weather - Jiangmen, Guangdong, CN","link":"http://us.rd.yahoo.com/dailynews/rss/weather/Country__Country/*https://weather.yahoo.com/country/state/city-2161847/","description":"Yahoo! Weather for Jiangmen, Guangdong, CN","language":"en-us","lastBuildDate":"Tue, 24 Jan 2017 11:33 PM CST","ttl":"60","location":{"city":"Jiangmen","country":"China","region":" Guangdong"},"wind":{"chill":"59","direction":"115","speed":"7"},"atmosphere":{"humidity":"74","pressure":"1025.0","rising":"0","visibility":"16.1"},"astronomy":{"sunrise":"7:9 am","sunset":"6:11 pm"},"image":{"title":"Yahoo! Weather","width":"142","height":"18","link":"http://weather.yahoo.com","url":"http://l.yimg.com/a/i/brand/purplelogo//uh/us/news-wea.gif"},"item":{"title":"Conditions for Jiangmen, Guangdong, CN at 11:00 PM CST","lat":"22.572001","long":"113.07682","link":"http://us.rd.yahoo.com/dailynews/rss/weather/Country__Country/*https://weather.yahoo.com/country/state/city-2161847/","pubDate":"Tue, 24 Jan 2017 11:00 PM CST","condition":{"code":"29","date":"Tue, 24 Jan 2017 11:00 PM CST","temp":"60","text":"Partly Cloudy"},"forecast":[{"code":"30","date":"24 Jan 2017","day":"Tue","high":"68","low":"55","text":"Partly Cloudy"},{"code":"30","date":"25 Jan 2017","day":"Wed","high":"71","low":"56","text":"Partly Cloudy"},{"code":"32","date":"26 Jan 2017","day":"Thu","high":"71","low":"52","text":"Sunny"},{"code":"34","date":"27 Jan 2017","day":"Fri","high":"71","low":"51","text":"Mostly Sunny"},{"code":"30","date":"28 Jan 2017","day":"Sat","high":"72","low":"55","text":"Partly Cloudy"},{"code":"28","date":"29 Jan 2017","day":"Sun","high":"72","low":"62","text":"Mostly Cloudy"},{"code":"28","date":"30 Jan 2017","day":"Mon","high":"73","low":"64","text":"Mostly Cloudy"},{"code":"28","date":"31 Jan 2017","day":"Tue","high":"71","low":"59","text":"Mostly Cloudy"},{"code":"30","date":"01 Feb 2017","day":"Wed","high":"74","low":"60","text":"Partly Cloudy"},{"code":"30","date":"02 Feb 2017","day":"Thu","high":"73","low":"60","text":"Partly Cloudy"}],"description":"<![CDATA[<img src=\"http://l.yimg.com/a/i/us/we/52/29.gif\"/>\n<BR />\n<b>Current Conditions:</b>\n<BR />Partly Cloudy\n<BR />\n<BR />\n<b>Forecast:</b>\n<BR /> Tue - Partly Cloudy. High: 68Low: 55\n<BR /> Wed - Partly Cloudy. High: 71Low: 56\n<BR /> Thu - Sunny. High: 71Low: 52\n<BR /> Fri - Mostly Sunny. High: 71Low: 51\n<BR /> Sat - Partly Cloudy. High: 72Low: 55\n<BR />\n<BR />\n<a href=\"http://us.rd.yahoo.com/dailynews/rss/weather/Country__Country/*https://weather.yahoo.com/country/state/city-2161847/\">Full Forecast at Yahoo! Weather</a>\n<BR />\n<BR />\n(provided by <a href=\"http://www.weather.com\" >The Weather Channel</a>)\n<BR />\n]]>","guid":{"isPermaLink":"false"}}}}}}

*/

func main() {

	f := yahoo.GetForecastlData("jiangmen,guangdong,china")
	fmt.Println(f)
}

