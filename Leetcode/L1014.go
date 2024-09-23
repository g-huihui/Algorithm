/**
 * @Author: Gong Yanhui
 * @Description: 1014. 最佳观光组合
 * @Date: 2024-09-23 11:20
 */

package main

import "math"

func maxScoreSightseeingPair(values []int) int {
	var begin, MaxRes = values[0], values[0]
	for i := 1; i < len(values); i++ {
		MaxRes = max(MaxRes, begin+values[i]-i)
		begin = max(begin, values[i]+i)
	}
	return MaxRes
}

// 超时
func maxScoreSightseeingPair1(values []int) int {
	var maxRes = math.MinInt
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			var tmp = values[i] + values[j] + i - j
			if maxRes < tmp {
				maxRes = tmp
			}
		}
	}
	return maxRes
}

func main() {
	println(maxScoreSightseeingPair([]int{8, 1, 5, 2, 6})) // [0,2] = 11 -> 8 + 5 - (2 - 0)
	println(maxScoreSightseeingPair([]int{1, 2}))          // [0,1] = 2 -> 2 + 1 - (1 - 0)
}
