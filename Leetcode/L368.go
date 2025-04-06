/**
 * @Author: Gong Yanhui
 * @Description: 368. 最大整除子集
 * @Date: 2025-04-06 20:57
 */

package main

import "sort"

// 先看懂的300题再来做这题 和官方题解还是有区别的
func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	var size = len(nums)
	var dp = make([][]int, size)
	// 每个索引i存储的 [0, 1] 之间的最大子集
	dp[0] = append([]int{}, nums[0])
	var res = append([]int{}, nums[0])
	for i := 1; i < size; i++ {
		dp[i] = append([]int{}, nums[i])
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 {
				if len(dp[j])+1 > len(dp[i]) {
					dp[i] = append([]int{}, dp[j]...)
					dp[i] = append(dp[i], nums[i])
				}
			}
		}
		if len(dp[i]) > len(res) {
			res = dp[i]
		}
	}

	return res
}
