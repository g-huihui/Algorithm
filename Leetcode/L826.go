/**
 * @Author: Gong Yanhui
 * @Description: 826. 安排工作以达到最大收益
 * @Date: 2024-05-17 10:43
 */

package main

import "sort"

// 通过率 50 / 57 没想通哪里有问题
func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	var m = make(map[int]int) // difficulty -> profit
	for i := 0; i < len(difficulty); i++ {
		m[difficulty[i]] = profit[i]
	}
	var resMaxProfit int
	sort.Ints(difficulty)
	for i := 0; i < len(worker); i++ {
		var index int
		var maxMoney int
		for index < len(difficulty) && difficulty[index] <= worker[i] {
			if m[difficulty[index]] > maxMoney {
				maxMoney = m[difficulty[index]]
			}
			index++
		}
		resMaxProfit += maxMoney
	}
	return resMaxProfit
}

func main() {
	//println(maxProfitAssignment([]int{13, 37, 58}, []int{4, 90, 96}, []int{34, 73, 45}))

	println(maxProfitAssignment(
		[]int{66, 1, 28, 73, 53, 35, 45, 60, 100, 44, 59, 94, 27, 88, 7, 18, 83, 18, 72, 63},
		[]int{66, 20, 84, 81, 56, 40, 37, 82, 53, 45, 43, 96, 67, 27, 12, 54, 98, 19, 47, 77},
		[]int{61, 33, 68, 38, 63, 45, 1, 10, 53, 23, 66, 70, 14, 51, 94, 18, 28, 78, 100, 16},
	)) // 1392
}

// 官方题解
func maxProfitAssignment2(difficulty []int, profit []int, worker []int) int {
	jobs := make([][2]int, len(difficulty))
	for i := range difficulty {
		jobs[i] = [2]int{difficulty[i], profit[i]}
	}
	sort.Slice(jobs, func(i, j int) bool { return jobs[i][0] < jobs[j][0] })
	sort.Ints(worker)

	res, i, best := 0, 0, 0
	for _, w := range worker {
		for i < len(jobs) && w >= jobs[i][0] {
			best = max(best, jobs[i][1])
			i++
		}
		res += best
	}
	return res
}
