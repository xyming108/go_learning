package main

import "bytes"

/*
解决方法：
重新分配新的 slice 并拷贝你需要的数据
使用完整的 slice 表达式：input[low:high:max]，容量便调整为 max - low

第 6 行中第三个参数是用来控制 dir1 的新容量，再往 dir1 中 append 超额元素时，
将分配新的 buffer 来保存。而不是覆盖原来的 path 底层数组
*/

// 使用 full slice expression
func main() {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/') // 4
	dir1 := path[:sepIndex:sepIndex]       // 此时 cap(dir1) 指定为4， 而不是先前的 16
	dir2 := path[sepIndex+1:]
	dir1 = append(dir1, "suffix"...)

	path = bytes.Join([][]byte{dir1, dir2}, []byte{'/'})
	println("dir1: ", string(dir1))     // AAAAsuffix
	println("dir2: ", string(dir2))     // BBBBBBBBB
	println("new path: ", string(path)) // AAAAsuffix/BBBBBBBBB
}
