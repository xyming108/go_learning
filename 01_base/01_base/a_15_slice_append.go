package main

import "fmt"

//append()为切片追加元素

func main() {
	s1 := []string{"北京", "上海", "广州"}
	s1[2] = "深圳"
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

	//调用append函数必须用原来的切片变量接收返回值
	s1 = append(s1, "杭州")
	s1 = append(s1, "西安", "贵州")
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	ss := []string{"武汉", "西安", "浙江"}
	s1 = append(s1, ss...) //...表示把数组中的元素单独拆开
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

	//关于append删除切片中的某个元素
	a1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	a2 := a1[:]
	//删除索引为1的那个
	a2 = append(a2[0:1], a2[2:]...)
	fmt.Println(a2)
	fmt.Println(a1)
}
