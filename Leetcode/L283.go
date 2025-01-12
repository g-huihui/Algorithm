/**
 * @Author: Gong Yanhui
 * @Description: 283. 移动零
 * @Date: 2025-01-12 11:50
 */

package main

func moveZeroes(nums []int) {
	var left, right, size = 0, 0, len(nums)
	for right < size {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

// [0, 1, 0, 3, 12]
// [1, 0, 0, 3, 12]
// [1, 3, 0, 0, 12]
// [1, 3, 12, 0, 0]
