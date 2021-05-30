package main

import (
	"fmt"
)

//defer：defer语句会将其后面跟随的语句进行延迟处理，并逆序执行
/*
	1. 关键字 defer 用于注册延迟调用。
    2. 这些调用直到 return 前才被执。因此，可以用来做资源清理。
    3. 多个defer语句，按先进后出的方式执行。
    4. defer语句中的变量，在defer声明时就决定了。
*/

func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("end")
}

/*
	结果：
		start
		end
		3
		2
		1

*/

func f11() int { //结果为5
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f22() (x int) { //返回值为6
	defer func() {
		x++
	}()
	return 5
}

func main() {
	deferDemo()
	fmt.Println()

	fmt.Println(f11())
	fmt.Println(f22())
}
