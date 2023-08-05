/**
 * @Author: Gong Yanhui
 * @Description: WithCancel 取消控制
 * @Date: 2023-08-05 16:09
 */

package main

import (
	"context"
	"time"
)

func Speak(ctx context.Context) {
	for range time.Tick(time.Second) {
		select {
		case <-ctx.Done():
			println("speak done")
			return
		default:
			println("speak")
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go Speak(ctx)
	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
}
