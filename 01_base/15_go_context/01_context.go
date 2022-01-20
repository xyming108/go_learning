package main

import (
	"context"
	"fmt"
)

func process(ctx context.Context) {
	a, ok := ctx.Value("a").(string) //类型断言
	if ok {
		fmt.Println("get successfully:", a)
	} else {
		fmt.Println("get fail")
	}
}

func main() {
	ctx := context.Background()
	process(ctx)

	ctx = context.WithValue(ctx, "a", "我是a")
	process(ctx)
}
