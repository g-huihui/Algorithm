/**
 * @Author: Gong Yanhui
 * @Description: 接口的比较
 * @Date: 2025-03-18 20:08
 */

package equality

import "fmt"

/**
接口值是一个两个字长度的 数据结构
1. 第一个字包含一个指向内部表的指针 这个内部表叫作iTable 包含了已存储的值的类型信息（动态类型）以及与这个值相关联的一组方法
2. 第二个字是一个指向所存储值（动态值）的指针

结论: 如果两个接口值具有相同的动态类型和相等的动态值 则它们相等
*/

type Speaker interface {
	Speak()
}

type Person struct {
	name string
}

func (p Person) Speak() {
	fmt.Println(p.name)
}

type Student struct {
	name string
}

func (s Student) Speak() {
	fmt.Println(s.name)
}

func InterfaceCompare1() {

	p1 := Person{"ball"}
	p2 := Person{"ball"}
	p3 := Person{"luna"}

	s1 := Student{"ball"}

	// true false false false
	fmt.Printf("%v %v %v %v\n",
		Speaker(p1) == p2,
		Speaker(p1) == Speaker(p3),
		Speaker(p1) == Speaker(s1),
		Speaker(s1) == nil,
	)
}

/**
接口与非接口比较
*/

func InterfaceCompare2() {

	p1 := Person{"ball"}
	p2 := Person{"ball"}
	p3 := Person{"luna"}

	s1 := Student{"ball"}

	//输出：true false false
	fmt.Printf("%v %v %v\n",
		Speaker(p1) == p2,
		Speaker(p1) == p3,
		Speaker(p1) == s1,
	)
}
