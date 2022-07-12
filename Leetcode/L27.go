/**
 * @Author: Gong Yanhui
 * @Description: 27. 移除元素
 * @Date: 2022-07-12 23:59
 */

package main

func removeElement(nums []int, val int) int {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}
