package main

import (
	"github.com/sndnvaps/go-yql-finance"
	"fmt"
)


func main() {

/*
请求地址
http://ichart.yahoo.com/table.csv?s=<string>&a=<int>&b=<int>&c=<int>&d=<int>&e=<int>&f=<int>&g=d&ignore=.csv
参数
s – 股票名称
a – 起始时间，月
b – 起始时间，日
c – 起始时间，年
d – 结束时间，月
e – 结束时间，日
f – 结束时间，年
g – 时间周期。Example: g=w, 表示周期是’周’。d->’日’(day), w->’周’(week)，m->’月’(mouth)，v->’dividends only’
一定注意月份参数，其值比真实数据-1。如需要9月数据，则写为08。
*/
	// The second parameter is the time interval.
	// You can pass in "d", "w", "m", "daily", "weekly", or "montly".
//select * from csv where url='http://ichart.finance.yahoo.com/table.csv?s=aapl&g=d' and columns='Date,Open,High,Low,Close,Volume,AdjClose'
//select * from csv where url='http://ichart.finance.yahoo.com/table.csv?s=aapl&g=m' and columns='Date,Open,High,Low,Close,Volume,AdjClose'

// 2016-09-07 ~ 2016-10-07
//select * from csv where url='http://ichart.finance.yahoo.com/table.csv?s=aapl&g=m&a=08&b=07&c=2016&d=09&e=07&f=2016&ignore=.csv and columns='Date,Open,High,Low,Close,Volume,AdjClose'

/*
	the_time, err := time.Parse("2006-01-02", "2014-01-08")
	if err == nil {
        unix_time := the_time.Unix() 
        fmt.Println(unix_time)
        fmt.Println(the_time.Year())
        fmt.Printf("%d\n",the_time.Month())
        fmt.Println(the_time.Day())
*/
	var date_duration [2]string = [2]string{"2016-01-01", "2016-02-01"}
	days := yql.GetHistoricalData("GOOG", "daily", date_duration)

	for _, day := range days {
		fmt.Println(day.Date)
		fmt.Println(day.Close)
	}
}
