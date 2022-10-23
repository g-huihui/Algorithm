/**
 * @Author: Gong Yanhui
 * @Description: 使用缓冲通道提升性能 https://laravelacademy.org/post/19896
 * @Date: 2022-10-23 14:49
 */

package main

import (
	"fmt"
	"time"
)

func test(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) // 只能关闭一次通道并且由发送端关闭通道(避免两种可能的panic)
}

func main() {
	start := time.Now()
	ch := make(chan int, 2)
	go test(ch)
	for i := range ch {
		fmt.Println("receive message : ", i)
	}
	end := time.Now()

	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时(s):", consume)
}
