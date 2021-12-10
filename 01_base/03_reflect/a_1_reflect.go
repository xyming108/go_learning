package main

import (
	"fmt"
	"reflect"
)

/**
Author: xym
Date: 2021/4/9 15:16
Project: Go_Learning
Description:
*/

func reflectType(x interface{}) {
	//现在不知道别人调用这个函数的时候会传进来什么类型的变量
	//1、方式一：通过类型断言
	//2、方式二：借助反射
	obj := reflect.TypeOf(x)
	fmt.Println(obj, obj.Name(), obj.Kind(), obj.Size(), obj.String())
	fmt.Printf("%T\n", obj)
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Printf("%v %T\n", v, v)
	k := v.Kind() //拿到值对应的类型种类
	//如何得到一个传入时候类型的变量
	switch k {
	case reflect.Float32:
		ret := float32(v.Float())
		fmt.Printf("%v %T\n", ret, ret)
	case reflect.Int32:
		ret := int32(v.Int())
		fmt.Printf("%v %T\n", ret, ret)
	}
}

func reflectValue1(x interface{}) {
	v := reflect.ValueOf(x)
	//Elem()用来根据指针取对应的值
	k := v.Elem().Kind() //拿到值对应的类型种类
	switch k {
	case reflect.Float32:
		v.Elem().SetFloat(3.14)
	case reflect.Int32:
		v.Elem().SetInt(1000)
	}
}

type Cat struct {
	name string
}

type Dog struct {
	age int
}

func main() {
	//基本数据类型
	var a float32 = 1.234
	reflectType(a)
	var b int8 = 18
	reflectType(b)
	fmt.Println()

	//结构体类型
	var c Cat
	reflectType(c)
	var d Dog
	reflectType(d)
	fmt.Println()

	//slice
	var e []int
	reflectType(e)
	var f []string
	reflectType(f)
	fmt.Println()

	var h float32 = 100
	reflectValue(h)
	fmt.Println()

	fmt.Println("----------")
	var i int32 = 10
	reflectValue1(&i)
	fmt.Printf("%v %T\n", i, i)
	fmt.Println("-----------")

	//IsNil用于判断指针是否为空，IsValid用于判断返回值是否有效
	var j *int
	fmt.Println(reflect.ValueOf(j).IsNil())
	fmt.Println(reflect.ValueOf(i).IsValid())
}
