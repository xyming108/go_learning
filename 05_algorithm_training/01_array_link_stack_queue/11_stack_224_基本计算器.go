package main

import "fmt"

/*
给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。
s 由数字、'+'、'-'、'('、')'、和 ' ' 组成
注意:不允许使用任何将字符串作为数学表达式计算的内置函数，比如 eval() 。

示例 1：
输入：s = "1 + 1"
输出：2

示例 2：
输入：s = " 2-1 + 2 "
输出：3

示例 3：
输入：s = "(1+(4+5+2)-3)+(6+8)"
输出：23
*/

/*
时间复杂度：O(n)
空间复杂度：O(n)
*/
func calculate(s string) int {
	var ans int
	stack := []int{1}
	sign := 1
	n := len(s)
	for i := 0; i < n; {
		switch s[i] {
		case ' ':
			i++
		case '+':
			sign = stack[len(stack)-1]
			i++
		case '-':
			sign = -stack[len(stack)-1]
			i++
		case '(':
			stack = append(stack, sign)
			i++
		case ')':
			stack = stack[:len(stack)-1]
			i++
		default:
			num := 0
			for ; i < n && '0' <= s[i] && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			ans += sign * num
		}
	}
	return ans
}

func main() {
	s := "(1+(4+5+2)-3)+(6+8)"
	fmt.Println(calculate(s))
}
