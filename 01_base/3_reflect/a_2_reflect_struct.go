package main

import (
	"fmt"
	"reflect"
)

/**
Author: xym
Date: 2021/4/9 17:02
Project: Go_Learning
Description:
*/

//结构体反射

type student struct {
	Name  string `json:"name" ini:"s_name"`
	Score int    `json:"score" ini:"s_score"`
}

func main() {
	stu := student{
		Name:  "小王子",
		Score: 90,
	}

	//通过反射去获取结构体中的所有字段信息
	t := reflect.TypeOf(stu)
	fmt.Printf("name:%v kind:%v\n", t.Name(), t.Kind())

	//根据结构体字段的索引取字段
	for i := 0; i < t.NumField(); i++ {
		filedObj := t.Field(i)
		fmt.Printf("name:%v type:%v tag:%v\n", filedObj.Name, filedObj.Type, filedObj.Tag)
		fmt.Println(filedObj.Tag.Get("json"), filedObj.Tag.Get("ini"))

	}
	fmt.Println()

	//根据名字去取结构体的字段
	filedObj2, ok := t.FieldByName("Score")
	if ok {
		fmt.Printf("name:%v type:%v tag:%v\n", filedObj2.Name, filedObj2.Type, filedObj2.Tag)
	} else {
		fmt.Println("找不到")
	}
}
