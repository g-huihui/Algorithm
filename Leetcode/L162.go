/**
 * @Author: Gong Yanhui
 * @Description: 162. 寻找峰值
 * @Date: 2024-12-22 21:41
 */

package main

func findPeakElement(nums []int) int {
	var begin, end = 0, len(nums) - 1
	for begin < end {
		mid := (begin + end) / 2
		if nums[mid] < nums[mid+1] {
			begin = mid + 1
		} else {
			end = mid
		}
	}

	return begin
}
