/**
 * @Author: Gong Yanhui
 * @Description: 976. 三角形的最大周长
 * @Date: 2025-02-15 16:39
 */

package main

import "sort"

func largestPerimeter(nums []int) int {
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 2; i-- { // 先确定第三条最长的边
		if nums[i-1]+nums[i-2] > nums[i] {
			return nums[i] + nums[i-1] + nums[i-2]
		}
	}
	return 0
}
