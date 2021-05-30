package __package_calc

import "fmt"

/**
Author: 1_base
Date: 2021/4/8 22:39
Project: Go_Learning
Description:
*/

func init() {
	fmt.Println("import 时自动执行！")
}

// Add 包中的标识符（变量名、函数名、结构体、接口）如果首字母小写，表示私有
//首字母大写的标识符可以对外部的包调用
func Add(x, y int) int {
	return x + y
}
