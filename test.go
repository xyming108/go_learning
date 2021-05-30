package main

import "fmt"

func main() {
	var a = new(int)
	*a = 100
	fmt.Println(*a)

	var b = make(map[string]int, 10)
	b["沙河娜扎"] = 100
	fmt.Println(b)
}
