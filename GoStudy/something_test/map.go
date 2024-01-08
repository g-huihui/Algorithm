/**
 * @Author: Gong Yanhui
 * @Description: 测试map并发触发panic问题
 * @Date: 2024-01-08 14:46
 */

package something_test

import (
	"fmt"
	"time"
)

func TestConcurrentMap() {
	m := make(map[int]int)
	go func() {
		time.Sleep(1 * time.Second)
		for i := 0; i < 1000; i++ {
			m[i] = i // 写
		}
		fmt.Println("write done")
	}()
	go func() {
		time.Sleep(1 * time.Second)
		for i := 0; i < 1000; i++ {
			_ = m[i] // 读
		}
		fmt.Println("read done")
	}()
	go func() {
		time.Sleep(1 * time.Second)
		for i := 0; i < 1000; i++ {
			m[i] = i // 写
		}
		fmt.Println("write done")
	}()
	go func() {
		time.Sleep(1 * time.Second)
		for i := 0; i < 1000; i++ {
			_ = m[i] // 读
		}
		fmt.Println("read done")
	}()

	time.Sleep(10 * time.Second)
}
