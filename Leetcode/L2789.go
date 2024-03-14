/**
 * @Author: Gong Yanhui
 * @Description: 2789. 合并后数组中的最大元素
 * @Date: 2024-03-14 19:41
 */

package main

func maxArrayValue(nums []int) int64 {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return int64(nums[0])
	}

	var maxNum int64
	for i := len(nums) - 1; i >= 0; i-- {
		if i == 0 && int64(nums[i]) > maxNum {
			maxNum = int64(nums[i])
			break
		}
		if nums[i] >= nums[i-1] {
			nums[i-1] = nums[i] + nums[i-1]
		} else {
			if int64(nums[i]) > maxNum {
				maxNum = int64(nums[i])
			}
		}
	}
	return maxNum
}

func main() {
	var nums1 = []int{1, 2, 3, 4, 5}
	var result1 = maxArrayValue(nums1)
	println(result1) // 15

	var nums2 = []int{2, 3, 7, 9, 3}
	var result2 = maxArrayValue(nums2)
	println(result2) // 21

	var nums3 = []int{5, 3, 3}
	var result3 = maxArrayValue(nums3)
	println(result3) // 11
}
