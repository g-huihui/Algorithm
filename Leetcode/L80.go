/**
 * @Author: Gong Yanhui
 * @Description: 80. 删除有序数组中的重复项 II
 * @Date: 2024-05-14 14:21
 */

package main

func removeDuplicates2(nums []int) int {
	length := len(nums)
	if length <= 2 {
		return length
	}
	var slow, fast = 2, 2
	for ; fast < length; fast++ {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
