/**
 * @Author: Gong Yanhui
 * @Description: 2176. 统计数组中相等且可以被整除的数对
 * @Date: 2025-04-18 19:12
 */

package main

func countPairs(nums []int, k int) int {
	var size = len(nums)
	var count, mul int
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if nums[i] == nums[j] {
				mul = i * j
				if mul == 0 || mul%k == 0 {
					count++
				}
			}
		}
	}

	return count
}
