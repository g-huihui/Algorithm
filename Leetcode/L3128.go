/**
 * @Author: Gong Yanhui
 * @Description: 3128. 直角三角形
 * @Date: 2024-08-02 11:14
 */

package main

func numberOfRightTriangles(grid [][]int) int64 {

	var colLen = len(grid[0]) // 列数
	var rowLen = len(grid)    // 行数

	var colCount = make([]int64, rowLen) // 每列1数量
	var rowCount = make([]int64, colLen) // 每行1数量

	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			if grid[i][j] == 1 {
				colCount[i]++
				rowCount[j]++
			}
		}
	}

	var count int64
	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			if grid[i][j] == 1 {
				count += (colCount[i] - 1) * (rowCount[j] - 1)
			}
		}
	}

	return count
}

func main() {
	println(numberOfRightTriangles([][]int{{0, 1, 0}, {0, 1, 1}, {0, 1, 0}}))          // 2
	println(numberOfRightTriangles([][]int{{1, 0, 0, 0}, {0, 1, 0, 1}, {1, 0, 0, 0}})) // 0
}
