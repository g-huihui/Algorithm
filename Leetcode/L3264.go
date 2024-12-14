/**
 * @Author: Gong Yanhui
 * @Description: 3264. K 次乘运算后的最终数组 I
 * @Date: 2024-12-14 12:23
 */

package main

func getFinalState(nums []int, k int, multiplier int) []int {

	// 循环k次即可
	for i := 0; i < k; i++ {
		// 找到最小值的索引下标
		var index = 0
		for j := 0; j < len(nums); j++ {
			if nums[j] < nums[index] {
				index = j
			}
		}
		nums[index] = nums[index] * multiplier
	}

	return nums
}
