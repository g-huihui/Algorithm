/**
 * @Author: Gong Yanhui
 * @Description: 通过共享内存来实现协程间的消息传递 https://laravelacademy.org/post/19888
 * @Date: 2022-10-15 15:32
 */

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 共享变量
var counter int = 0

func add(a, b int, lock *sync.Mutex) {
	c := a + b
	lock.Lock()
	counter++
	fmt.Printf("%d: %d + %d = %d\n", counter, a, b, c)
	lock.Unlock()
}

func main() {
	var (
		start = time.Now()
		lock  = sync.Mutex{}
	)
	for i := 0; i < 10; i++ {
		go add(1, i, &lock)
	}

	for {
		lock.Lock()
		c := counter
		lock.Unlock()
		runtime.Gosched()
		if c >= 10 {
			break
		}
	}

	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时(s):", consume)
}
