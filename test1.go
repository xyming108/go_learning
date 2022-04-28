package main

import "fmt"

func sort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := len(arr) - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}

	return arr
}

func main() {
	a := []int{2, 4, 6, 9, 1}
	fmt.Println(sort(a))
}
