/**
 * @Author: Gong Yanhui
 * @Description: 1822. 数组元素积的符号
 * @Date: 2025-01-23 14:23
 */

package main

func arraySign(nums []int) int {
	var count int // 记录负数的个数
	for _, num := range nums {
		if num == 0 {
			return 0
		}
		if num < 0 {
			count++
		}
	}
	if count%2 == 0 {
		return 1
	} else {
		return -1
	}
}
