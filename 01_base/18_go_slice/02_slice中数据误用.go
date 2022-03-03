package main

import "bytes"

/*
举个简单例子，重写文件路径（存储在 slice 中）
分割路径来指向每个不同级的目录，修改第一个目录名再重组子目录名，创建新路径：

拼接的结果不是正确的 AAAAsuffix/BBBBBBBBB，因为 dir1、 dir2 两个 slice 引用的数据都是 path 的底层数组，
第 13 行修改 dir1 同时也修改了 path，也导致了 dir2 的修改
*/

// 错误使用 slice 的拼接示例
func main() {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/') // 4
	println(sepIndex)

	dir1 := path[:sepIndex]
	dir2 := path[sepIndex+1:]
	println("dir1: ", string(dir1)) // AAAA
	println("dir2: ", string(dir2)) // BBBBBBBBB

	dir1 = append(dir1, "suffix"...)
	println("current path: ", string(path)) // AAAAsuffixBBBB

	path = bytes.Join([][]byte{dir1, dir2}, []byte{'/'})
	println("dir1: ", string(dir1)) // AAAAsuffix
	println("dir2: ", string(dir2)) // uffixBBBB

	println("new path: ", string(path)) // AAAAsuffix/uffixBBBB    // 错误结果
}
