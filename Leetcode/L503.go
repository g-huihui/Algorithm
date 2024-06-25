/**
 * @Author: Gong Yanhui
 * @Description: 503. 下一个更大元素 II
 * @Date: 2024-06-25 11:17
 */

package main

func nextGreaterElements(nums []int) []int {
	var res []int
	for i := 0; i < len(nums); i++ {
		res = append(res, findNextBigger(nums, i))
	}
	return res
}

func findNextBigger(nums []int, index int) int {
	for i := (index + 1) % len(nums); i != index; i = (i + 1) % len(nums) {
		if nums[i] > nums[index] {
			return nums[i]
		}
	}
	return -1
}

func main() {
	res1 := nextGreaterElements([]int{1, 2, 1})
	println(res1) // [2, -1, 2]

	res2 := nextGreaterElements([]int{1, 2, 3, 4, 3})
	println(res2) // [2, 3, 4, -1, 4]
}
