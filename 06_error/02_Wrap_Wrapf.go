package main

import (
	"fmt"
	"github.com/pkg/errors"
)

/*
Wrap returns an error annotating err with a stack trace
at the point Wrap is called, and the supplied message.
If err is nil, Wrap returns nil.

Wrap 返回一个错误注释错误，并在调用 Wrap 时使用堆栈跟踪和提供的消息。如果 err 为 nil，则 Wrap 返回 nil

选择 wrap error 是只有 applications 可以选择应用的策略。具有最高可重用性的包只能返回根错误值。此机制与 Go 标准库中使用的相同(kit 库的 sql.ErrNoRows)
*/

func err0() error {
	err0 := errors.New("err0")
	return errors.Wrap(err0, "错误0")
}

func err1() error {
	err1 := err0()
	return errors.Wrap(err1, "错误1")
}

func err2() error {
	err2 := err1()
	f := []int{1, 2, 3}
	/*相比于Wrap，可带参数*/
	return errors.Wrapf(err2, "错误2%d", f)
}

func main() {
	err := err2()
	if err != nil {
		fmt.Printf("%v\n\n", err)
		/*
			错误2[1 2 3]: 错误1: 错误0: err0
		*/

		/*Cause 返回错误的根本原因*/
		temp := errors.Cause(err)
		fmt.Printf("%v\n\n", temp)
		/*
			err0
		*/

		//打印出堆栈信息
		fmt.Printf("%+v\n\n", err)
		/*
			err0
			main.err0
			        /Users/anker/kevin_go/src/GO_Learning/06_error/02_Wrap_Wrapf.go:17
			main.err1
			        /Users/anker/kevin_go/src/GO_Learning/06_error/02_Wrap_Wrapf.go:22
			main.err2
			        /Users/anker/kevin_go/src/GO_Learning/06_error/02_Wrap_Wrapf.go:27
			main.main
			        /Users/anker/kevin_go/src/GO_Learning/06_error/02_Wrap_Wrapf.go:34
			runtime.main
			        /opt/homebrew/Cellar/go/1.17.2/libexec/src/runtime/proc.go:255
			runtime.goexit
			        /opt/homebrew/Cellar/go/1.17.2/libexec/src/runtime/asm_arm64.s:1133
			错误0
			main.err0
			        /Users/anker/kevin_go/src/GO_Learning/06_error/02_Wrap_Wrapf.go:18
			main.err1
			        /Users/anker/kevin_go/src/GO_Learning/06_error/02_Wrap_Wrapf.go:22
			main.err2
			        /Users/anker/kevin_go/src/GO_Learning/06_error/02_Wrap_Wrapf.go:27
			main.main
			        /Users/anker/kevin_go/src/GO_Learning/06_error/02_Wrap_Wrapf.go:34
			runtime.main
			        /opt/homebrew/Cellar/go/1.17.2/libexec/src/runtime/proc.go:255
			runtime.goexit
			        /opt/homebrew/Cellar/go/1.17.2/libexec/src/runtime/asm_arm64.s:1133
			错误1
			main.err1
			        /Users/anker/kevin_go/src/GO_Learning/06_error/02_Wrap_Wrapf.go:23
			main.err2
			        /Users/anker/kevin_go/src/GO_Learning/06_error/02_Wrap_Wrapf.go:27
			main.main
			        /Users/anker/kevin_go/src/GO_Learning/06_error/02_Wrap_Wrapf.go:34
			runtime.main
			        /opt/homebrew/Cellar/go/1.17.2/libexec/src/runtime/proc.go:255
			runtime.goexit
			        /opt/homebrew/Cellar/go/1.17.2/libexec/src/runtime/asm_arm64.s:1133
			错误2[1 2 3]
			main.err2
			        /Users/anker/kevin_go/src/GO_Learning/06_error/02_Wrap_Wrapf.go:30
			main.main
			        /Users/anker/kevin_go/src/GO_Learning/06_error/02_Wrap_Wrapf.go:34
			runtime.main
			        /opt/homebrew/Cellar/go/1.17.2/libexec/src/runtime/proc.go:255
			runtime.goexit
			        /opt/homebrew/Cellar/go/1.17.2/libexec/src/runtime/asm_arm64.s:1133
		*/
	}
}
