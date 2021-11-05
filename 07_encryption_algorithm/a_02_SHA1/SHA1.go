package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func main() {
	fmt.Println()
	rawStr := "我是原文"
	hash := sha1.New()
	//SHA-1可以生成一个被称为消息摘要的160位（20字节）散列值，散列值通常的呈现形式为40个十六进制数
	fmt.Println(hash.Size()) //20
	hash.Write([]byte(rawStr))
	//16进制显示结果
	result := hex.EncodeToString(hash.Sum(nil))
	fmt.Println(result) //959acf97bfe0d42ba5f1474c26b4160144226639
}
