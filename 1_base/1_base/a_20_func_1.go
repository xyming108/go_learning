package main

import "fmt"

//函数

//没有返回值
func f1(x int, y int) {
	fmt.Println(x + y)
}

//没有参数、没有返回值
func f2() {
	fmt.Println("f2")
}

//没有参数但有返回值
func f3() int {
	return 3
}

//有返回值，a可以在内部使用
func sum(x int, y int) (a int) {
	/*return x+y*/
	a = x + y
	return a //此处a可以省略
}

//多个返回值
func f4() (int, int) {
	return 1, 2
}

//参数类型的简写
func f5(x, y int) int {
	return x + y
}

//可变长参数
func f6(x string, y ...int) {
	fmt.Print(x)
	fmt.Println(y) //y的类型是切片 []int
}

func main() {
	ret := sum(1, 5)
	fmt.Println(ret)
	f1(1, 3)
	f2()
	fmt.Println(f3())
	m, n := f4()
	fmt.Println(m, n)
	i := f5(6, 2)
	fmt.Println(i)
	f6("sad ", 1, 2, 3)

}
