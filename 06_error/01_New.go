package main

import (
	error "errors"
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	e := error.New("errors中New一个错误")
	fmt.Println(e)
	fmt.Printf("%+v\n", e)
	fmt.Println()
	/*
		errors中New一个错误
		errors中New一个错误
	*/

	ee := errors.New("github.com/pkg/errors中New一个错误")
	fmt.Println(ee)
	fmt.Printf("%+v\n", ee)
	/*
		github.com/pkg/errors中New一个错误
		github.com/pkg/errors中New一个错误
		main.main
		        /Users/anker/kevin_go/src/GO_Learning/06_error/01_New.go:15
		runtime.main
		        /opt/homebrew/Cellar/go/1.16.5/libexec/src/runtime/proc.go:225
		runtime.goexit
		        /opt/homebrew/Cellar/go/1.16.5/libexec/src/runtime/asm_arm64.s:1130
	*/

	info := fmt.Sprint("哈哈哈哈")
	eee := errors.Errorf("试一试可以带参数的Errorf:%s", info)
	fmt.Printf("%+v\n", eee)
	/*
		试一试可以带参数的Errorf:哈哈哈哈
		main.main
		        /Users/anker/kevin_go/src/GO_Learning/06_error/01_New.go:20
		runtime.main
		        /opt/homebrew/Cellar/go/1.17.2/libexec/src/runtime/proc.go:255
		runtime.goexit
		        /opt/homebrew/Cellar/go/1.17.2/libexec/src/runtime/asm_arm64.s:1133
	*/
}
