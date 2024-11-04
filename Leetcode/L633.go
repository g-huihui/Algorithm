/**
 * @Author: Gong Yanhui
 * @Description: 633. 平方数之和
 * @Date: 2024-11-04 20:21
 */

package main

import "math"

func judgeSquareSum(c int) bool {
	var left, right = 0, int(math.Sqrt(float64(c)))
	for left <= right {
		sum := left*left + right*right
		if sum == c {
			return true
		} else if sum > c {
			right--
		} else {
			left++
		}
	}

	return false
}
