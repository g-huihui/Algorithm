/**
 * @Author: Gong Yanhui
 * @Description: 3254. 长度为 K 的子数组的能量值 I
 * @Date: 2024-11-06 12:00
 */

package main

import "fmt"

// 暴力破解
func resultsArray(nums []int, k int) []int {
	var res []int
	for i := 0; i < len(nums)-k+1; i++ {
		var j int
		for j = 1; j < k; j++ {
			if nums[i+j]-nums[i+j-1] != 1 {
				break
			}
		}
		if j == k {
			res = append(res, nums[i+j-1])
		} else {
			res = append(res, -1)
		}
	}

	return res
}

func main() {
	fmt.Println(resultsArray([]int{1, 2, 3, 4, 3, 2, 5}, 3)) // [3,4,-1,-1,-1]
}
