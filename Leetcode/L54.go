/**
 * @Author: Gong Yanhui
 * @Description: 54. 螺旋矩阵
 * @Date: 2024-12-01 11:23
 */

package main

func spiralOrder(matrix [][]int) []int {
	// 定义上下边界
	top, bottom := 0, len(matrix)-1
	// 定义左右边界
	left, right := 0, len(matrix[0])-1
	// 定义结果数组
	res := make([]int, 0, len(matrix)*len(matrix[0]))

	for top <= bottom && left <= right {
		// 从左到右
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		// 从上到下
		for i := top + 1; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		// 从右到左
		for i := right - 1; i >= left && top < bottom; i-- {
			res = append(res, matrix[bottom][i])
		}
		// 从下到上
		for i := bottom - 1; i > top && left < right; i-- {
			res = append(res, matrix[i][left])
		}
		// 缩小边界
		top++
		bottom--
		left++
		right--
	}

	return res
}
