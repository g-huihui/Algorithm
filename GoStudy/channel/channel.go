/**
 * @Author: Gong Yanhui
 * @Description: 测试往channel中放入nil 管道的读取
 * @Date: 2024-01-08 14:39
 */

package channel

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func TestChannel() {
	type Test struct {
		A int
	}

	cha := make(chan *Test, 3)

	cha <- &Test{1}
	cha <- nil
	cha <- &Test{2}
	close(cha)

	a := <-cha
	if a == nil {
		fmt.Println("nil")
	} else {
		fmt.Println(a)
	}
	// Print: &{1}

	b := <-cha
	if b == nil {
		fmt.Println("nil")
	} else {
		fmt.Println(b)
	}
	// Print: nil

	c := <-cha
	if c == nil {
		fmt.Println("nil")
	} else {
		fmt.Println(c)
	}
	// Print: &{2}

	d := <-cha
	if d == nil {
		fmt.Println("nil")
	} else {
		fmt.Println(d)
	}
	// Print: nil
}

// ChannelPrint
//
//	@Description: 使用channel按照顺序打印dog fish cat指定次数
//	@param num 打印次数
func ChannelPrint(num int) {

	var fun = func(c chan<- string, str string) {
		for i := 0; i < num; i++ {
			c <- str
		}
	}

	var (
		dogChan  = make(chan string)
		fishChan = make(chan string)
		catChan  = make(chan string)
	)
	go fun(dogChan, "dog")
	go fun(fishChan, "fish")
	go fun(catChan, "cat")

	for i := 0; i < num; i++ {
		var str string
		str = <-dogChan
		println(str)
		str = <-fishChan
		println(str)
		str = <-catChan
		println(str)
	}
}

// ChannelPrint2
//
//	@Description: 使用3个channel按照顺序打印每个channel中的数据
func ChannelPrint2() {

	var flag int32 = 0

	var send = func(ch chan<- int, begin int) {
		for num := begin; ; num += 3 {
			ch <- num
			time.Sleep(time.Second)
		}
	}

	var printNum = func(ch <-chan int, flagNum int32) {
		for {
			if atomic.LoadInt32(&flag) == flagNum {
				println(<-ch)
				atomic.StoreInt32(&flag, (flagNum+1)%3)
			}
		}
	}

	var ch1 = make(chan int)
	var ch2 = make(chan int)
	var ch3 = make(chan int)

	go send(ch1, 0)
	go send(ch2, 1)
	go send(ch3, 2)

	go printNum(ch1, 0)
	go printNum(ch2, 1)
	go printNum(ch3, 2)

	select {}
}

// ChannelPrint3
//
//	@Description: 三个协程 1 打印a 2 打印b 3 打印c 输出字符串abc的数量
//
// @param n 输出字符串abc的数量
func ChannelPrint3(n int) {
	// 输入[1] 输出[abc]
	// 输入[3] 输出[abcabcabc]

	// 三个channel 分别存放a b c
	var ch1 = make(chan string)
	var ch2 = make(chan string)
	var ch3 = make(chan string)

	// 标记tag 0 1 2
	var flag int32 = 0

	// 全局等待
	var w = sync.WaitGroup{}
	w.Add(3) // 3个读取协程

	// 发送函数 往channel无限发送
	var sendFunc = func(ch chan<- string, str string) {
		for {
			ch <- str
		}
	}

	go sendFunc(ch1, "a")
	go sendFunc(ch2, "b")
	go sendFunc(ch3, "c")

	var printFunc = func(ch <-chan string, flagNum *int32, n int, w *sync.WaitGroup, tag int32) {
		for {
			select {
			case str := <-ch:
				if atomic.LoadInt32(flagNum) == tag {
					fmt.Print(str)
					atomic.StoreInt32(flagNum, (tag+1)%3)
					n--
					if n == 0 {
						w.Done()
						return
					}
				}
			}
		}
	}

	go printFunc(ch1, &flag, n, &w, 0)
	go printFunc(ch2, &flag, n, &w, 1)
	go printFunc(ch3, &flag, n, &w, 2)

	w.Wait()
}

// ChannelPrint4
//
//	@Description: N个协程 按序分别打印各自的序号
//
// @param n 协程数量
// @param count 打印次数
func ChannelPrint4(n int, count int) {

	// 创建n个channel 并且往每个channel中都填满数据
	var chs = make([]chan int, n)
	for i := 0; i < n; i++ {
		chs[i] = make(chan int)
		go func(ch chan int, num int) {
			for {
				ch <- num
			}
		}(chs[i], i+1)
	}

	// 标记该哪个协程打印了
	var flag int64

	var w = sync.WaitGroup{}
	w.Add(n)
	// 定义打印协程的函数
	var printFunc = func(ch <-chan int, tag int64) {
		var c int
		for num := range ch {
			if atomic.LoadInt64(&flag) == tag {
				print(num)
				atomic.StoreInt64(&flag, (tag+1)%int64(n))
				c++
				if c == count {
					w.Done()
					return
				}
			}
		}
	}

	// 开启n个协程
	for i := 0; i < n; i++ {
		go printFunc(chs[i], int64(i))
	}

	w.Wait()
}
