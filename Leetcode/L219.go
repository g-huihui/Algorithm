/**
 * @Author: Gong Yanhui
 * @Description: 219. 存在重复元素 II
 * @Date: 2024-12-02 09:31
 */

package main

func containsNearbyDuplicate(nums []int, k int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if j-i > k {
				break
			}
			if nums[i] == nums[j] {
				return true
			}
		}
	}

	return false
}

// [99,99] ==> true but false
func containsNearbyDuplicate1(nums []int, k int) bool {
	for i := 0; i < len(nums)-k; i++ {
		for j := i + 1; j <= i+k; j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}
