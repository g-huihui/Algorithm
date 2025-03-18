/**
 * @Author: Gong Yanhui
 * @Description: 结构体比较
 * @Date: 2025-03-18 20:19
 */

package equality

import "fmt"

/**
不同struct不可比较
相同struct 如果struct中所有的字段都是可比较的 那么两个struct是可比较的
相同struct 如果struct对应的非空白字段相等 则它们相等
struct不能与nil比较
*/

func StructCompare() {
	type person struct {
		name string
		age  int
	}

	p1 := person{
		name: "luna",
	}
	p2 := person{"ele", 0}
	p3 := person{"luna", 0}

	//输出：false true
	fmt.Printf("%v %v\n", p1 == p2, p1 == p3)

	var p4 person
	var p5 person

	//输出：true
	fmt.Printf("%v\n", p4 == p5)

	type person2 struct {
		age int
		ch  chan string
	}

	p6 := person2{age: 1, ch: make(chan string)}
	p7 := person2{age: 1, ch: make(chan string)}
	p8 := person2{age: 1, ch: make(chan string)}

	//输出：false false
	fmt.Printf("%v %v\n", p6 == p7, p6 == p8)

	//type person3 struct {
	//	name string
	//	age  int
	//}
	//
	//p9 := person3{"ele", 0}
	// 不同struct不可比较
	//fmt.Printf("%v\n", p2 == p9)
}
