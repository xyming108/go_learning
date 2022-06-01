package main

import (
	"fmt"
	"github.com/golang-module/carbon"
	"time"
)

func getStartAnd(s time.Time) (int, time.Time, time.Time) {
	//获取当前时间处于这一年的第几周
	weekNum := carbon.Parse(s.String()[0:10]).WeekOfYear()
	//获取周的星期一和星期天，得到的结果格式如：2022-05-02 00:00:00
	startDay := carbon.Parse(s.String()[0:10]).SetWeekStartsAt(carbon.Monday).StartOfWeek().ToDateTimeString()
	endDay := carbon.Parse(s.String()[0:10]).SetWeekStartsAt(carbon.Monday).EndOfWeek().ToDateTimeString()
	//把结果转换成时间戳
	stampStart := carbon.Parse(startDay).Timestamp() //startTime的周一时间戳
	stampEnd := carbon.Parse(endDay).Timestamp()     //endTime的周天时间戳
	a := time.Unix(stampStart, 0)
	b := time.Unix(stampEnd, 0)
	return weekNum, a, b
}

func getWeekAxis(startTime, endTime int64) ([]string, []string) {
	resultWeekAndDay := make([]string, 0, 10)
	resultWeek := make([]string, 0, 10)
	filterMap := make(map[string]int)
	s := time.Unix(startTime, 0)
	e := time.Unix(endTime, 0)
	//After方法 a.After(b) a,b Time类型 如果a时间在b时间之后，则返回true
	for e.After(s) {
		weekNum, a, b := getStartAnd(s)
		//%02d 表示整数固定位数为2，不够前面用0填充
		r := fmt.Sprintf("%d-%02d周(%d/%d-%d/%d)", a.Year(), weekNum, a.Month(), a.Day(), b.Month(), b.Day())
		if _, ok := filterMap[r]; !ok {
			resultWeekAndDay = append(resultWeekAndDay, r)
			resultWeek = append(resultWeek, r[0:10])
			filterMap[r] = 1
		}
		s = s.AddDate(0, 0, 7)
	}
	return resultWeekAndDay, resultWeek
}

func main() {
	xAxisWeekAndDay, xAxisWeek := getWeekAxis(1622095215, 1653631215)
	fmt.Println(xAxisWeekAndDay)
	fmt.Println()
	fmt.Println(xAxisWeek)
}
