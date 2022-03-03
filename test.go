package main

import "fmt"

type data struct {
	num   int
	key   *string
	items map[string]bool
}

func (this *data) pointerFunc() {
	this.num = 7
}

func (this data) valueFunc() {
	this.num = 8
	*this.key = "valueFunc.key"
	this.items["valueFunc"] = true
}

func main() {
	key := "key1"

	d := data{1, &key, make(map[string]bool)}
	fmt.Printf("num=%v  key=%v  items=%v\n", d.num, *d.key, d.items)

	d.pointerFunc() // 修改 num 的值为 7
	fmt.Printf("num=%v  key=%v  items=%v\n", d.num, *d.key, d.items)

	d.valueFunc() // 修改 key 和 items 的值
	fmt.Printf("num=%v  key=%v  items=%v\n", d.num, *d.key, d.items)
}
