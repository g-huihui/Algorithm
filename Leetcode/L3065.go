/**
 * @Author: Gong Yanhui
 * @Description: 3065. 超过阈值的最少操作数 I
 * @Date: 2025-01-14 21:57
 */

package main

func minOperations(nums []int, k int) int {
	var res int
	for i := 0; i < len(nums); i++ {
		if nums[i] < k {
			res++
		}
	}

	return res
}
