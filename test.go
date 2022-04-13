package main

import "fmt"

func main() {
	a := []int{5, 4, 6, 2}
	b := []int{0, 0, 0, 0, 0, 0}
	copy(b, a)
	fmt.Println(a, b)
}
