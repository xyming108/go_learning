package main

import (
	"fmt"
	"os"
)

/**
Author: 1_base
Date: 2021/4/6 9:51
Project: Go_Learning
Description:
*/

/*
	函数版学生管理系统：
		查看、新增、删除
*/

var (
	allStudent map[int64]*student
)

type student struct {
	id   int64
	name string
}

//构造函数
func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

func showAllStudent() {
	//把所有学生打印出来
	for k, v := range allStudent {
		fmt.Printf("学号：%d 姓名：%s\n", k, v.name)
	}
}

func addStudent() {
	var (
		id   int64
		name string
	)
	fmt.Println("请输入学生的学号：")
	fmt.Scanln(&id)
	fmt.Println("请输入学生姓名：")
	fmt.Scanln(&name)
	newStu := newStudent(id, name)
	allStudent[id] = newStu
}

func deleteStudent() {
	var (
		deleteId int64
	)
	fmt.Println("请输入要删除学生的学号：")
	fmt.Scanln(&deleteId)
	delete(allStudent, deleteId)
}

func main() {
	allStudent = make(map[int64]*student, 48) //初始化，开辟内存空间
	for {
		//打印菜单
		fmt.Println("欢迎光临学生管理系统！")
		fmt.Println(`
		1.查看所有学生
		2.新增学生
		3.删除学生
		4.退出
		`)
		fmt.Println("请输入你要选择的功能...")
		//等待用户选择要做什么
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了%d这个选项\n", choice)
		//执行对应的函数
		switch choice {
		case 1:
			showAllStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("结束")
		}
	}
}
