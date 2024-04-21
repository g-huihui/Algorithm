/**
 * @Author: Gong Yanhui
 * @Description: 216. 组合总和 III
 * @Date: 2024-04-21 21:49
 */

package main

// 和为n的k个数组合
func combinationSum3(k int, n int) [][]int {
	var res [][]int
	var path []int
	backtrack3(k, n, 1, path, &res)
	return res
}

func backtrack3(k int, n int, start int, path []int, res *[][]int) {
	if n == 0 && len(path) == k {
		*res = append(*res, append([]int{}, path...))
		return
	}
	for num := start; num < 10; num++ {
		if n < num {
			break
		}
		path = append(path, num)
		backtrack3(k, n-num, num+1, path, res)
		path = path[:len(path)-1]
	}
}
