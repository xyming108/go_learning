package _1_context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

const shortDuration = 1 * time.Millisecond

func TestContextValue(t *testing.T) {
	ctx := context.WithValue(context.Background(), "key1", "val1")
	value := ctx.Value("key1")
	fmt.Println(value)

	subCtx := context.WithValue(ctx, "", "")
	subValue := subCtx.Value("key1")
	fmt.Println(subValue)
}

func TestControl(t *testing.T) {
	ctx1, cancelFunc1 := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc1()
	ctx2, cancelFunc2 := context.WithTimeout(ctx1, 3*time.Second)
	defer cancelFunc2()
	go func() {
		t1 := time.Now().Second()
		<-ctx2.Done()
		fmt.Println("timeout")
		fmt.Println(time.Now().Second()-t1)
	}()
	time.Sleep(2 * time.Second)
}

func ExampleWithCancel() {
	// gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal goroutine started by gen.
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
}

func ExampleWithDeadline() {
	d := time.Now().Add(100 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is good practice to call its
	// cancellation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

	// Output:
	// context deadline exceeded
}

func ExampleWithTimeout() {
	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	ctx, cancel := context.WithTimeout(context.Background(), shortDuration)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}

	// Output:
	// context deadline exceeded
}
