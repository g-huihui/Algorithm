/**
 * @Author: Gong Yanhui
 * @Description: 3375. 使数组的值全部为 K 的最少操作次数
 * @Date: 2025-04-09 15:19
 */

package main

import "sort"

// 题目未看懂 全看题解
func minOperations1(nums []int, k int) int {
	sort.Ints(nums)

	if k > nums[0] { // 若数字存在比k小的数字 则无解
		return -1
	}

	var m = make(map[int]struct{}) // 统计数组中所有大于k的数字种类个数 即为操作数
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > k {
			m[nums[i]] = struct{}{}
		} else {
			break
		}
	}

	return len(m)
}
