package main

import "fmt"

/*
给定一个数组，它的第 i 个元素是某只股票第 i 天的价格。你可以多次买卖该股票，但最多只能持有1股，也不允许卖空。
请写一个函数来计算你所能获取的最大利润。

例如，给你数组1,4,2,3}，那么你所能获取的最大利润为4。
*/

func maxProfit(s []int) int {
	if len(s) == 0 || len(s) == 1 {
		return 0
	}
	if len(s) == 2 {
		if s[0] < s[1] {
			return s[1] - s[0]
		} else {
			return 0
		}
	}
	curProfit := 0
	var j int
	for i := 0; i < len(s)-1; i++ {
		j = i + 1
		for s[j-1] < s[j] && j <= len(s)-1 {
			j++
		}
		curProfit += s[j-1] - s[i]
		i = j
	}

	if s[len(s)-2] < s[len(s)-1] {
		curProfit += s[len(s)-1] - s[len(s)-2]
	}

	return curProfit
}

func main() {
	s := []int{1, 4, 2, 3, 5}
	fmt.Println(maxProfit(s))
}
