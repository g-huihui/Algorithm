/**
 * @Author: Gong Yanhui
 * @Description: 3033. 修改矩阵
 * @Date: 2024-07-06 16:44
 */

package main

func modifiedMatrix(matrix [][]int) [][]int {

	// 获取每列中最大元素
	var row, col = len(matrix), len(matrix[0])
	var maxCol = make([]int, col)
	var changedIndex = make([][]int, row)
	for i := 0; i < row; i++ {
		changedIndex[i] = make([]int, 0)
	}

	for i := 0; i < col; i++ {
		for j := 0; j < row; j++ {
			if matrix[j][i] > maxCol[i] {
				maxCol[i] = matrix[j][i]
			}
			if matrix[j][i] == -1 { // 记录-1的位置
				changedIndex[j] = append(changedIndex[j], i)
			}
		}
	}

	// 将-1的地方修改为当前列最大值
	for i := 0; i < row; i++ {
		for j := 0; j < len(changedIndex[i]); j++ {
			matrix[i][changedIndex[i][j]] = maxCol[changedIndex[i][j]]
		}
	}

	return matrix
}
