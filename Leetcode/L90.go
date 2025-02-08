/**
 * @Author: Gong Yanhui
 * @Description: 90. 子集 II
 * @Date: 2025-02-08 12:13
 */

package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	used := make([]bool, len(nums))
	back90(nums, used, 0, []int{}, &res)
	return res
}

func back90(nums []int, used []bool, index int, subArr []int, res *[][]int) {

	// index为遍历数组的下标 如果index等于数组长度 则说明已经遍历完了
	if index == len(nums) {
		*res = append(*res, append([]int{}, subArr...))
		return
	}

	// 不选择当前元素
	back90(nums, used, index+1, subArr, res)

	// 去除重复元素
	if index > 0 && nums[index] == nums[index-1] && !used[index-1] {
		return
	}

	// 选择当前元素
	used[index] = true
	subArr = append(subArr, nums[index])
	back90(nums, used, index+1, subArr, res)
	subArr = subArr[:len(subArr)-1]
	used[index] = false
}

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
	fmt.Println(subsetsWithDup([]int{0}))
}
