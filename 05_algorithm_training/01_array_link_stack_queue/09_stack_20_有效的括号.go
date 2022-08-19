package main

import "fmt"

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。

示例 1：
输入：s = "()[]{}"
输出：true

示例 2：
输入：s = "{[]}"
输出：true

示例 3：
输入：s = "([)]"
输出：false
*/

//时间复杂度：O(N)
//空间复杂度：O(N+E)，E表示括号的个数
//栈的思想
func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	mapTemp := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []byte
	for i := 0; i < len(s); i++ {
		if v, ok := mapTemp[s[i]]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != v {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

func main() {
	s := "()[]{}"
	b := isValid(s)
	fmt.Println(b)
}
