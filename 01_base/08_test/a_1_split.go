package main

import "strings"

/**
Author: xym
Date: 2021/4/11 12:06
Project: Go_Learning
Description:
*/

// Split Split将s按照sep进行分割，返回一个字符串的切片
func Split(s, sep string) (ret []string) {
	idx := strings.Index(s, sep)
	for idx > -1 {
		ret = append(ret, s[:idx])
		s = s[idx+len(sep):]
		idx = strings.Index(s, sep)
	}
	ret = append(ret, s)
	return
}
