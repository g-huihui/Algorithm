/**
 * @Author: Gong Yanhui
 * @Description: 78. 子集
 * @Date: 2025-03-02 13:25
 */

package main

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	var res [][]int
	var path []int
	dfs78(nums, 0, path, &res)

	return res
}

func dfs78(nums []int, index int, path []int, res *[][]int) {
	if index == len(nums) {
		*res = append(*res, append([]int{}, path...))
		return
	}

	// 不选择当前元素
	dfs78(nums, index+1, path, res)

	// 选择当前元素
	path = append(path, nums[index])
	dfs78(nums, index+1, path, res)
	path = path[:len(path)-1] // 恢复现场
}
