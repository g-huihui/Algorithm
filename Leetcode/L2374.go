/**
 * @Author: Gong Yanhui
 * @Description: 2374. 边积分最高的节点
 * @Date: 2024-09-21 19:31
 */

package main

func edgeScore(edges []int) int {
	var length = len(edges)
	var totalScore = make([]int, length)
	for i := 0; i < length; i++ {
		totalScore[edges[i]] += i
	}

	var maxRes = 0
	var maxIndex = -1
	for i := 0; i < length; i++ {
		if totalScore[i] > maxRes {
			maxRes = totalScore[i]
			maxIndex = i
		}
	}

	return maxIndex
}
