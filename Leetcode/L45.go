/**
 * @Author: Gong Yanhui
 * @Description: 45. 跳跃游戏 II
 * @Date: 2025-01-27 14:02
 */

package main

import "math"

// 官方解法 贪心算法 反向查找出发位置 时间复杂度 O(n) 空间复杂度 O(1)
func jump(nums []int) int {
	var position = len(nums) - 1 // 代表当前需要到达的位置
	var steps = 0                // 代表跳跃的步数

	for position > 0 { // 只要没有到达起始位置 就继续
		for i := 0; i < position; i++ { // 从前往后找到第一个可以到达当前位置的位置
			if nums[i]+i >= position { // 如果可以到达
				position = i // 更新当前位置
				steps++      // 步数加一
				break
			}
		}
	}

	return steps
}

// 官方解法2 优化 正向查找可到达的最大位置 时间复杂度 O(n) 空间复杂度 O(1)
func jump2(nums []int) int {
	var end = 0         // 已经到达的最远位置
	var maxPosition = 0 // 可以到达的最远位置
	var steps = 0

	for i := 0; i < len(nums)-1; i++ {
		maxPosition = max(maxPosition, nums[i]+i)
		if i == end { // 到达了上一次跳跃的最远位置
			end = maxPosition
			steps++
		}
	}

	return steps
}

// 从后往前看到达最后一天所需要的最小步数 (超出时间限制)
func jump1(nums []int) int {
	return back45(nums, len(nums)-1)
}

// 回溯算法
func back45(nums []int, index int) int {
	if index == 0 {
		return 0
	}

	var step = math.MaxInt
	for i := 0; i < index; i++ {
		if nums[i]+i >= index {
			step = min(step, back45(nums, i)+1)
		}
	}

	return step
}

func main() {
	println(jump([]int{2, 3, 1, 1, 4})) // 2
	println(jump([]int{2, 3, 0, 1, 4})) // 2

	println(jump1(
		[]int{
			5, 6, 4, 4, 6, 9, 4, 4, 7, 4, 4, 8, 2, 6, 8, 1, 5, 9, 6, 5, 2, 7, 9, 7, 9, 6, 9, 4, 1, 6, 8, 8, 4, 4, 2, 0, 3, 8, 5,
		})) // 5 超时
}
