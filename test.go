package main

import "fmt"

func a(flag *bool) {
	*flag = true
}

func main() {
	var flag bool
	fmt.Println(flag)
	a(&flag)
	fmt.Println(flag)
}
