/**
 * @Author: Gong Yanhui
 * @Description: WithTimeout 超时控制
 * @Date: 2023-08-05 15:59
 */

package main

import (
	"context"
	"fmt"
	"time"
)

func NewContextWithTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 3*time.Second)
}

func HttpHandler1() {
	ctx, cancel := NewContextWithTimeout()
	defer cancel()

	for i := 0; ; i++ {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Printf("deal time is %d\n", i)
		}
	}
}

func HttpHandler2() {
	ctx, cancel := NewContextWithTimeout()
	defer cancel()

	for i := 0; ; i++ {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Printf("deal time is %d\n", i)
			cancel()
		}
	}
}

func main() {
	HttpHandler1() //    达到超时时间终止接下来的执行
	HttpHandler2() // 没有达到超时时间终止接下来的执行
}
