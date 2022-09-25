package _4_memory_leak

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	initSlice := []int{3, 4, 5, 6, 7}
	//partSlice := initSlice[1:3]

	partSlice := make([]int, 2)
	copy(partSlice, initSlice[1:3])

	fmt.Printf("initSlice addr: %p", &initSlice)
	fmt.Println()
	fmt.Printf("partSlice addr: %p", &partSlice)
	fmt.Println()

	for i := 0; i < len(initSlice); i++ {
		fmt.Printf("%v:[%v]  ", initSlice[i], &initSlice[i])
	}
	fmt.Println()
	for i := 0; i < len(partSlice); i++ {
		fmt.Printf("%v:[%v]  ", partSlice[i], &partSlice[i])
	}
	fmt.Println()
}
