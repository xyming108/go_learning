package main

import "fmt"

/*
可以通过拷贝临时 slice 的数据，而不是重新切片来解决
*/

func gget() (res []byte) {
	raw := make([]byte, 10000)
	fmt.Println(len(raw), cap(raw), &raw[0]) // 10000 10000 0xc420080000
	res = make([]byte, 3)
	copy(res, raw[:3])
	return
}

func main() {
	data := gget()
	fmt.Println(len(data), cap(data), &data[0]) // 3 3 0xc4200160b8
}
