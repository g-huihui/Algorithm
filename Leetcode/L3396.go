/**
 * @Author: Gong Yanhui
 * @Description: 3396. 使数组元素互不相同所需的最少操作次数
 * @Date: 2025-04-08 10:45
 */

package main

// 从后往前扫一次 时间复杂度O(n)
func minimumOperations(nums []int) int {
	var m = make(map[int]struct{}, len(nums))
	var index = -1 // 从后往前找 全是唯一元素的分界点
	for i := len(nums) - 1; i >= 0; i-- {
		if _, ok := m[nums[i]]; !ok {
			m[nums[i]] = struct{}{}
		} else {
			index = i // 找到第一个重复元素 index是要排除的位置
			break
		}
	}
	if index == -1 {
		return 0
	}

	var count = (index + 1) / 3
	if ((index + 1) % 3) == 0 {
		return count
	} else {
		return count + 1
	}
}
