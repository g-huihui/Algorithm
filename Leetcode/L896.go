/**
 * @Author: Gong Yanhui
 * @Description: 896. 单调数列
 * @Date: 2025-01-23 15:48
 */

package main

func isMonotonic(nums []int) bool {
	var size = len(nums)
	if size <= 2 {
		return true
	}

	var inc, dec = true, true
	for i := 1; i < size; i++ {
		if nums[i] < nums[i-1] {
			inc = false
		}
		if nums[i] > nums[i-1] {
			dec = false
		}
	}

	return inc || dec
}
