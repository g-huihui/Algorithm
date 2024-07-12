/**
 * @Author: Gong Yanhui
 * @Description: 2974. 最小数字游戏
 * @Date: 2024-07-12 20:49
 */

package main

import "sort"

func numberGame(nums []int) []int {
	sort.Ints(nums)

	length := len(nums)
	var res = make([]int, length)

	for i := 0; i < length; i += 2 {
		res[i] = nums[i+1]
		res[i+1] = nums[i]
	}

	return res
}
