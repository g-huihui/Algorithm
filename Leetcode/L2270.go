/**
 * @Author: Gong Yanhui
 * @Description: 2270. 分割数组的方案数
 * @Date: 2025-01-14 22:03
 */

package main

func waysToSplitArray(nums []int) int {

	var left = make([]int, len(nums)) // 前缀和
	left[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		left[i] = left[i-1] + nums[i]
	}

	var right = make([]int, len(nums)) // 后缀和
	right[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = right[i+1] + nums[i]
	}

	var res int
	for i := 0; i < len(nums)-1; i++ {
		if left[i] >= right[i+1] {
			res++
		}
	}

	return res
}

func main() {
	println(waysToSplitArray([]int{10, 4, -8, 7})) // 2
	println(waysToSplitArray([]int{2, 3, 1, 0}))   // 2
}
