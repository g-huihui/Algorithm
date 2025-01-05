/**
 * @Author: Gong Yanhui
 * @Description: 120. 三角形最小路径和
 * @Date: 2025-01-05 12:03
 */

package main

func minimumTotal(triangle [][]int) int {
	var height = len(triangle) // 树高
	var dp = make([][]int, height)
	for i := 0; i < height; i++ {
		dp[i] = make([]int, len(triangle[i]))
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < height; i++ { // 每一层
		for j := 0; j < len(triangle[i]); j++ { // 每一层的每一个节点
			if j == 0 { // 每一层的第一个节点
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			} else if j == len(triangle[i])-1 { // 每一层的最后一个节点
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
			}
		}
	}

	// 找到最后一层的最小值
	var res = dp[height-1][0]
	for i := 1; i < len(dp[height-1]); i++ {
		res = min(res, dp[height-1][i])
	}

	return res
}

// 动态规划
// triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
// dp		= [[2],[5,6],[11,10,13],[15,11,18,16]]
