package main

import (
	"fmt"
	"os"
)

/**
Author: 01_base
Date: 2021/4/6 18:47
Project: Go_Learning
Description:
*/

//学生管理系统
var smr studentMgr //声明一个全局的变量：学生管理对象

//菜单函数
func showMenu() {
	fmt.Println("---------------Welcome to sms---------------")
	fmt.Println(`
	1、查看所有学生
	2、添加学生
	3、修改学生
	4、删除学生
	5、退出
	`)
}

func main() {
	//初始化
	smr = studentMgr{ //修改的全局那个变量
		allStudent: make(map[int]student8, 100),
	}
	for {
		showMenu()
		fmt.Println("请输入要执行的编号...")
		var choice int
		fmt.Scanln(&choice)
		fmt.Println("你输入的是：", choice)
		switch choice {
		case 1:
			smr.showStudents()
		case 2:
			smr.addStudents()
		case 3:
			smr.editStudents()
		case 4:
			smr.deleteStudents()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("退出")
		}
	}
}
