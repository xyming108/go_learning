package main

import (
	"fmt"
)

/**
Author: 01_base
Date: 2021/4/8 22:40
Project: Go_Learning
Description:
*/

//包的路径从：GOPATH/src 后面的路径开始写起
//想要被别的包调用的标识符都要首字母大写
//单行导入好多行导入
//导入包时可以指定别名
//导入包不想使用包内部的标识符需要使用匿名导入
//每个包导入的时候会自动执行一个名为init()的函数，它没有参数也没有返回值也不能手动调用
//多个包中都定义了init()函数，则它们的执行顺序从头部开始执行

import (
	//自定义包名：import 别名 "包的路径"
	cc "01_base/2_package_calc"
	//匿名导入：只是希望导入包，但是不使用内部的任何数据
	//import _ "包的路径"
)

/*
执行顺序
	：全局声明 -> init() -> main()
*/

var x int8 = 10

const pi = 3.14

func init() {
	fmt.Println("自动执行！")
	fmt.Println(x, pi)
}

func main() {
	ret := cc.Add(10, 20)
	fmt.Println(ret)
}
