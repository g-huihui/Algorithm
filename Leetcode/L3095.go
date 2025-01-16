/**
 * @Author: Gong Yanhui
 * @Description: 3095. 或值至少 K 的最短子数组 I
 * @Date: 2025-01-16 23:10
 */

package main

import "math"

func minimumSubarrayLength(nums []int, k int) int {
	var res = math.MaxInt

	for i := 0; i < len(nums); i++ {
		var sum = 0
		for j := i; j < len(nums); j++ {
			sum |= nums[j]
			if sum >= k {
				res = min(res, j-i+1)
				break
			}
		}
	}

	if res == math.MaxInt {
		return -1
	}

	return res
}

// 想使用双指针的办法 但是失败
func minimumSubarrayLength1(nums []int, k int) int {
	var res = 1
	var left, right, sum = 0, 0, 0
	for right < len(nums) {
		if sum < k {
			sum |= nums[right]
			right++
			if left == 0 && right == len(nums) {
				return -1
			}
		} else {
			res = min(res, right-left)
			sum ^= nums[left]
			left++
		}
	}

	return res
}

func main() {
	println(minimumSubarrayLength([]int{1, 2, 3}, 2))  // 1
	println(minimumSubarrayLength([]int{2, 1, 8}, 10)) // 3
	println(minimumSubarrayLength([]int{1, 2}, 0))     // 1

	println(minimumSubarrayLength([]int{1, 12, 2, 5}, 43)) // -1
}
