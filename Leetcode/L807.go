/**
 * @Author: Gong Yanhui
 * @Description: 807. 保持城市天际线
 * @Date: 2024-07-14 18:58
 */

package main

func maxIncreaseKeepingSkyline(grid [][]int) int {

	var rowLen = len(grid)
	var colLen = len(grid[0])

	// 获取每行每列最大值
	var rowMax = make([]int, rowLen)
	var colMax = make([]int, colLen)
	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			if rowMax[i] < grid[i][j] {
				rowMax[i] = grid[i][j]
			}
			if colMax[j] < grid[i][j] {
				colMax[j] = grid[i][j]
			}
		}
	}

	var res int
	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			if grid[i][j] < rowMax[i] && grid[i][j] < colMax[j] {
				res = res + (min(rowMax[i], colMax[j]) - grid[i][j])
			}
		}
	}

	return res
}
