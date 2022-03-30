package main

import (
	"fmt"
)

/*
给你一个字符串 s，找到 s 中最长的回文子串。

示例 1：
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。

示例 2：
输入：s = "cbbd"
输出："bb"
*/

/*
暴力解法
时间复杂度：O(n³)
空间复杂度：O(1)

题解：https://leetcode-cn.com/problems/longest-palindromic-substring/solution/zui-chang-hui-wen-zi-chuan-by-leetcode-solution/
*/
func longestPalindrome1(s string) string {
	length := len(s)
	if length < 2 {
		return s
	}
	maxLen, begin := 1, 0
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if j-i+1 > maxLen && validPalindrome(s, i, j) {
				maxLen = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxLen]
}

func validPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

/*
中心扩散法
时间复杂度：O(n²)
空间复杂度：O(1)
*/
func longestPalindrome2(s string) string {
	length := len(s)
	if length < 2 {
		return s
	}
	maxLen, begin := 1, 0
	for i := 0; i < length-1; i++ {
		oddLen := expendAroundCenter(s, i, i)
		evenLen := expendAroundCenter(s, i, i+1)
		if max(oddLen, evenLen) > maxLen {
			maxLen = max(oddLen, evenLen)
			begin = i - (maxLen-1)/2
		}
	}
	return s[begin : begin+maxLen]
}

func expendAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) {
		if s[left] == s[right] {
			left--
			right++
		} else {
			break
		}
	}
	return right - left - 1
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

/*
动态规划
时间复杂度：O(n²)
空间复杂度：O(n²)
*/
func longestPalindrome3(s string) string {
	length := len(s)
	if length < 2 {
		return s
	}

	maxLen, begin := 1, 0
	dp := make([][]bool, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)
	}

	for i := 0; i < length; i++ {
		dp[i][i] = true
	}

	for L := 2; L <= length; L++ {
		for i := 0; i < length; i++ {
			j := L + i - 1
			if j >= length {
				break
			}

			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				begin = i
			}
		}
	}

	return s[begin : begin+maxLen]
}

func main() {
	s := "babad"
	fmt.Println(longestPalindrome1(s))
	fmt.Println(longestPalindrome2(s))
	fmt.Println(longestPalindrome3(s))
}
