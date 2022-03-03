package main

import (
	"fmt"
	"time"
)

type field struct {
	name string
}

func (p *field) print() {
	fmt.Println(p.name)
}

// 错误示例
func main() {
	data := []field{{"one"}, {"two"}, {"three"}}
	for _, v := range data {
		go v.print()
	}
	time.Sleep(3 * time.Second)
	// 输出 three three three
}

// 正确示例
/*func main() {
	data := []field{{"one"}, {"two"}, {"three"}}
	for _, v := range data {
		v := v
		go v.print()
	}
	time.Sleep(3 * time.Second)
	// 输出 one two three
}*/

// 正确示例
/*func main() {
	data := []*field{{"one"}, {"two"}, {"three"}}
	for _, v := range data {    // 此时迭代值 v 是三个元素值的地址，每次 v 指向的值不同
		go v.print()
	}
	time.Sleep(3 * time.Second)
	// 输出 one two three
}*/
