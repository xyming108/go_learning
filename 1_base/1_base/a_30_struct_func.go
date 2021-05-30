package main

import (
	"fmt"
)

/**
Author: 1_base
Date: 2021/4/5 18:25
Project: Go_Learning
Description:
*/

type person1 struct {
	name string
	age  int
}

type dog struct {
	name string
}

//构造函数
//返回的是结构体还是结构体指针
//当结构体比较大的时候尽量使用结构体指针，减少程序的内存开销
func newPerson(name string, age int) person1 {
	return person1{
		name: name,
		age:  age,
	}
}

func newDog(name string) dog {
	return dog{
		name: name,
	}
}

func main() {
	p1 := newPerson("元帅", 18)
	p2 := newPerson("周芷若", 9000)
	fmt.Println(p1, p2)

	d1 := newDog("狗")
	fmt.Println(d1)
}
