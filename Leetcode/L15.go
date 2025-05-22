/**
 * @Author: Gong Yanhui
 * @Description: 15. 三数之和
 * @Date: 2025-05-22 12:39
 */

package main

import "sort"

// 双指针解法
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var size = len(nums)

	var res [][]int
	for i := 0; i < size-2; i++ {
		if i > 0 && nums[i] == nums[i-1] { // 去重 跳过相同的 nums[i]
			continue
		}
		var left, right = i + 1, size - 1
		for left < right {
			var total = nums[i] + nums[left] + nums[right]
			if total == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				// 去重 跳过相同的 left 和 right
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if total < 0 {
				left++
			} else { // total > 0
				right--
			}
		}
	}

	return res
}

// 使用回溯算法暴力破解 超时
func threeSum1(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	backend15(0, 0, nums, []int{}, &res)

	return res
}

func backend15(index, total int, nums, curArr []int, res *[][]int) {
	if len(curArr) > 3 {
		return
	}

	if total == 0 && len(curArr) == 3 {
		*res = append(*res, append([]int{}, curArr...))
		return
	}

	for i := index; i < len(nums); i++ {
		if i > index && nums[i] == nums[i-1] {
			continue
		}
		curArr = append(curArr, nums[i])
		backend15(i+1, total+nums[i], nums, curArr, res)
		curArr = curArr[:len(curArr)-1]
	}
}
