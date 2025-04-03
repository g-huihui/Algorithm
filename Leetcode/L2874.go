/**
 * @Author: Gong Yanhui
 * @Description: 2874. 有序三元组中的最大值 II
 * @Date: 2025-04-03 19:18
 */

package main

func maximumTripletValue2(nums []int) int64 {

	var size = len(nums)

	// 获取从左到右的最大值
	var leftDp = make([]int, size)
	leftDp[0] = nums[0]
	for i := 1; i < size; i++ {
		leftDp[i] = max(leftDp[i-1], nums[i])
	}

	// 获取从又到左的最大值
	var rightDp = make([]int, size)
	rightDp[size-1] = nums[size-1]
	for i := size - 2; i >= 0; i-- {
		rightDp[i] = max(rightDp[i+1], nums[i])
	}

	// 计算最大值 遍历j的值
	var res int
	for i := 1; i < size-1; i++ {
		var num = (leftDp[i-1] - nums[i]) * rightDp[i+1]
		if num > res {
			res = num
		}
	}

	return int64(res)
}

func main() {
	// (num[i] - num[j]) * num[k]
	println(maximumTripletValue2([]int{12, 6, 1, 2, 7}))  // (0, 2, 4) => 77
	println(maximumTripletValue2([]int{1, 10, 3, 4, 19})) // (1, 2, 4) => 133
	println(maximumTripletValue2([]int{1, 2, 3}))         // (0, 1, 2) => 0
}
