/**
 * @Author: Gong Yanhui
 * @Description: 利用多核 CPU 实现并行计算 https://laravelacademy.org/post/19910
 * @Date: 2022-10-24 21:50
 */

package main

import (
	"fmt"
	"runtime"
	"time"
)

func sum(seq int, ch chan int) {
	defer close(ch)
	sum := 0
	for i := 1; i <= 10000000; i++ {
		sum += i
	}
	fmt.Printf("子协程%d运算结果:%d\n", seq, sum)
	ch <- sum
}

func main() {
	// 启动时间
	start := time.Now()
	// 最大 CPU 核心数
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus) // 设置程序运行时可以使用的最大核心数
	chs := make([]chan int, cpus)
	for i := 0; i < len(chs); i++ {
		chs[i] = make(chan int, 1)
		go sum(i, chs[i])
	}
	sum := 0
	for _, ch := range chs {
		res := <-ch
		sum += res
	}
	// 结束时间
	end := time.Now()
	// 打印耗时
	fmt.Printf("最终运算结果: %d, 执行耗时(s): %f\n", sum, end.Sub(start).Seconds())
}

/**
另外，需要注意的是，目前 Go 语言默认就是支持多核的，
所以如果上述示例代码中没有显式设置 runtime.GOMAXPROCS(cpus) 这行代码，
编译器也会利用多核 CPU 来执行代码，其结果是运行耗时和设置多核是一样的。
*/
