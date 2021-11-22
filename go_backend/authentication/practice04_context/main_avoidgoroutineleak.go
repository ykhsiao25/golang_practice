package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//這邊的 return value 直接用chan即可, 不是int
	for i := range gen(ctx) {
		fmt.Println(i)
		if i == 5 {
			cancel()
			break
		}
	}
	time.Sleep(1 * time.Second)
}

func gen(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("return")
				return
			case ch <- n:
				n++
			}
		}
	}()
	return ch
}
