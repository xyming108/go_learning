package main

import "fmt"

/*
输入描述：
输入一行，为一个只包含小写字母的字符串。

输出描述：
输出该字符串反转后的字符串。

示例1
输入：
abcd

输出：
dcba
*/

func main() {
	var s string
	fmt.Scanln(&s)
	b := []byte(s)
	for i := 0; i < len(s)/2; i++ {
		b[i], b[len(s)-i-1] = b[len(s)-i-1], b[i]
	}

	fmt.Println(string(b))
}
