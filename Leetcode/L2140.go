/**
 * @Author: Gong Yanhui
 * @Description: 2140. 解决智力问题
 * @Date: 2025-04-01 23:32
 */

package main

func mostPoints(questions [][]int) int64 {
	var size = len(questions)

	var dp = make([]int, size+1)
	dp[size-1] = questions[size-1][0]
	for i := size - 1; i >= 0; i-- {
		// 如果解当前题 则当前题后面会冻结的题不能解 此时得分
		var num = questions[i][0]
		if i+questions[i][1]+1 < size {
			num += dp[i+questions[i][1]+1]
		}
		// 如果不解当前题 则 dp[i] = dp[i+1]
		dp[i] = max(dp[i+1], num)
	}

	return int64(dp[0])
}

func main() {
	println(mostPoints([][]int{{3, 2}, {4, 3}, {4, 4}, {2, 5}}))         // 5
	println(mostPoints([][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}})) // 7

	println(mostPoints([][]int{{12, 46}, {78, 19}, {63, 15}, {79, 62}, {13, 10}}))
}
