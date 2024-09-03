/**
 * @Author: Gong Yanhui
 * @Description: 2708. 一个小组的最大实力值
 * @Date: 2024-09-03 20:54
 */

package main

import "sort"

func maxStrength(nums []int) int64 {
	if len(nums) == 1 {
		return int64(nums[0])
	}

	sort.Ints(nums)

	var res int64 = 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		if int64(nums[i]) > 0 {
			if res == 0 {
				res = 1
			}
			res = res * int64(nums[i])
			continue
		}
		// 负数的情况下
		if i < len(nums)-1 && nums[i+1] < 0 {
			if res == 0 {
				res = 1
			}
			res = res * int64(nums[i]) * int64(nums[i+1])
			i++
		}
	}

	return res
}

func main() {
	println(maxStrength([]int{3, -1, -5, 2, 5, -9})) // 1350
	println(maxStrength([]int{-4, -5, -4}))          // 20
	println(maxStrength([]int{0, -1}))               // 0
}
