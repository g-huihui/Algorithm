/**
 * @Author: Gong Yanhui
 * @Description: 3255. 长度为 K 的子数组的能量值 II
 * @Date: 2024-11-08 10:37
 */

package main

import "fmt"

// nums = [1,2,3,4,3,2,5], k = 3  => [3,4,-1,-1,-1]
// 用一个数组记录前面连续的长度 [1,2,3,4,1,1,1]

func resultsArray2(nums []int, k int) []int {
	if k == 1 {
		return nums
	}
	var dp = make([]int, len(nums))
	dp[0] = 1
	var res = make([]int, 0, len(nums)-k+1)
	for i := 1; i < len(nums); i++ {
		if nums[i]-1 == nums[i-1] {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = 1
		}
		if i >= k-1 {
			if dp[i] >= k {
				res = append(res, nums[i])
			} else {
				res = append(res, -1)
			}
		}
	}
	return res
}

func main() {
	fmt.Println(resultsArray2([]int{1, 2, 3, 4, 3, 2, 5}, 3)) // [3,4,-1,-1,-1]
	fmt.Println(resultsArray2([]int{1}, 1))                   // [1]
	fmt.Println(resultsArray2([]int{2}, 1))                   // [2]
	fmt.Println(resultsArray2([]int{1, 1}, 1))                // [1,1]
}
