/**
 * @Author: Gong Yanhui
 * @Description: 350. 两个数组的交集 II
 * @Date: 2025-01-30 18:56
 */

package main

import "fmt"

// 方法一 哈希表 时间复杂度O(n+m) 空间复杂度O(min(n,m))
func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		// 确保每次都是nums1更短
		return intersect(nums2, nums1)
	}

	var m = make(map[int]int)
	for _, num := range nums1 {
		m[num]++
	}

	var res = make([]int, 0)
	for _, num := range nums2 {
		if m[num] > 0 {
			res = append(res, num)
			m[num]--
		}
	}

	return res
}

func main() { // 题目表述 可以不考虑输出结果的顺序
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))       // [2,2]
	fmt.Println(intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4})) // [4,9]
}
