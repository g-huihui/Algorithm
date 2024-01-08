/**
 * @Author: Gong Yanhui
 * @Description: 测试往channel中放入nil 管道的读取
 * @Date: 2024-01-08 14:39
 */

package something_test

import "fmt"

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
