/**
 * @Author: Gong Yanhui
 * @Description: 3132. 找出与数组相加的整数 II
 * @Date: 2024-08-09 10:36
 */

package main

import (
	"math"
	"sort"
)

func minimumAddedInteger(nums1 []int, nums2 []int) int {

	var minX = math.MaxInt
	sort.Ints(nums1)
	sort.Ints(nums2)

	// 遍历排除的两个数索引
	for i := 0; i < len(nums1)-1; i++ {
		for j := i + 1; j < len(nums1); j++ {
			var tmpArr = make([]int, 0, len(nums2))
			for index, num := range nums1 {
				if index != i && index != j {
					tmpArr = append(tmpArr, num)
				}
			}
			var x = nums2[0] - tmpArr[0]
			var k int
			for k = 1; k < len(nums2); k++ {
				if nums2[k]-tmpArr[k] != x {
					break
				}
			}
			if k == len(nums2) && x < minX {
				minX = x
			}
		}
	}

	if minX == math.MaxInt {
		return 0
	}

	return minX
}

func main() {
	println(minimumAddedInteger([]int{4, 20, 16, 12, 8}, []int{14, 18, 10})) // -2
	println(minimumAddedInteger([]int{3, 5, 5, 3}, []int{7, 7}))             // 2
}
