package main

import "fmt"

func quickSort(a []int, left, right int) {
	if left >= right {
		return
	}
	l, r := left, right
	pivot := a[left]
	for left < right {
		for left < right && a[right] >= pivot {
			right--
		}
		a[left] = a[right]
		for left < right && a[left] <= pivot {
			left++
		}
		a[right] = a[left]
	}
	a[left] = pivot
	middle := left
	quickSort(a, l, middle-1)
	quickSort(a, middle+1, r)
}

func main() {
	a := []int{1, 4, 7, 2, 8, 3, 3, 9, 0, 1, 10}
	quickSort(a, 0, len(a)-1)
	fmt.Println(a)
}
