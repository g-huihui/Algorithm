/**
 * @Author: Gong Yanhui
 * @Description: 2529. 正整数和负整数的最大计数
 * @Date: 2024-04-09 14:28
 */

package main

func maximumCount(nums []int) int {
	var negativeCount, positiveCount int
	for _, num := range nums {
		if num == 0 {
			continue
		}
		if num < 0 {
			negativeCount++
		} else {
			positiveCount++
		}
	}
	if negativeCount > positiveCount {
		return negativeCount
	} else {
		return positiveCount
	}
}
