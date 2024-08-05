/**
 * @Author: Gong Yanhui
 * @Description: 600. 不含连续1的非负整数
 * @Date: 2024-08-05 20:16
 */

package main

// 官方题解
func findIntegers(n int) int {
	var ans int

	dp := [31]int{1, 1} // dp[i] 表示长度为 i 的合法数的个数
	for i := 2; i < 31; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	for i, pre := 29, 0; i >= 0; i-- {
		val := 1 << i
		if n&val > 0 {
			ans += dp[i+1]
			if pre == 1 {
				break
			}
			pre = 1
		} else {
			pre = 0
		}
		if i == 0 { // 说明 n 是 1111...1
			ans++
		}
	}
	return ans
}

func main() {
	println(findIntegers(5)) // 5
	println(findIntegers(1)) // 2
	println(findIntegers(2)) // 3
}

/**
0 : 0
1 : 1
2 : 10
3 : 11
4 : 100
5 : 101
6 : 110
7 : 111
8 : 1000
*/
