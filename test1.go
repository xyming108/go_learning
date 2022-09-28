package main

import (
	"fmt"
	"github.com/golang-module/carbon"
	sort2 "sort"
	"time"
)

func sort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := len(arr) - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}

	return arr
}

func GetWeekXAxis(startWeek string, duration int) (string, string, []string) {
	week := make([]string, 0, 10)
	for i := 0; i < duration; i++ {
		if i == 0 {
			week = append(week, startWeek)
		} else {

		}
	}
	return "", "", nil
}

func GetStartAndEnd(s time.Time) (int, time.Time, time.Time) {
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

func main() {
	/*a := []int{2, 4, 6, 9, 1}
	fmt.Println(sort(a))*/

	/*startWeek, endWeek, xAxis := GetWeekXAxis("2022-40", 10)
	fmt.Println(startWeek)
	fmt.Println(endWeek)
	fmt.Println(xAxis)

	a := "2022-05"
	a = a[5:7]
	b, _ := strconv.Atoi(a)
	fmt.Println(b - 1)*/

	var startTime int64 = 1136185445
	var endTime int64 = 1664259112
	s := time.Unix(startTime, 0)
	e := time.Unix(endTime, 0)
	// 先往后算7天，方便下面循环的计算
	e = e.AddDate(0, 0, 7)
	duration := 10
	week := make([]string, 0, duration)
	for e.After(s) {
		if duration == 0 {
			break
		}
		e = e.AddDate(0, 0, -7)
		// 计算这一天所在周的星期一
		everyMonday := carbon.Parse(e.String()[0:10]).SetWeekStartsAt(carbon.Monday).StartOfWeek().ToDateTimeString()
		// 根据星期一计算属于该年中的哪一周
		weekNum, a, b := GetStartAndEnd(e)
		//weekNum := carbon.Parse(everyMonday).WeekOfYear()
		r := fmt.Sprintf("%s-%02d周(%d/%d-%d/%d)", everyMonday[0:4], weekNum, a.Month(), a.Day(), b.Month(), b.Day())
		week = append(week, r)
		fmt.Println(r)
		duration--
	}
	sort2.Slice(week, func(i, j int) bool {
		return week[i] < week[j]
	})
	fmt.Println(week)
	startWeek, endWeek := (week[0])[0:7], (week[len(week)-1])[0:7]
	fmt.Println(startWeek)
	fmt.Println(endWeek)
}
