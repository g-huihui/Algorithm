/**
 * @Author: Gong Yanhui
 * @Description: 209. 长度最小的子数组
 * @Date: 2024-07-04 18:26
 */

package main

// 2024-11-30 15:45 第二次尝试
func minSubArrayLen(target int, nums []int) int {
	var minDistance = len(nums)
	var head, tail, total int
	for {
		if tail == len(nums) && total < target {
			// 结束条件
			break
		}

		if total < target {
			total += nums[tail]
			tail++
		} else {
			if minDistance > tail-head {
				minDistance = tail - head
			}
			total -= nums[head]
			head++
		}
	}

	if minDistance == len(nums) && head == 0 {
		return 0
	}

	return minDistance
}

// 暴力破解 超出时间限制
func minSubArrayLen1(target int, nums []int) int {
	var minDistance = len(nums)
	var flag bool
	for head := 0; head < len(nums); head++ {
		var total int
		for tail := head; tail < len(nums); tail++ {
			total += nums[tail]
			if total >= target && minDistance >= (tail-head+1) {
				flag = true
				minDistance = tail - head + 1
				break
			}
		}
	}

	if minDistance == len(nums) && !flag {
		return 0
	} else {
		return minDistance
	}
}

func main() {
	println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))        // 2
	println(minSubArrayLen(4, []int{1, 4, 4}))                 // 1
	println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1})) // 0

	println(minSubArrayLen(15, []int{5, 1, 3, 5, 10, 7, 4, 9, 2, 8})) // 2
	println(minSubArrayLen(15, []int{1, 2, 3, 4, 5}))                 // 5
}
