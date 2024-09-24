/**
 * @Author: Gong Yanhui
 * @Description: 测试往channel中放入nil 管道的读取
 * @Date: 2024-01-08 14:39
 */

package something

import (
	"fmt"
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

//
// ChannelPrint
//  @Description: 使用channel按照顺序打印dog fish cat指定次数
//  @param num 打印次数
//
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

//
// ChannelPrint2
//  @Description: 使用3个channel按照顺序打印每个channel中的数据
//
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
