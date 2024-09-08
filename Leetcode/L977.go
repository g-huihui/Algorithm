/**
 * @Author: Gong Yanhui
 * @Description: 977. 有序数组的平方
 * @Date: 2024-09-08 22:12
 */

package main

import "sort"

func sortedSquares(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] * nums[i]
	}
	sort.Ints(nums)

	return nums
}
