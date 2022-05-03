package main

import "fmt"

func main() {
	a := 1 << 3
	fmt.Printf("%b\n", a)
	fmt.Println(a)
	b := 6 >> 1
	fmt.Printf("%b\n", b)
	fmt.Println(b)
	fmt.Println(b&(1))
}
