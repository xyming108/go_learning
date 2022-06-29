package main

import "fmt"

/*
给你一个字符串 s ，请你统计并返回这个字符串中 回文子串 的数目。
回文字符串 是正着读和倒过来读一样的字符串。
子字符串 是字符串中的由连续字符组成的一个序列。
具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
s 由小写英文字母组成

示例 1：
输入：s = "abc"
输出：3
解释：三个回文子串: "a", "b", "c"

示例 2：
输入：s = "aaa"
输出：6
解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
*/

/*
中心拓展
时间复杂度：O(n^2)
空间复杂度：O(1)
*/
func countSubstrings(s string) int {
	n := len(s)
	ans := 0
	for i := 0; i < 2*n-1; i++ {
		l, r := i/2, i/2+i%2
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
			ans++
		}
	}
	return ans
}

/*
动态规划法
时间复杂度：O(n^2)
空间复杂度：O(n^2)
*/
func countSubstrings2(s string) int {
	n := len(s)
	ans := 0
	check := make([][]bool, n)
	for i := range check {
		check[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			if s[i] == s[j] && (j-i < 2 || check[i+1][j-1]) {
				ans++
				check[i][j] = true
			}
		}
	}
	return ans
}

func main() {
	s := "aaa"
	fmt.Println(countSubstrings(s))
	fmt.Println(countSubstrings2(s))
}
