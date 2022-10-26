/**
 * @Author: Gong Yanhui
 * @Description: sync 包（一）：sync.Mutex 和 sync.RWMutex https://laravelacademy.org/post/19928
 * @Date: 2022-10-26 20:23
 */

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var counter_RWMutex int = 0

func add_RWMutex(a, b int, lock *sync.RWMutex) {
	c := a + b
	lock.Lock()
	counter_RWMutex++
	fmt.Printf("%d: %d + %d = %d\n", counter_RWMutex, a, b, c)
	lock.Unlock()
}

func main() {
	start := time.Now()
	lock := &sync.RWMutex{}
	for i := 0; i < 10; i++ {
		go add_RWMutex(1, i, lock)
	}

	for {
		lock.RLock()
		c := counter_RWMutex
		lock.RUnlock()
		runtime.Gosched()
		if c >= 10 {
			break
		}
	}
	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时(s)：", consume)
}

/**
不要重复锁定互斥锁；
不要忘记解锁互斥锁，必要时使用 defer 语句；
不要对尚未锁定或者已解锁的互斥锁解锁；
不要在多个函数之间直接传递互斥锁。
*/
