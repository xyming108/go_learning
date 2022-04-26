package main

import (
	"fmt"
	"github.com/golang-module/carbon"
)

func main() {
	a := carbon.Parse("2022-04-26").SetWeekStartsAt(carbon.Monday).StartOfWeek().ToDateTimeString()
	b := carbon.Parse("2022-04-26").SetWeekStartsAt(carbon.Monday).EndOfWeek().ToDateTimeString()
	fmt.Println(a[0:10])
	fmt.Println(b[0:10])

	fmt.Println(a)
	fmt.Println(b)
}
