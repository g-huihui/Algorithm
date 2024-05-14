/**
 * @Author: Gong Yanhui
 * @Description: 122. 买卖股票的最佳时机 II
 * @Date: 2024-05-14 16:15
 */

package main

func maxProfit2(prices []int) int {
	var length = len(prices)
	var dp = make([][2]int, length)
	dp[0][1] = -prices[0]
	for i := 1; i < length; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[length-1][0]
}
