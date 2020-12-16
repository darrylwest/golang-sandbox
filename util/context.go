package main

import (
	"context"
	"fmt"
)

const ctxkey = "MyKey"

type Params struct {
	appkey string
}

func gen(ctx context.Context) <-chan int {
	fmt.Printf("app key params %v\n", ctx.Value(ctxkey))

	ch := make(chan int)
	go func() {
		var n int
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- n:
				n++
			}
		}
	}()
	return ch
}

func main() {
	params := Params{"12345"}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctxv := context.WithValue(ctx, ctxkey, params)

	for n := range gen(ctxv) {
		fmt.Println(n)
		if n == 5 {
			cancel()
			break
		}
	}
}
