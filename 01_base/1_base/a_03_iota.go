package main

import "fmt"

const (
	n1 = 100 //100
	n2       //100
	n3       //100
)

//iota
const (
	b1 = iota //0
	b2        //1
	b3        //2
	b4        //3
	_         //值丢弃了
	b5        //5
	b6 = 100  //100
	b7        //100
	b8 = iota //8
	b9        //9
)

func main() {
	fmt.Println(n1, n2, n3)
	fmt.Println(b1, b2, b3, b4, b5, b6, b7, b8, b9)
}
