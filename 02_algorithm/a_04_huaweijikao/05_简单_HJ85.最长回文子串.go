package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
描述
给定一个仅包含小写字母的字符串，求它的最长回文子串的长度。
所谓回文串，指左右对称的字符串。
所谓子串，指一个字符串删掉其部分前缀和后缀（也可以不删）的字符串

输入描述：
输入一个仅包含小写字母的字符串
输出描述：
返回最长回文子串的长度

示例1
输入：
cdabbacc

输出：
4

说明：
abba为最长的回文子串
*/

func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	start, end := 0, 0
	for i := 0; i < n; {
		l, r := i, i
		//如果字符串相同则分别从前一个和后一个开始回文
		for r < n-1 && s[r] == s[r+1] {
			r++
		}
		i = r + 1
		for l > 0 && r < n-1 && s[l-1] == s[r+1] {
			l--
			r++
		}
		if end < r-l {
			start = l
			end = r - l
		}
	}
	return s[start : start+end+1]
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		str := longestPalindrome(input.Text())
		fmt.Println(len(str))
	}
}
