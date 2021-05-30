package main

import (
	"fmt"
	"strings"
)

func main() {
	//Go语言字符串只能用双引号
	var s = "气我判决书查克拉"
	fmt.Println(s)

	//多行的字符串
	ss := `
高处不胜寒
	起舞弄清影
	`
	fmt.Println(ss)

	//*********************************
	//字符串的常用操作
	p1 := "小星大星啊"
	fmt.Println(len(p1)) //一个中文3字节
	p2 := "你好"
	p3 := p1 + p2 //拼接
	fmt.Println(p3)
	ss1 := fmt.Sprintf("%s%s", p1, p2) //拼接，并且有返回值
	fmt.Println(ss1)

	ss2 := strings.Split(p1, "星") //分隔
	fmt.Printf("%T\n", ss2)
	fmt.Println(ss2)

	ss3 := strings.Contains(p1, "星") //是否包含
	fmt.Println(ss3)

	ss4 := strings.HasPrefix(p1, "小") //前缀判断
	fmt.Println(ss4)
	ss5 := strings.HasSuffix(p1, "啊") //后缀判断
	fmt.Println(ss5)

	s11 := "abcdefg"
	ss6 := strings.Index(s11, "a")     //第一次出现的位置下标
	ss7 := strings.LastIndex(s11, "g") //最后一次出现的位置下标
	fmt.Printf("%d, %d\n", ss6, ss7)

	ss8 := strings.Join(ss2, "+") //拼接
	fmt.Println(ss8)

}
