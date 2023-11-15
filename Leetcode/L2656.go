/**
 * @Author: Gong Yanhui
 * @Description: 2656. K 个元素的最大和
 * @Date: 2023-11-15 14:37
 */

package main

import "sort"

func maximizeSum(nums []int, k int) int {
	sort.Ints(nums)
	total := nums[len(nums)-1] * k
	for i := 0; i < k; i++ {
		total += i
	}
	return total
}
