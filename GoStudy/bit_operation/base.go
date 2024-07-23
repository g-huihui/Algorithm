/**
 * @Author: Gong Yanhui
 * @Description: 位运算基础
 * @Date: 2024-07-23 20:27
 */

package bit_operation

func bit_operation() {

	var tag uint32 = 1 << 3 // 0001 -> 1000 = 8

	println(1 | tag) // 0001 | 1000 = 1001 = 9	放入
	println(9 | tag) // 1001 | 1000 = 1001 = 9	已经存在放入不变

	println(9 & tag) // 1001 & 1000 = 1000 = 8 检查是否存在
	println(1 & tag) // 0001 & 1000 = 0000 = 0 不存在情况下 (未完全验证)

	println(9 ^ tag) // 1001 ^ 1000 = 0001 = 1	取出
}
