package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	fmt.Println()
	rawStr := "我是原文"
	hash := sha256.New()
	hash.Write([]byte(rawStr))
	fmt.Println(hash.Size()) //32
	result := hex.EncodeToString(hash.Sum(nil))
	fmt.Println(result) //f0bbaad46a079bf0ec6d64da697de44f6c0e7aa8bd402d80b27b0808c1335323
}
