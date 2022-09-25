package main

import (
	"bytes"
	"fmt"
)

func main() {
	var sqlStr bytes.Buffer
	sqlStr.WriteString("(123, 2131, 21312,")


	byt := sqlStr.Bytes()
	lastElem := string(byt[len(byt)-1])
	if lastElem == "," {
		byt = byt[0 : len(byt)-1]
		sqlStr.Reset()
		sqlStr.WriteString(string(byt))
	}
	sqlStr.WriteString(");")

	fmt.Println(sqlStr.String())
}
