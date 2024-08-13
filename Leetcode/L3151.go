/**
 * @Author: Gong Yanhui
 * @Description: 3151. 特殊数组 I
 * @Date: 2024-08-13 09:57
 */

package main

func isArraySpecial(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		return true
	}

	for i := 1; i < len(nums); i++ {
		if isOdd(nums[i]) == isOdd(nums[i-1]) {
			return false
		}
	}
	return true
}

// 判断奇偶数 奇数
func isOdd(num int) bool {
	return num%2 == 1
}
