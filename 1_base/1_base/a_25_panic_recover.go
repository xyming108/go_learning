package main

import "fmt"

//panic和recover

func funcA() {
	fmt.Println("a")
}

func funcB() {
	defer func() {
		//recover()必须搭配defer使用
		//defer语句一定要在可能引发panic的语句之前定义
		err := recover()
		fmt.Println(err)
		fmt.Println("1")
		fmt.Println("2")
		fmt.Println("3")
	}()
	panic("出现了严重的错误") //程序会报错，崩溃退出，后面的语句都不会执行
	fmt.Println("b")
}

func funcC() {
	fmt.Println("c")
}

func test11() {
	defer func() { //必须是最后接收的recover()才有效
		fmt.Println(recover()) //有效	test panic
	}()
	defer recover()              //无效！
	defer fmt.Println(recover()) //无效！	nil
	defer func() {
		func() {
			println("defer inner") //defer inner
			recover()              //无效！
		}()
	}()

	panic("test panic")
}

func main() {
	/*funcA()
	funcB()
	funcC()*/

	test11()
}
