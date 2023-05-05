package main

import (
	"crypto/sha1"
	"fmt"
	"io"
)

func Encrypt(s string) string {
	h := sha1.New()
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func main() {
	a := "ttt.xu@anker.com"

	encrypt := Encrypt(a)
	fmt.Println(encrypt)
}
