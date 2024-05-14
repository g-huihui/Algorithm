/**
 * @Author: Gong Yanhui
 * @Description: 169. 多数元素
 * @Date: 2024-05-14 11:15
 */

package main

import "sort"

func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// 数组排序后 索引为2/n的元素一定是众数
