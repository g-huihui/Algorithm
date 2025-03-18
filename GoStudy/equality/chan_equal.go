/**
 * @Author: Gong Yanhui
 * @Description: chan的比较
 * @Date: 2025-03-18 20:00
 */

package equality

import "fmt"

func ChanCompare() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	var ch3 chan int = ch1
	var ch4 chan int
	var ch5 chan int
	// var ch6 chan string

	// 同一个make创建的通道才相等
	// 输出：false true
	fmt.Printf("%v %v\n", ch1 == ch2, ch3 == ch1)

	// 通道值可与nil比较
	// 输出：true false
	fmt.Printf("%v %v\n", ch4 == ch5, ch5 == ch1)

	// 两个不同类型的通道，即使都是空值，也不能比较
	// invalid operation: ch5 == ch6 (mismatched types chan int and chan string)
	// fmt.Printf("%v\n", ch5 == ch6)
}
