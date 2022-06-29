package main

import (
	"fmt"
	sort2 "sort"
)

func main() {
	a := []int{1, 4, 2, 6, 3, 6}
	//sort2.Ints(a)
	sort2.Slice(a, func(i, j int) bool {
		return a[i]>a[j]
	})
	fmt.Println(a)
}
