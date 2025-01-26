/**
 * @Author: Gong Yanhui
 * @Description: 682. 棒球比赛
 * @Date: 2025-01-26 16:43
 */

package main

import "strconv"

func calPoints(operations []string) int {
	var stack []int // 用于存储每轮的得分
	for _, op := range operations {
		switch op {
		case "+":
			stack = append(stack, stack[len(stack)-1]+stack[len(stack)-2])
		case "D":
			stack = append(stack, stack[len(stack)-1]*2)
		case "C":
			stack = stack[:len(stack)-1]
		default:
			stack = append(stack, strToInt(op))
		}
	}

	return sum(stack)
}

func sum(arr []int) int {
	var total = 0
	for _, num := range arr {
		total += num
	}
	return total
}

func strToInt(op string) int {
	num, _ := strconv.Atoi(op)
	return num
}
