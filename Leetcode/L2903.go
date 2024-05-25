/**
 * @Author: Gong Yanhui
 * @Description: 2903. 找出满足差值条件的下标 I
 * @Date: 2024-05-25 11:53
 */

package main

func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	var res = []int{-1, -1}
	for i := 0; i <= len(nums)-indexDifference; i++ {
		for j := i + indexDifference; j < len(nums); j++ {
			if abs(nums[i]-nums[j]) >= valueDifference {
				return []int{i, j}
			}
		}
	}
	return res
}
