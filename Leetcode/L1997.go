/**
 * @Author: Gong Yanhui
 * @Description: 1997. 访问完所有房间的第一天
 * @Date: 2024-03-28 10:55
 */

package main

// 正常解法 超时
func firstDayBeenInAllRooms1(nextVisit []int) int {
	var length = len(nextVisit)
	var dp = make([]int, length)
	var index, count = 0, 0
	for {
		if dp[index] == 0 { // 如果当前房间未访问过 访问当前房间
			length--
		}
		dp[index]++
		if dp[index]%2 == 1 {
			index = nextVisit[index]
		} else {
			index = (index + 1) % len(nextVisit)
		}
		if length == 0 { // 如果长度为0说明已经访问完所有房间
			break
		}
		count++
	}
	return count
}

func firstDayBeenInAllRooms(nextVisit []int) int {
	mod := 1000000007
	dp := make([]int, len(nextVisit))

	dp[0] = 2 //初始化原地待一天 + 访问下一个房间一天
	for i := 1; i < len(nextVisit); i++ {
		to := nextVisit[i]
		dp[i] = dp[i-1] + 2
		if to != 0 {
			dp[i] = (dp[i] - dp[to-1] + mod) % mod //避免负数
		}
		dp[i] = (dp[i] + dp[i-1]) % mod //题目保证n >= 2
	}

	return dp[len(nextVisit)-2]
}

func main() {
	println(firstDayBeenInAllRooms([]int{0, 0}))    // 2
	println(firstDayBeenInAllRooms([]int{0, 0, 2})) // 6
}
