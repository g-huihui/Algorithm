/**
 * @Author: Gong Yanhui
 * @Description: 2713. 矩阵中严格递增的单元格数
 * @Date: 2024-06-19 15:03
 */

package main

// 超时 需要优化
func maxIncreasingCells(mat [][]int) int {
	n := len(mat)
	m := len(mat[0])

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	var count int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			dp[i][j] = 1
			res := deepFunc(mat, dp, i, j, 1)
			if res > count {
				count = res
			}
			dp[i][j] = 0
		}
	}

	return count
}

func deepFunc(mat [][]int, dp [][]int, i int, j int, count int) int {

	n := len(mat)
	m := len(mat[0])

	var res = count

	// 整行遍历
	for k := 0; k < m; k++ {
		if k == j {
			continue
		}

		if dp[i][k] == 0 && mat[i][k] > mat[i][j] {
			dp[i][k] = 1
			tmp := deepFunc(mat, dp, i, k, count+1)
			if tmp > res {
				res = tmp
			}
			dp[i][k] = 0
		}
	}

	// 整列遍历
	for k := 0; k < n; k++ {
		if k == i {
			continue
		}

		if dp[k][j] == 0 && mat[k][j] > mat[i][j] {
			dp[k][j] = 1
			tmp := deepFunc(mat, dp, k, j, count+1)
			if tmp > res {
				res = tmp
			}
			dp[k][j] = 0
		}
	}

	return res
}

func main() {
	println(maxIncreasingCells([][]int{{3, 1}, {3, 4}}))        // 2
	println(maxIncreasingCells([][]int{{1, 1}, {1, 1}}))        // 1
	println(maxIncreasingCells([][]int{{3, 1, 6}, {-9, 5, 7}})) // 4

	println(maxIncreasingCells(
		[][]int{{39, -7, 48, -13, -23, 34, 34, 13, 23, -14, -49, 24, 34, 1, 19, 47,
			-36, 29, -1, 1, -50, 42, 27, 11, -5, -37, 20, 38, 43, 3, -23, -41, 22}}))
}
