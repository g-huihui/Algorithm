/**
 * @Author: Gong Yanhui
 * @Description: 1742. 盒子中小球的最大数量
 * @Date: 2025-02-13 13:09
 */

package main

func countBalls(lowLimit int, highLimit int) int {
	var maxRes = 0
	var m = make(map[int]int) // 用于存储每个盒子中小球的数量
	for i := lowLimit; i <= highLimit; i++ {
		var total = getNumSum(i)
		m[total]++
		if m[total] > maxRes {
			maxRes = m[total]
		}
	}

	return maxRes
}

func getNumSum(num int) int {
	var total = 0
	for num != 0 {
		total += num % 10
		num /= 10
	}
	return total
}
