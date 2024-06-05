/**
 * @Author: Gong Yanhui
 * @Description: 134. 加油站
 * @Date: 2024-06-05 11:21
 */

package main

import "math"

func canCompleteCircuit(gas []int, cost []int) int {
	var (
		length   = len(gas)
		lastGas  int // 剩余油量
		minIndex = 0
		minCost  = math.MaxInt
	)

	for i := 0; i < length; i++ {
		lastGas += gas[i] - cost[i]
		if lastGas <= minCost {
			minCost = lastGas
			minIndex = i
		}
	}

	if lastGas < 0 {
		return -1
	}

	return (minIndex + 1) % length
}
