package main

import "fmt"

type S1 struct {
	s string
	i int
	b bool
	p *int
	//slice []int
}

type S2 struct {
	s string
	i int
	b bool
	p *int
	slice []int
}

func main() {
	//var s1 S1
	//var s2 S2
	//
	//s3 := S1(s2)
	//fmt.Println(s3 == s1)
	m := make(map[S1]string, 0)
	fmt.Println(m)
}
