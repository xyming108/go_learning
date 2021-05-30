package main

import "fmt"

//闭包
//闭包是一个函数，这个函数包含了他外部作用域的一个变量

func ff1(f func()) {
	fmt.Println("ff1")
	f()
}

func ff2(x, y int) {
	fmt.Println("ff2")
	fmt.Println(x + y)
}

func ff3(f func(int, int), x, y int) func() {
	tmp := func() {
		f(x, y)
	}
	return tmp
}

func adder(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

//---------------------------------------
func a5() func() int {
	i := 0
	b := func() int {
		i++
		fmt.Println(i)
		return i
	}
	return b
}

func main() {
	/*	ret := adder(100)
		ret2 := ret(200)
		fmt.Printf("%T\n", ret)
		fmt.Printf("%T\n", ret2)
		fmt.Println(ret2)
		fmt.Println()

		pp := ff3(ff2, 100, 300)
		ff1(pp)*/

	c := a5()
	c()
	//此时c() == b， 即
	/*
		func() int {
				i++
				fmt.Println(i)
				return i
			}
		之后的c()都是如此
	*/
	c()
	c()
	c()

	//a5() //不会输出i
}
