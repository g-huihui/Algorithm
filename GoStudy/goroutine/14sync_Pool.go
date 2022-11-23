/**
 * @Author: Gong Yanhui
 * @Description: 临时对象池 sync.Pool	https://laravelacademy.org/post/20015
 * @Date: 2022-11-23 21:55
 */

package main

import (
	"fmt"
	"sync"
)

func testPool() {
	var pool = &sync.Pool{
		New: func() interface{} {
			return "return object When pool is empty"
		},
	}
	value := "A object"
	pool.Put(value)
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
}

func main() {
	testPool()

	var wg sync.WaitGroup
	wg.Add(1)
	var pool = &sync.Pool{
		New: func() interface{} {
			return "return object When pool is empty"
		},
	}
	go test_put(pool, wg.Done)
	wg.Wait()
	fmt.Println(pool.Get())
}

func test_put(pool *sync.Pool, deferFunc func()) {
	defer func() {
		deferFunc()
	}()
	value := "Put A Object"
	pool.Put(value)
}
