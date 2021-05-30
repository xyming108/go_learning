package main

import (
	"fmt"
)

func main() {
	maxFloat32 := 3.14
	fmt.Printf("%T\n", maxFloat32) //默认go语言中小数都是64位的
	f2 := float32(maxFloat32)      //强制转换
	fmt.Printf("%T\n", f2)

}
