package main

import "fmt"

func main() {
	/*
		byte：针对一个字符
		rune：针对含有中文，韩文，日文等一系列字符串
	*/

	s := "hello 서예명哈哈哈"
	n := len(s) //len()求的是byte字节的数量，一个中文或韩文占3个字节
	fmt.Println(n)

	for i := 0; i < n; i++ {
		fmt.Print(s[i], " ")
		fmt.Printf("%c ", s[i])
	}
	fmt.Println()
	for _, c := range s { //从字符串中拿出具体字符
		fmt.Printf("%c ", c)
	}
	fmt.Println()
	fmt.Println()

	//************************************
	//字符串不能修改，只能转换成rune或byte类型修改后在转换回string类型
	s1 := "白菜啊"
	s2 := []rune(s1) //把字符串强制转换成一个rune切片
	fmt.Println(s2)
	s2[0] = '青'
	fmt.Println(string(s2)) //把rune切片强制转换成字符串
	fmt.Println(s2)
	fmt.Println()

	aa1 := "big"
	aa2 := []byte(aa1)
	fmt.Println(aa2)
	aa2[0] = 'p'
	fmt.Println(string(aa2))
	aa3 := byte('H')
	fmt.Printf("%T\n", aa3)
	fmt.Println()

	c1 := "红" //string类型
	c2 := '红' //rune(int32)类型
	fmt.Printf("c1:%T, c2:%T\n", c1, c2)
	c3 := "H" //string类型
	c4 := 'H' //int32类型
	fmt.Printf("c1:%T, c2:%T\n", c3, c4)
	fmt.Println(c1, c2, c3, c4)
	fmt.Println()

	//类型转换
	nn := 10
	var f float64
	f = float64(nn)
	fmt.Println(f)
	fmt.Printf("%T\n", f)

}
