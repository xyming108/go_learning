package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func err00() error {
	err0 := errors.New("message0")
	return errors.WithMessage(err0, "信息0")
}

func err11() error {
	err1 := err00()
	f := fmt.Sprint("啦啦啦")
	/*WithMessage和WithMessagef不会往上抛错误，因为没有调用堆栈*/
	return errors.WithMessagef(err1, "信息1%s", f)
}

func err22() error {
	err2 := err11()
	return errors.WithMessage(err2, "信息2")
}

func main() {
	err := err22()
	if err != nil {
		fmt.Printf("%v\n\n", err)
		/*
			信息2: 信息1啦啦啦: 信息0: message0
		*/

		/*Cause 返回错误的根本原因*/
		temp := errors.Cause(err)
		fmt.Printf("%v\n\n", temp)
		/*
			message0
		*/

		/*返回上一层的错误信息*/
		u := errors.Unwrap(err)
		fmt.Printf("%v\n\n", u)
		/*
			信息1啦啦啦: 信息0: message0
		*/

		fmt.Printf("%+v\n\n", err)
		/*
			message0
			main.err00
			        /Users/anker/kevin_go/src/GO_Learning/06_error/03_WithMessage.go:9
			main.err11
			        /Users/anker/kevin_go/src/GO_Learning/06_error/03_WithMessage.go:14
			main.main
			        /Users/anker/kevin_go/src/GO_Learning/06_error/03_WithMessage.go:20
			runtime.main
			        /opt/homebrew/Cellar/go/1.17.2/libexec/src/runtime/proc.go:255
			runtime.goexit
			        /opt/homebrew/Cellar/go/1.17.2/libexec/src/runtime/asm_arm64.s:1133
			信息0
			信息1啦啦啦
			信息2
		*/

	}
}
