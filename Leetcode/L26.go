/**
 * @Author: Gong Yanhui
 * @Description: 26. 删除有序数组中的重复项
 * @Date: 2022-07-11 23:20
 */

package main

func removeDuplicates(nums []int) int {
	index := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}
