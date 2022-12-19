package main

import "fmt"

/*
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

示例 1：
输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]

示例 2：
输入：n = 1
输出：["()"]
*/

/*
递归，分治
*/
func generateParenthesis(n int) []string {
	var result []string
	var backtrace func(left, right int, str string)
	backtrace = func(leftRemain, rightRemain int, str string) {
		if len(str) == 2*n { //递归出口
			result = append(result, str)
			return
		}
		if leftRemain > 0 { //剪枝条件1
			backtrace(leftRemain-1, rightRemain, str+"(")
		}
		if leftRemain < rightRemain { //剪枝条件2
			backtrace(leftRemain, rightRemain-1, str+")")
		}
	}
	backtrace(n, n, "")
	return result
}

func main() {
	fmt.Println(generateParenthesis(3))
}
