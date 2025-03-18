/**
 * @Author: Gong Yanhui
 * @Description: 数组比较
 * @Date: 2025-03-18 20:24
 */

package equality

import "fmt"

/**
如果数组中的元素类型是可比的 则数组也是可比较的
如果数组中对应的元素都相等 那么两个数组是相等的
数组不能与nil比较
*/

func ArrayCompare() {
	a1 := [3]int{1, 2, 3}
	a2 := [3]int{1, 2, 3}
	a3 := [3]int{2, 1, 3}

	// 元素顺序必须一样
	// 输出：true false
	fmt.Printf("%v %v\n", a1 == a2, a2 == a3)

	// 无法比较
	// fmt.Printf("%v\n", a3 == nil)
}
