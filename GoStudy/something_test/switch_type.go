/**
 * @Author: Gong Yanhui
 * @Description: 使用switch方式获取结构类型
 * @Date: 2024-09-24 10:54
 */

package something

import "fmt"

type Test struct {
	A string
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	case bool:
		fmt.Printf("the %v type is bool\n", v)
	case Test:
		fmt.Printf("the %v type is %T\n", v.A, v)
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func TestSwitchType() {
	do(21)
	do("hello")
	do(true)
	do(Test{"test"})
	do(3.14) // I don't know about type float64!
}
