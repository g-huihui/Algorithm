/**
 * @Author: Gong Yanhui
 * @Description: 2639. 查询网格图中每一列的宽度
 * @Date: 2024-04-27 11:23
 */

package main

func findColumnWidth(grid [][]int) []int {
	n, m := len(grid), len(grid[0])
	res := make([]int, m)
	for j := 0; j < m; j++ {
		for i := 0; i < n; i++ {
			width := getNumWidth(grid[i][j])
			if width > res[j] {
				res[j] = width
			}
		}
	}
	return res
}

func getNumWidth(num int) int {
	res := 0
	if num <= 0 {
		res++
		num = -num
	}
	for num > 0 {
		res++
		num /= 10
	}
	return res
}
