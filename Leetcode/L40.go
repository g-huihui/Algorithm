/**
 * @Author: Gong Yanhui
 * @Description: 40. 组合总和 II
 * @Date: 2024-04-21 21:26
 */

package main

import (
	"fmt"
	"sort"
)

//
// combinationSum2
//  @Description:
//  @param candidates 每个数字在每个组合中只能使用 一次
//  @param target
//  @return [][]int
//
func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	var path []int
	sort.Ints(candidates) // 排序优化下选择 当减到负数就不需要往后继续了
	backtrack2(candidates, target, 0, path, &res)
	return res
}

func backtrack2(candidates []int, target int, index int, path []int, res *[][]int) {
	if target == 0 {
		*res = append(*res, append([]int{}, path...))
		return
	}
	for i := index; i < len(candidates); i++ {
		if target < candidates[i] {
			break
		}
		if i > index && candidates[i] == candidates[i-1] {
			continue
		}
		path = append(path, candidates[i])
		backtrack2(candidates, target-candidates[i], i+1, path, res)
		path = path[:len(path)-1]
	}
}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)) // [[1 1 6] [1 2 5] [1 7] [2 6]]
	fmt.Println(combinationSum2([]int{2, 5, 2, 1, 2}, 5))        // [[1 2 2] [5]]
}
