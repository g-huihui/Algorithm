/**
 * @Author: Gong Yanhui
 * @Description: 63. 不同路径 II
 * @Date: 2025-02-08 11:35
 */

package main

// 使用动态规划的方式 判断到每个点的路径可能数
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	var row = len(obstacleGrid)
	var col = len(obstacleGrid[0])

	// 创建动态规划递归数组
	var dp = make([][]int, row)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, col)
	}

	// 初始化第一个点
	if obstacleGrid[0][0] != 0 {
		return 0
	}
	dp[0][0] = 1

	// 遍历每个点 判断到达该点的路径数
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			// 如果当前是障碍物
			if obstacleGrid[i][j] != 0 {
				dp[i][j] = 0
				continue
			}

			// 如果是第一个点
			if i == 0 && j == 0 {
				continue
			}

			// 如果是第一行
			if i == 0 {
				dp[i][j] = dp[i][j-1]
				continue
			}

			// 如果是第一列
			if j == 0 {
				dp[i][j] = dp[i-1][j]
				continue
			}

			// 如果是其他点
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[row-1][col-1]
}

func uniquePathsWithObstacles1(obstacleGrid [][]int) int {
	res := 0
	back63(obstacleGrid, 0, 0, &res)

	return res
}

// 使用回溯的方式找到所有的路径 会超时
func back63(obstacleGrid [][]int, i, j int, res *int) {
	// 走到终点了 结果+1
	if i == len(obstacleGrid)-1 && j == len(obstacleGrid[0])-1 && obstacleGrid[i][j] == 0 {
		*res++
		return
	}

	// 如果当前位置不是0
	if obstacleGrid[i][j] != 0 {
		return
	}

	// 向下走
	if i < len(obstacleGrid)-1 {
		back63(obstacleGrid, i+1, j, res)
	}
	// 向右走
	if j < len(obstacleGrid[0])-1 {
		back63(obstacleGrid, i, j+1, res)
	}
}

func main() {
	println(uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}})) // 2
	println(uniquePathsWithObstacles([][]int{{0, 1}, {0, 0}}))                  // 1

	println(uniquePathsWithObstacles([][]int{{0, 0}, {0, 1}})) // 0
}
