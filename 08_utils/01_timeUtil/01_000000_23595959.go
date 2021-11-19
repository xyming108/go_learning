package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println()
	t := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(t)
	fmt.Println()

	t1 := time.Now().Unix()
	fmt.Println(t1)
	fmt.Println()

	day := time.Now().Format("2006-01-02")
	fmt.Println(day)
	fmt.Println()

	start := day + " 00:00:00"
	startTime, _ := time.ParseInLocation("2006-01-02 15:04:05", start, time.Local)
	fmt.Println(startTime)
	fmt.Println(startTime.Unix())
	fmt.Println()

	// 一天的时间戳相差86400（秒为单位），如果想算其他日期的，直接 天数*86400（秒为单位）（如果要算毫秒，后面加再加 000）

	end := day + " 23:59:59"
	endTime, _ := time.ParseInLocation("2006-01-02 15:04:05", end, time.Local)
	fmt.Println(endTime)
	fmt.Println(endTime.Unix())
}
