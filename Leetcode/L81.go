/**
 * @Author: Gong Yanhui
 * @Description: 81. 搜索旋转排序数组 II
 * @Date: 2025-02-02 16:24
 */

package main

func search(nums []int, target int) bool {
	var size = len(nums)
	if size == 0 {
		return false
	}
	if size == 1 {
		return nums[0] == target
	}
	var left, right = 0, size - 1
	for left <= right {
		// mid = (left + right) / 2
		var mid = left + (right-left)/2 // 防止int溢出
		if nums[mid] == target {
			return true
		}
		if nums[left] == nums[mid] && nums[mid] == nums[right] {
			left++
			right--
		} else if nums[left] <= nums[mid] { // 左边有序 只看有序的一边来反推无序那边
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // 右边有序
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return false
}
