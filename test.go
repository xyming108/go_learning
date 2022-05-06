package main

import (
	"fmt"
	"github.com/golang-module/carbon"
	"time"
)

func getXAxis(startTime, endTime int64) ([]string, []string) {
	resultWeekAndDay := make([]string, 0, 10)
	resultWeek := make([]string, 0, 10)
	filterMap := make(map[string]int)
	s := time.Unix(startTime, 0)
	e := time.Unix(endTime, 0)
	//After方法 a.After(b) a,b Time类型 如果a时间在b时间之后，则返回true
	for e.After(s) {
		weekNum, a, b := getStartAndEnd(s)
		r := fmt.Sprintf("%d周(%d/%d-%d/%d)", weekNum, a.Month(), a.Day(), b.Month(), b.Day())
		if _, ok := filterMap[r]; !ok {
			resultWeekAndDay = append(resultWeekAndDay, r)
			resultWeek = append(resultWeek, r[0:5])
			filterMap[r] = 1
		}
		s = s.AddDate(0, 0, 7)
	}
	weekNum, a, b := getStartAndEnd(s)
	r := fmt.Sprintf("%d周(%d/%d-%d/%d)", weekNum, a.Month(), a.Day(), b.Month(), b.Day())
	if _, ok := filterMap[r]; !ok {
		resultWeekAndDay = append(resultWeekAndDay, r)
		resultWeek = append(resultWeek, r[0:5])
	}
	return resultWeekAndDay, resultWeek
}

func getStartAndEnd(s time.Time) (int, time.Time, time.Time) {
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

func main() {
	startTme, endTime := int64(1651818927), int64(1654497327)
	//时间戳转换成日期字符串，如：2022-05-06 14:35:27 +0800 CST
	start := carbon.CreateFromTimestamp(startTme).ToString()
	end := carbon.CreateFromTimestamp(endTime).ToString()
	fmt.Println(start)
	fmt.Println(end)
	//获取周的星期一和星期天，得到的结果格式如：2022-05-02 00:00:00
	startDay := carbon.Parse(start[0:10]).SetWeekStartsAt(carbon.Monday).StartOfWeek().ToDateTimeString()
	endDay := carbon.Parse(end[0:10]).SetWeekStartsAt(carbon.Monday).EndOfWeek().ToDateTimeString()
	fmt.Println("------------startDay: ", startDay)
	fmt.Println("------------endDay: ", endDay)
	//把结果转换成时间戳
	stampStart := carbon.Parse(startDay).Timestamp() //startTme的周一时间戳
	stampEnd := carbon.Parse(endDay).Timestamp()     //endTime的周天时间戳
	fmt.Println("1: ", stampStart)
	fmt.Println("2: ", stampEnd)

	sInt := carbon.Parse(start[0:10]).WeekOfYear()
	fmt.Println(sInt)

	aa, bb := getXAxis(startTme, endTime)
	fmt.Println(aa, "\n", bb)
}
