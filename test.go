package main

import "fmt"

func main() {

	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := s[0:5:6]
	fmt.Println(b)
}
