/**
 * @Author: Gong Yanhui
 * @Description: 2956. 找到两个数组中的公共元素
 * @Date: 2024-07-16 10:52
 */

package main

func findIntersectionValues(nums1 []int, nums2 []int) []int {

	var len1 = len(nums1)
	var len2 = len(nums2)

	var m1 = make(map[int]struct{}, len1)
	var m2 = make(map[int]struct{}, len2)
	for i := 0; i < len1; i++ {
		m1[nums1[i]] = struct{}{}
	}
	for i := 0; i < len2; i++ {
		m2[nums2[i]] = struct{}{}
	}

	var res = make([]int, 2)
	for i := 0; i < len1; i++ {
		if _, ok := m2[nums1[i]]; ok {
			res[0]++
		}
	}
	for i := 0; i < len2; i++ {
		if _, ok := m1[nums2[i]]; ok {
			res[1]++
		}
	}

	return res
}
