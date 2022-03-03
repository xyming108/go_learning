package main

import "fmt"

/*
从 slice 中重新切出新 slice 时，新 slice 会引用原 slice 的底层数组。
如果跳了这个坑，程序可能会分配大量的临时 slice 来指向原底层数组的部分数据，将导致难以预料的内存使用。
*/

func get() []byte {
	raw := make([]byte, 10000, 10000)
	fmt.Println(len(raw), cap(raw), &raw[0]) // 10000 10000 0xc420080000
	return raw[:3]                           // 重新分配容量为 10000 的 slice
}

func main() {
	data := get()
	fmt.Println(len(data), cap(data), &data[0]) // 3 10000 0xc420080000
}
