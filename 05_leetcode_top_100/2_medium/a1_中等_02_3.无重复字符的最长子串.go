package main

import "fmt"

/*
给定一个字符串 s ，请你找出其中不含有重复字符的最长子串的长度。

示例1:
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

示例 2:
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例 3:
输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是"wke"，所以其长度为 3。

请注意，你的答案必须是 子串 的长度，"pwke"是一个子序列，不是子串。
*/

/*
滑动窗口法
时间复杂度：O(n)
空间复杂度：O(字符集个数)
*/
func lengthOfLongestSubstring(s string) int {
	//哈希集合，用于去重
	m := map[byte]int{}
	//移动的指针
	rk, ans := -1, 0
	for i := 0; i < len(s); i++ {
		if i != 0 {
			//i每移动一次，把之前的一个删掉
			delete(m, s[i-1])
		}
		//若不重复，则不断向后查找
		for rk+1 < len(s) && m[s[rk+1]] == 0 {
			m[s[rk+1]]++
			rk++
		}
		//比较长度的最大值
		if ans < rk-i+1 {
			ans = rk - i + 1
		}
	}
	return ans
}

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}
