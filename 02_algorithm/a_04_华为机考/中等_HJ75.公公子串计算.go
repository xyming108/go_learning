package main

import "fmt"

/*
输入描述：
输入两个只包含小写字母的字符串

输出描述：
输出一个整数，代表最大公共子串的长度

示例1
输入：
asdfas
werasdfaswer

输出：
6
*/

func MaxLength(str1, str2 string) int {
	m, n := len(str1), len(str2)

	//方法一：初始化一个二维切片
	//var dp [][]int
	//for i := 0; i < m+1; i++ {
	//	ar := make([]int, n+1)
	//	dp = append(dp, ar)
	//}

	//方法二：初始化一个二维切片
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	var maxLen int
	for i := 0; i <= m; i++ {
		dp[i][0] = 0
	}
	for i := 0; i <= n; i++ {
		dp[0][i] = 0
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = 0
			}
			if dp[i][j] > maxLen {
				maxLen = dp[i][j]
			}
		}
	}
	return maxLen
}

func main() {
	var s1, s2 string
	fmt.Scanln(&s1)
	fmt.Scanln(&s2)
	fmt.Println(MaxLength(s1, s2))
}
