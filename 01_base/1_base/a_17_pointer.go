package main

import "fmt"

/*
	&：取地址
	*：根据地址取值
*/

func main() {
	n := 18
	fmt.Println(&n)
	//1、取地址
	p := &n
	fmt.Printf("%T\n", p)
	//2、根据地址取值
	m := *p
	fmt.Printf("%T\n", m)
	fmt.Println(m)
	fmt.Println()

	//new函数申请一个内存地址
	var a = new(int)
	fmt.Println(a)
	fmt.Println(*a)
	*a = 100
	fmt.Println(*a)

	/*
			make和new的区别
			1、make和new都是用来申请内存的
			2、new一般用来给基本数据类型申请内存，如string\int... ，new返回的是指针类型
			3、make是用来给slice、map、chan申请内存的，
		       make函数返回的是对应的这三个类型本身
	*/

}
