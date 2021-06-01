package main

import "fmt"

/**
Author: 1_base
Date: 2021/4/5 19:40
Project: Go_Learning
Description:
*/

//方法
//标识符：变量名、函数名、类型名、方法名
//Go语言中如果标识符首字母是大写的，就表示对外部可见（暴露的，公有的）

type dog1 struct {
	name string
}

type person2 struct {
	name string
	age  int
}

//构造函数
func newDog1(name string) dog1 {
	return dog1{
		name: name,
	}
}

func newPerson2(name string, age int) person2 {
	return person2{
		name: name,
		age:  age,
	}
}

//方法是作用于特定类型的函数
//接受者表示的是调用该方法的具体类型变量，多用类型名首字母小写表示
func (d dog1) wang() {
	fmt.Printf("%s：汪汪汪", d.name)
}

//值接收者，传拷贝进去
func (p person2) guonian() {
	p.age++
}

/*
	指针接收者使用场景：
		需要修改接收者中的值
		接收者是拷贝代价比较大的大对象
		保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者
*/
//指针接收者，传内存地址进去
func (p *person2) zhizhenguonian() {
	(*p).age++
}

func main() {
	/*d1 := newDog1("哈士奇")
	d1.wang()*/

	p2 := newPerson2("元帅", 18)
	fmt.Println(p2.age) //18
	p2.guonian()
	fmt.Println(p2.age) //18
	p2.zhizhenguonian()
	fmt.Println(p2.age) //19
}
