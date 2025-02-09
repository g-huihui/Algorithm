/**
 * @Author: Gong Yanhui
 * @Description: 1572. 矩阵对角线元素的和
 * @Date: 2025-02-09 15:21
 */

package main

func diagonalSum(mat [][]int) int {
	var size, total = len(mat), 0
	for i := 0; i < size; i++ {
		total += mat[i][i]
		total += mat[i][len(mat)-1-i]
	}
	if size%2 == 1 {
		total -= mat[size/2][size/2]
	}

	return total
}
