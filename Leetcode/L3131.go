/**
 * @Author: Gong Yanhui
 * @Description: 3131. 找出与数组相加的整数 I
 * @Date: 2024-08-08 17:44
 */

package main

import "sort"

func addedInteger(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	if len(nums1) == 0 || len(nums2) == 0 {
		return 0
	}
	return nums2[0] - nums1[0]
}
