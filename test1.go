package main

import (
	"fmt"
)

func main() {
	var zero float64
	inf := 1 / zero
	finf := -1 / zero
	fmt.Println(inf, finf)
	fmt.Println(10 < inf)
	fmt.Println(10 < finf)
	fmt.Printf("%T", inf)
}
