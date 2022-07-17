/**
 * @Author: Gong Yanhui
 * @Description: 69. x 的平方根
 * @Date: 2022-07-17 22:49
 */

package main

func mySqrt(x int) int {
	// 二分查找
	left, right := 0, x
	res := -1
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid <= x {
			res = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return res
}
