/**
 * @Author: Gong Yanhui
 * @Description: 39. 组合总和
 * @Date: 2024-04-20 17:31
 */

package main

import (
	"fmt"
	"sort"
)

//
// combinationSum
//  @Description:
//  @param candidates 无重复元素的数组 同一个数字可以无限制重复被选取
//  @param target
//  @return [][]int
//
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var path []int
	sort.Ints(candidates) // 排序优化下选择 当减到负数就不需要往后继续了
	backtrack(candidates, target, 0, path, &res)
	return res
}

//
// backtrack
//  @Description: 回溯算法
//  @param candidates
//  @param target
//  @param start 当前candidates应该被选择的位置
//  @param path 存放已经选择好的路径
//  @param res 最终结果
//
func backtrack(candidates []int, target int, start int, path []int, res *[][]int) {
	if target == 0 { // 代表一次计算正确的结果组合
		*res = append(*res, append([]int{}, path...)) // 将改正确结果存储到res中结束当前分支
		return
	}
	for i := start; i < len(candidates); i++ {
		if target < candidates[i] { // 说明当前组合已经不合适 可以结束当前分支了
			break
		}
		path = append(path, candidates[i]) // flag
		backtrack(candidates, target-candidates[i], i, path, res)
		path = path[:len(path)-1] // 对flag行的操作进行回滚 继续下一个candidates数据尝试
	}
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7)) // [[2 2 3] [7]]
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))    // [[2 2 2 2] [2 3 3] [3 5]]
	fmt.Println(combinationSum([]int{2}, 1))          // []
}
