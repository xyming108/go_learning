package main

import (
	"fmt"
)

/*
一个重复字符串是由两个相同的字符串首尾拼接而成，
例如abcabc便是长度为6的一个重复字符串，而abcba则不存在重复字符串。
给定任意字符串，请帮小强找出其中的最长重复子串。


输入描述:
输入一个字符串s，其中s长度小于1e4而且只包含数字和字母。

输出描述:
输出一个整数，表示s的最长重复子串长度，若不存在则输出0

输入例子1:
xabcabcx

输出例子1:
6
*/

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func GetMaxRepeatLength(s string) int {
	res := 0
	for i := 1; i < len(s)/2; i++ {
		count := 0
		for j := 0; j < len(s)-i; j++ {
			if s[j] == s[j+i] {
				count++
			} else {
				count = 0
			}
			if count == i {
				res = max(count*2, res)
			}
		}
	}
	return res
}

func main() {
	//s := "xabcabcx"
	var s string
	fmt.Scanln(&s)
	fmt.Println(GetMaxRepeatLength(s))
}
