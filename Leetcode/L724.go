/**
 * @Author: Gong Yanhui
 * @Description: 724. 寻找数组的中心下标
 * @Date: 2024-07-11 20:38
 */

package main

func pivotIndex(nums []int) int {
	var length = len(nums)

	var leftSum = make([]int, length)
	leftSum[0] = nums[0]
	for i := 1; i < length; i++ {
		leftSum[i] = leftSum[i-1] + nums[i]
	}

	var rightSum = make([]int, length)
	rightSum[length-1] = nums[length-1]
	for i := length - 2; i >= 0; i-- {
		rightSum[i] = rightSum[i+1] + nums[i]
	}

	for i := 0; i < length; i++ {
		if leftSum[i] == rightSum[i] {
			return i
		}
	}

	return -1
}

func main() {
	println(pivotIndex([]int{1, 7, 3, 6, 5, 6})) // 3
	println(pivotIndex([]int{1, 2, 3}))          // -1
	println(pivotIndex([]int{2, 1, -1}))         // 0
}
