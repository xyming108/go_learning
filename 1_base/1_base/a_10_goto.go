package main

import "fmt"

func main() {
	//goto+label实现跳出多层for循环
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				goto xx //跳到我指定的那个标签
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	}
xx: //label标签
	fmt.Println("over")
}
