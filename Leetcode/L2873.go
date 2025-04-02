/**
 * @Author: Gong Yanhui
 * @Description: 2873. 有序三元组中的最大值 I
 * @Date: 2025-04-02 19:20
 */

package main

func maximumTripletValue(nums []int) int64 {
	var res int64
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				var cur = int64((nums[i] - nums[j]) * nums[k])
				if cur > res {
					res = cur
				}
			}
		}
	}

	return res
}
