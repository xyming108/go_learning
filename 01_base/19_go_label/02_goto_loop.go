package main

import "fmt"

func main() {
	for {
		switch {
		case true:
			fmt.Println("breaking out...")
			//break    // 死循环，一直打印 breaking out...
			goto loop
		}
	}
loop:
	fmt.Println("out...")
}
