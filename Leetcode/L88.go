/**
 * @Author: Gong Yanhui
 * @Description: 88. 合并两个有序数组
 * @Date: 2024-06-05 11:35
 */

package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	var index1, index2 int
	var tmpNum = make([]int, 0, m+n)

	for {
		if index1 == m {
			tmpNum = append(tmpNum, nums2[index2:]...)
			break
		}
		if index2 == n {
			tmpNum = append(tmpNum, nums1[index1:]...)
			break
		}
		if nums1[index1] < nums2[index2] {
			tmpNum = append(tmpNum, nums1[index1])
			index1++
		} else {
			tmpNum = append(tmpNum, nums2[index2])
			index2++
		}
	}

	copy(nums1, tmpNum)
}
