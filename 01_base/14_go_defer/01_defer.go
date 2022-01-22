package main

import (
	"fmt"
	"time"
)

func f0() int {
	i := 5
	defer func() {
		fmt.Println("---0")
		i++
	}()
	return i
}

func f1() (result int) {
	defer func() {
		fmt.Println("----1")
		result++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		fmt.Println("-----2")
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		fmt.Println("------3")
		r = r + 5
	}(r)
	return 1
}

func foo() {
	i := 1
	go func() {
		i++
	}()
	time.Sleep(1 * time.Second)
	println(i)
}

func bar() {
	i := 1
	go func(i int) {
		i++
	}(i)
	time.Sleep(1 * time.Second)
	println(i)
}

func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

func f12() *int {
	a := 10
	b := &a
	fmt.Println("b", b)
	defer func() {
		c := 12
		b = &c
		fmt.Println("defer", b)
	}()
	return b
}

func main() {
	//println(f0())
	//println(f1())
	//println(f2())
	//println(f3())
	//foo()
	//bar()
	//a()

	//fmt.Println("===============")
	//aa := [3]int{1, 2, 3}
	//defer fmt.Println(aa)
	//aa[0] = 100
	//
	//b := []int{1, 2, 3}
	//defer fmt.Println(b)
	//b[0] = 200

	bb := f12()
	fmt.Println("main", bb, " ", *bb)
}
