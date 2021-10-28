package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func err000() error {
	err0 := errors.New("err0")
	return errors.WithStack(err0)
}

func err111() error {
	err1 := err000()
	return errors.WithStack(err1)
}

func main() {
	err := err111()
	if err != nil {
		fmt.Printf("%v\n\n", err)
		/*
			err0
		*/

		/*回退一层堆栈错误*/
		fmt.Printf("%+v\n\n", errors.Unwrap(err))
		/*
		    main.err000
			        /Users/anker/kevin_go/src/GO_Learning/06_error/04_WithStack.go:9
			main.err111
			        /Users/anker/kevin_go/src/GO_Learning/06_error/04_WithStack.go:14
			main.main
			        /Users/anker/kevin_go/src/GO_Learning/06_error/04_WithStack.go:19
			runtime.main
			        /opt/homebrew/Cellar/go/1.17.2/libexec/src/runtime/proc.go:255
			runtime.goexit
			        /opt/homebrew/Cellar/go/1.17.2/libexec/src/runtime/asm_arm64.s:1133
			main.err000
			        /Users/anker/kevin_go/src/GO_Learning/06_error/04_WithStack.go:10
			main.err111
			        /Users/anker/kevin_go/src/GO_Learning/06_error/04_WithStack.go:14
			main.main
			        /Users/anker/kevin_go/src/GO_Learning/06_error/04_WithStack.go:19
			runtime.main
			        /opt/homebrew/Cellar/go/1.17.2/libexec/src/runtime/proc.go:255
			runtime.goexit
			        /opt/homebrew/Cellar/go/1.17.2/libexec/src/runtime/asm_arm64.s:1133

		*/

		/*直达根部错误*/
		fmt.Printf("%+v\n\n", errors.Cause(err))
		/*
				err0
			main.err000
			        /Users/anker/kevin_go/src/GO_Learning/06_error/04_WithStack.go:9
			main.err111
			        /Users/anker/kevin_go/src/GO_Learning/06_error/04_WithStack.go:14
			main.main
			        /Users/anker/kevin_go/src/GO_Learning/06_error/04_WithStack.go:19
			runtime.main
			        /opt/homebrew/Cellar/go/1.17.2/libexec/src/runtime/proc.go:255
			runtime.goexit
			        /opt/homebrew/Cellar/go/1.17.2/libexec/src/runtime/asm_arm64.s:1133

		*/

		/*所有错误的堆栈信息*/
		fmt.Printf("%+v\n\n", err)
		/*
			err0
			main.err000
			        /Users/anker/kevin_go/src/GO_Learning/06_error/04_WithStack.go:9
			main.err111
			        /Users/anker/kevin_go/src/GO_Learning/06_error/04_WithStack.go:14
			main.main
			        /Users/anker/kevin_go/src/GO_Learning/06_error/04_WithStack.go:19
			runtime.main
			        /opt/homebrew/Cellar/go/1.17.2/libexec/src/runtime/proc.go:255
			runtime.goexit
			        /opt/homebrew/Cellar/go/1.17.2/libexec/src/runtime/asm_arm64.s:1133
			main.err000
			        /Users/anker/kevin_go/src/GO_Learning/06_error/04_WithStack.go:10
			main.err111
			        /Users/anker/kevin_go/src/GO_Learning/06_error/04_WithStack.go:14
			main.main
			        /Users/anker/kevin_go/src/GO_Learning/06_error/04_WithStack.go:19
			runtime.main
			        /opt/homebrew/Cellar/go/1.17.2/libexec/src/runtime/proc.go:255
			runtime.goexit
			        /opt/homebrew/Cellar/go/1.17.2/libexec/src/runtime/asm_arm64.s:1133
			main.err111
			        /Users/anker/kevin_go/src/GO_Learning/06_error/04_WithStack.go:15
			main.main
			        /Users/anker/kevin_go/src/GO_Learning/06_error/04_WithStack.go:19
			runtime.main
			        /opt/homebrew/Cellar/go/1.17.2/libexec/src/runtime/proc.go:255
			runtime.goexit
			        /opt/homebrew/Cellar/go/1.17.2/libexec/src/runtime/asm_arm64.s:1133
		*/
	}
}
