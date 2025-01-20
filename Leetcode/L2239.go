/**
 * @Author: Gong Yanhui
 * @Description: 2239. 找到最接近 0 的数字
 * @Date: 2025-01-20 16:07
 */

package main

import "math"

func findClosestNumber(nums []int) int {
	var res = -math.MaxInt

	for i := 0; i < len(nums); i++ {
		if abs(nums[i]) < abs(res) || (abs(nums[i]) == abs(res) && nums[i] > res) {
			res = nums[i]
		}
	}

	return res
}
