/**
 * @Author: Gong Yanhui
 * @Description: sync.WaitGroup 类型	https://laravelacademy.org/post/20001
 * @Date: 2022-11-23 15:14
 */

package main

import (
	"fmt"
	"sync"
)

func addNum(a, b int, deferFunc func()) {
	defer deferFunc()
	c := a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)
}

func main() {
	var (
		wg    sync.WaitGroup
		count = 10
	)
	wg.Add(count)
	for i := 0; i < count; i++ {
		go addNum(i, 1, wg.Done)
	}
	wg.Wait()
	fmt.Println("All goroutine is over...")
}
