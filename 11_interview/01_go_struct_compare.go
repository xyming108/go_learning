package main

import (
	"fmt"
	"reflect"
)

type S struct {
	s     string
	i     int
	b     bool
	p     *int
	slice []int
}

func main() {
	s1 := S{
		s:     "你好",
		i:     1,
		b:     false,
		slice: []int{1, 2},
	}

	s2 := S{
		s:     "你好",
		i:     1,
		b:     false,
		slice: []int{1, 2},
	}

	fmt.Println(reflect.DeepEqual(s1, s2))
}
