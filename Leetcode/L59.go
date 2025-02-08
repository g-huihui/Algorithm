/**
 * @Author: Gong Yanhui
 * @Description: 59. 螺旋矩阵 II
 * @Date: 2025-02-08 11:25
 */

package main

func generateMatrix(n int) [][]int {

	// 定义返回的二维数组
	var res = make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	// 定义上下左右边界
	var left, right, top, bottom = 0, n - 1, 0, n - 1

	// 定义要循环的次数
	var all = n * n
	// 定义开始放置的数字
	var num = 1

	// 开始循环放置数字
	for num <= all {
		// 先放置最上面一行
		for i := left; i <= right; i++ {
			res[top][i] = num
			num++
		}
		top++

		// 放置最右边一列
		for i := top; i <= bottom; i++ {
			res[i][right] = num
			num++
		}
		right--

		// 放置最下面一行
		for i := right; i >= left; i-- {
			res[bottom][i] = num
			num++
		}
		bottom--

		// 放置最左边一列
		for i := bottom; i >= top; i-- {
			res[i][left] = num
			num++
		}
		left++
	}

	return res
}
