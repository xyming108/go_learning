package main

import (
	"errors"
	"fmt"
)

type ErrorString struct {
	s string
}

func (e *ErrorString) Error() string {
	return e.s
}

func main() {
	/*
		errors.Is
		当多层调用返回的错误被一次次地包装起来，我们在调用链上游拿到的错误如何判断是否是底层的某个错误呢？
		它递归调用 Unwrap 并判断每一层的 err 是否相等，如果有任何一层 err 和传入的目标错误相等，则返回 true。
	*/
	var err01 = errors.New("not found")
	var err02 = fmt.Errorf("err02 and err01:%w", err01)
	if errors.Is(err02, err01) {
		fmt.Println("对上啦")
		/*
			对上啦
		*/
	}

	fmt.Println(err02)
	fmt.Println(errors.Unwrap(err02))
	fmt.Println(errors.Unwrap(errors.Unwrap(err02)))
	/*
		err02 and err01:not found
		not found
		<nil>

	*/

	/*
		errors.As
		这个和上面的 errors.Is 大体上是一样的，区别在于 Is 是严格判断相等，即两个 error 是否相等。
		而 As 则是判断类型是否相同，并提取第一个符合目标类型的错误，用来统一处理某一类错误。
	*/
	var targetErr *ErrorString
	err03 := fmt.Errorf("err03 and target err:%w", &ErrorString{s: "target err"})
	if errors.As(err03, &targetErr) {
		fmt.Println("是这个类型")
		/*
			是这个类型
		*/
	}
}
