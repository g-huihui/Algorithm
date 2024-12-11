/**
 * @Author: Gong Yanhui
 * @Description: 2717. 半有序排列
 * @Date: 2024-12-11 23:52
 */

package main

func semiOrderedPermutation(nums []int) int {
	var begin, end int

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			begin = i
			continue
		}
		if nums[i] == len(nums) {
			end = i
			continue
		}
	}

	res := begin + (len(nums) - end - 1)
	if begin > end {
		res--
	}

	return res
}
