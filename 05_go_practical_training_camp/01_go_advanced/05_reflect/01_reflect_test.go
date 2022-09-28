package _5_reflect

import (
	"fmt"
	"reflect"
	"testing"
)

/*
学习资料
http://c.biancheng.net/view/5131.html
*/

// 反射可以将“接口类型变量”转换为“反射类型对象”
// 反射类型指 reflect.Type 和 reflect.Value
func A() {
	testA := 10
	typ := reflect.TypeOf(testA)
	fmt.Println(typ)
	fmt.Println(typ.String())
	fmt.Println(typ.Name())
	fmt.Println(typ.Kind())
	fmt.Println(typ.Align())
	fmt.Println(typ.Comparable())
	fmt.Println()
}
func B() {
	testB := 10.0
	val := reflect.ValueOf(testB)
	fmt.Println(val)
	fmt.Println(val.Float())
	fmt.Println(val.Kind() == reflect.Float64)
	fmt.Println(val.Type())
	fmt.Println(val.Kind())
	fmt.Println(val.String())
	fmt.Println()
}

// 反射可以将“反射类型对象”转换为“接口类型变量”
func C() {
	testC1 := 12.34
	value1 := reflect.ValueOf(testC1)
	fmt.Println(value1.Interface())
	fmt.Println(value1.Interface().(float64))

	type structTyp struct {
		Name string
	}
	testC2 := structTyp{Name: "xym"}
	value2 := reflect.ValueOf(testC2)
	fmt.Println(value2.Interface())
	fmt.Println(value2.Interface().(structTyp))

	fmt.Println()
}

// 如果要修改“反射类型对象”其值必须是“可写的”
func D() {
	testD1 := 100
	value1 := reflect.ValueOf(testD1)
	fmt.Println(value1.CanSet())

	testD2 := 100
	value2 := reflect.ValueOf(&testD2)
	fmt.Println(value2.CanSet())

	testD3 := 100
	// 要可写就必须传地址
	value3 := reflect.ValueOf(&testD3)
	elem3 := value3.Elem()
	fmt.Println(elem3)
	fmt.Println(elem3.CanSet())
	elem3.SetInt(200)
	fmt.Println(testD3)
	fmt.Println(elem3)
}

func E() {
	type T struct {
		A int
		B string
	}
	t := T{23, "haha"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}

	s.Field(0).SetInt(24)
	s.Field(1).SetString("hi")
	fmt.Println("t:", t)
	fmt.Println()
}

func add(a, b int) int {
	return a + b
}
func F() {
	// 将函数包装为反射值对象
	funcValue := reflect.ValueOf(add)
	// 构造函数参数, 传入两个整型值
	paramList := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20)}
	// 反射调用函数
	retList := funcValue.Call(paramList)
	// 获取第一个返回值, 取整数值
	fmt.Println(retList[0].Int())
}

func TestReflect(t *testing.T) {
	//A()
	//B()
	//C()
	//D()
	//E()
	F()
}
