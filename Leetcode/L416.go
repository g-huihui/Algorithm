/**
 * @Author: Gong Yanhui
 * @Description: 416. 分割等和子集
 * @Date: 2025-04-20 20:20
 */

package main

func canPartition(nums []int) bool {
	var size = len(nums)
	if size == 0 {
		return false
	}

	var total int
	for i := 0; i < size; i++ {
		total += nums[i]
	}
	if total%2 == 1 {
		return false
	}

	var result = total / 2

	return f416(nums, result, size)
}

// 0 1 背包解法
func f416(nums []int, result, size int) bool {
	var dp = make([][]int, size+1)
	for i := range dp {
		dp[i] = make([]int, result+1)
	}

	for i := 1; i <= size; i++ {
		for j := 1; j <= result; j++ {
			var num int
			if j-nums[i-1] >= 0 {
				num = dp[i-1][j-nums[i-1]] + nums[i-1]
			} else {
				num = nums[i-1]
			}
			dp[i][j] = max(dp[i-1][j], num)
			if dp[i][j] == result {
				return true
			}
		}
	}
	return false
}

func main() {
	println(canPartition([]int{1, 5, 11, 5})) // true
	println(canPartition([]int{1, 2, 3, 5}))  // false

	println(canPartition([]int{1, 2, 5}))                        // false
	println(canPartition([]int{3, 3, 6, 8, 16, 16, 16, 18, 20})) // true 未通过
}
