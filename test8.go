package main

import (
	"fmt"
	"github.com/golang-module/carbon"
)

func getMonthXAxis(s, e string) ([]string, []string) {
	resultMonthAndUnit, resultMonth := make([]string, 0), make([]string, 0)
	startMonthTime, endMonthTime := fmt.Sprintf(s+"-01 13:14:15"), fmt.Sprintf(e+"-01 13:14:15")
	resultMonth = append(resultMonth, startMonthTime[0:7])
	resultMonthAndUnit = append(resultMonthAndUnit, startMonthTime[0:7]+"月")
	for {
		nextMonthTime := carbon.Parse(startMonthTime).AddMonth().ToDateTimeString()
		resultMonth = append(resultMonth, nextMonthTime[0:7])
		resultMonthAndUnit = append(resultMonthAndUnit, nextMonthTime[0:7]+"月")
		startMonthTime = nextMonthTime
		if nextMonthTime[0:7] == endMonthTime[0:7] {
			break
		}
	}
	return resultMonth, resultMonthAndUnit
}

func GetXAxisPreMonth(yearMonth string, monthNum int) ([]string, []string) {
	resultMonthAndUnit, resultMonth := make([]string, 0), make([]string, 0)
	yearMonthTime := fmt.Sprintf(yearMonth + "-01 13:14:15")
	resultMonth = append(resultMonth, yearMonthTime[0:7])
	resultMonthAndUnit = append(resultMonthAndUnit, yearMonthTime[0:7]+"月")
	for i := 0; i < monthNum; i++ {
		nextMonthTime := carbon.Parse(yearMonthTime).SubMonth().ToDateTimeString()
		resultMonth = append(resultMonth, nextMonthTime[0:7])
		resultMonthAndUnit = append(resultMonthAndUnit, nextMonthTime[0:7]+"月")
		yearMonthTime = nextMonthTime
	}
	return resultMonthAndUnit, resultMonth
}

func main() {
	resultMonthAndUnit, resultMonth := getMonthXAxis("2022-09", "2023-08")
	fmt.Println(resultMonthAndUnit)
	fmt.Println(resultMonth)
	resultMonthAndUnit, resultMonth = GetXAxisPreMonth("2022-09", 5)
	fmt.Println(resultMonthAndUnit)
	fmt.Println(resultMonth)
}
