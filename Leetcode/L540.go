/**
 * @Author: Gong Yanhui
 * @Description: 540. 有序数组中的单一元素
 * @Date: 2024-11-10 10:52
 */

package main

func singleNonDuplicate(nums []int) int {
	var res int
	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
	}

	return res
}
