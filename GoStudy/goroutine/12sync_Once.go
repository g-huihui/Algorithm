/**
 * @Author: Gong Yanhui
 * @Description: sync.Once 类型	https://laravelacademy.org/post/20001
 * @Date: 2022-11-23 15:50
 */

package main

import (
	"fmt"
	"sync"
	"time"
)

// 只提供了一个 Do 方法 该方法只接受一个参数 且这个参数的类型必须是 func() 即无参数无返回值的函数类型
// 可以保证该方法只会执行一次

func doSomething(o *sync.Once) {
	fmt.Println("Start:")
	o.Do(func() {
		fmt.Println("Do Something...")
	})
	fmt.Println("Finished.")
}

func main() {
	var (
		o     = &sync.Once{}
		count = 3
	)
	for i := 0; i < count; i++ {
		go doSomething(o)
	}
	time.Sleep(time.Second * time.Duration(count))
}
