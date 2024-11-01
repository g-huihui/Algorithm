/**
 * @Author: Gong Yanhui
 * @Description: 3259. 超级饮料的最大强化能量
 * @Date: 2024-11-01 16:56
 */

package main

func maxEnergyBoost(energyDrinkA []int, energyDrinkB []int) int64 {
	n := len(energyDrinkA)
	dp := make([][2]int64, n+1)
	for i := 1; i <= n; i++ {
		dp[i][0] = dp[i-1][0] + int64(energyDrinkA[i-1])
		dp[i][1] = dp[i-1][1] + int64(energyDrinkB[i-1])
		if i >= 2 {
			dp[i][0] = maxInt64(dp[i][0], dp[i-2][1]+int64(energyDrinkA[i-1]))
			dp[i][1] = maxInt64(dp[i][1], dp[i-2][0]+int64(energyDrinkB[i-1]))
		}
	}

	return maxInt64(dp[n][0], dp[n][1])
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	energyDrinkA := []int{4, 1, 1}
	energyDrinkB := []int{1, 1, 3}
	println(maxEnergyBoost(energyDrinkA, energyDrinkB))
}
