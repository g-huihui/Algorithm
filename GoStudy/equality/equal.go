/**
 * @Author: Gong Yanhui
 * @Description: 记录一些不可比较的类型
 * @Date: 2025-03-18 20:17
 */

package equality

import (
	"bytes"
	"fmt"
	"reflect"
)

// 可比较的简单类型
// Boolean(布尔值)、Integer(整型)、Floating-point(浮点数)、Complex(复数)、String(字符)

// 以下为不可比较的类型
// map
// slice
// func
// 特殊情况 当他们值是 nil 时 可以与nil比较

//func MapCompare() {
//	m1 = make(map[string]int)
//	m2 = make(map[string]int)
//	m3 = make(map[string]string)
//
//	fmt.Printf("%v %v\n", m1 == m2, m1 == m3)
//}

//func SliceCompare() {
//	s1 = make([]int, 0, 1)
//	s2 = make([]int, 0, 1)
//
//	fmt.Printf("%v\n", s1 == s2)
//}

func DynamicTyping() {
	// 静态类型 由声明所定义
	var x int
	_ = x

	// 动态类型 运行时存储在变量中的值的实际类型 动态类型在执行过程中可能会有所不同
	var someVariable interface{} = 101 // int
	fmt.Println(someVariable)
	someVariable = "str" // string
	fmt.Println(someVariable)
}

type Data struct {
	UUID    string
	Content interface{}
}

func Compare() {
	var x, y Data
	x = Data{
		UUID:    "856f5555806443e98b7ed04c5a9d6a9a",
		Content: 1,
	}
	y = Data{
		UUID:    "745dee7719304991862e6985ea9c02a9",
		Content: 2,
	}
	// string 类型和 interface 是可比较类型 那么两个 Data 类型的数据 我们可以通过 == 操作符进行比
	fmt.Println(x == y) // false

	// 如果在 Content 中的动态类型是 map 会怎样 ?
	var m, n Data
	m = Data{
		UUID:    "9584dba3fe26418d86252d71a5d78049",
		Content: map[int]string{1: "GO", 2: "Python"},
	}
	n = Data{
		UUID:    "9584dba3fe26418d86252d71a5d78049",
		Content: map[int]string{1: "GO", 2: "Python"},
	}
	// 我们程序编译通过 但会发生运行时错误
	// fmt.Println(m == n) // panic: runtime error: comparing uncomparable type map[int]string

	// 那针对这种需求 即对于不可比较类型 因为不能使用比较操作符 == 但我们想要比较它们包含的值是否相等时 应该怎么办 ?

	// 我们可以采用 reflect.DeepEqual() 函数进行比较
	fmt.Println(reflect.DeepEqual(m, n)) // true

	// 还有一种类型 []byte
	b1 := []byte("123")
	b2 := []byte("123")
	// fmt.Printf("%v\n", b1 == b2) // 无法使用这种方式比较
	fmt.Printf("%v\n", bytes.Equal(b1, b2)) // true

	// 相关文章 https://www.51cto.com/article/690436.html
}
