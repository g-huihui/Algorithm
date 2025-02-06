/**
 * @Author: Gong Yanhui
 * @Description: 47. 全排列 II
 * @Date: 2025-02-06 19:26
 */

package main

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	used := make([]bool, len(nums))
	var res [][]int
	back47(nums, []int{}, used, &res)

	return res
}

// 回溯算法 index为当前应该放置的位置
func back47(nums, cur []int, used []bool, res *[][]int) {
	if len(cur) == len(nums) {
		*res = append(*res, append([]int{}, cur...))
		return
	}

	for i := range nums {
		if used[i] { // 如果改位置已经被使用过
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && !used[i-1] { // 如果当前位置和前一个相同并且前一个未使用
			// 防止重复
			continue
		}
		used[i] = true
		cur = append(cur, nums[i])
		back47(nums, cur, used, res)
		used[i] = false
		cur = cur[:len(cur)-1]
	}
}
