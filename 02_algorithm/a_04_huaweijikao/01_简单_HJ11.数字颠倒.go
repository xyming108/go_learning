package main

import (
	"fmt"
)

/*
输入描述：
输入一个int整数

输出描述：
将这个整数以字符串的形式逆序输出

示例1
输入：
1516000

输出：
0006151
*/

func main() {
	var input string
	fmt.Scanln(&input)
	for i := len(input) - 1; i >= 0; i-- {
		fmt.Print(string(input[i]))
	}
}
