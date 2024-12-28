/**
 * @Author: Gong Yanhui
 * @Description: 3046. 分割数组
 * @Date: 2024-12-28 08:56
 */

package main

func isPossibleToSplit(nums []int) bool {
	var m = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
		if m[nums[i]] > 2 {
			return false
		}
	}
	return true
}
