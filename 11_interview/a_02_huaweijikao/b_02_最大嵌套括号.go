package main

import "fmt"

func BB(s string) int {
	mmap := make(map[string]string, 3)
	mmap[")"] = "("
	mmap["]"] = "["
	mmap["}"] = "{"
	var maxLen, curLen, flag int
	var data []string
	data = append(data, string(s[0]))
	for i := 1; i < len(s); i++ {
		data = append(data, string(s[i]))
		if len(data) >= 2 && mmap[data[len(data)-1]] == data[len(data)-2] {
			data = data[0:len(data)-2]
			flag = 1
		}
		curLen = len(data)
		if maxLen < curLen {
			maxLen = curLen
		}
	}
	if flag == 0 {
		return 0
	} else {
		return maxLen
	}
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(BB(s))
}
