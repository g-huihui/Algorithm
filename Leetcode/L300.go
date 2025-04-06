/**
 * @Author: Gong Yanhui
 * @Description: 300. 最长递增子序列
 * @Date: 2025-04-06 20:42
 */

package main

// 详细见题解 看懂了才做的
func lengthOfLIS(nums []int) int {
	var size = len(nums)
	var dp = make([]int, size)
	var res = 1
	dp[0] = 1
	for i := 1; i < size; i++ {
		var m = 1 // 当前区间的最大长度 [j, i]
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				m = max(m, dp[j]+1)
			}
		}
		dp[i] = m
		res = max(res, m)
	}

	return res
}
