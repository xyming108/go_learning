package main

import (
	"fmt"
	"github.com/golang-module/carbon"
	sort2 "sort"
	"time"
)

func main() {
	xAxis := make([]string, 0, 7)
	xAxis = append(xAxis, "2022-09-21")
	stampStart := carbon.Parse("2022-09-21").Timestamp()
	fmt.Println(stampStart)
	s := time.Unix(stampStart, 0)
	for i := 0; i < 6; i++ {
		duration, err := time.ParseDuration("-24h")
		if err != nil {
			return
		}
		tmpTime := s.Add(duration)
		xAxis = append(xAxis, tmpTime.String()[0:10])
		s = tmpTime
	}
	sort2.Strings(xAxis)
	fmt.Println(xAxis)
}