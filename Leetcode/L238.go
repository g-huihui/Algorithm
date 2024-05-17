/**
 * @Author: Gong Yanhui
 * @Description: 238. 除自身以外数组的乘积
 * @Date: 2024-05-17 17:53
 */

package main

func productExceptSelf(nums []int) []int {
	length := len(nums)
	var (
		left  = make([]int, length)
		right = make([]int, length)
		res   = make([]int, length)
	)

	// Left初始化第一个为1
	left[0] = 1
	for i := 1; i < length; i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	// Right初始化最后一个为1
	right[length-1] = 1
	for i := length - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}

	// 左右两边相乘及中间值
	for i := 0; i < length; i++ {
		res[i] = left[i] * right[i]
	}

	return res
}

func main() {
	var nums = []int{1, 2, 3, 4}
	// Left  arr [1, 1, 2, 6]
	// Right arr [24, 12, 4, 1]
	println(productExceptSelf(nums)) // [24,12,8,6]
}
