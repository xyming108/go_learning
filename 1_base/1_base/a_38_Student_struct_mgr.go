package main

import (
	"fmt"
)

/**
Author: 1_base
Date: 2021/4/6 15:06
Project: Go_Learning
Description:
*/

//学生管理系统
//特性：
//1、保存了一些数据 --->结构体的字段
//2、有4个功能     --->结构体的方法

type student8 struct {
	id   int
	name string
}

//造一个学生的管理者
type studentMgr struct {
	allStudent map[int]student8
}

//查看学生
func (s studentMgr) showStudents() {
	for _, stu := range s.allStudent {
		fmt.Printf("学号：%d 姓名:%s\n", stu.id, stu.name)
	}
}

//增加学生
func (s studentMgr) addStudents() {
	var (
		stuId   int
		stuName string
	)
	fmt.Println("请输入学号：")
	fmt.Scanln(&stuId)
	fmt.Println("请输入姓名：")
	fmt.Scanln(&stuName)
	newStu := student8{
		id:   stuId,
		name: stuName,
	}
	s.allStudent[newStu.id] = newStu
	fmt.Println("添加成功")
}

//修改学生
func (s studentMgr) editStudents() {
	var stuId int
	fmt.Println("请输入要修改的编号")
	fmt.Scanln(&stuId)
	stuObj, ok := s.allStudent[stuId]
	if !ok {
		fmt.Println("查无此人")
		return
	}
	fmt.Printf("你要修改的学生信息如下：学号：%d 姓名：%s\n", stuObj.id, stuObj.name)
	fmt.Println("请输入学生的新名字：")
	var newName string
	fmt.Scanln(&newName)
	stuObj.name = newName
	s.allStudent[stuId] = stuObj
}

//删除学生
func (s studentMgr) deleteStudents() {
	var stuId int
	fmt.Println("请输入要删除学生的编号")
	fmt.Scanln(&stuId)
	_, ok := s.allStudent[stuId]
	if !ok {
		fmt.Println("查无此人")
		return
	}
	delete(s.allStudent, stuId)
	fmt.Println("删除成功！！")
}
