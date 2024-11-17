/**
 * @Author: Gong Yanhui
 * @Description: 3239. 最少翻转次数使二进制矩阵回文 I
 * @Date: 2024-11-17 17:10
 */

package main

func minFlips(grid [][]int) int {

	var row = len(grid)
	var col = len(grid[0])

	// 1. 先按照行进行翻转
	var rowCount = 0
	for i := 0; i < row; i++ {
		for j := 0; j < col/2; j++ {
			if grid[i][j] != grid[i][col-j-1] {
				rowCount++
			}
		}
	}

	// 2. 再按照列进行翻转
	var colCount = 0
	for i := 0; i < col; i++ {
		for j := 0; j < row/2; j++ {
			if grid[j][i] != grid[row-j-1][i] {
				colCount++
			}
		}
	}

	// 取最小值
	return min(rowCount, colCount)
}

func main() {
	println(minFlips([][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 1}})) // 2
	println(minFlips([][]int{{0, 1}, {0, 1}, {0, 0}}))          // 1
	println(minFlips([][]int{{1}, {0}}))                        // 0
}
