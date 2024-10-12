/**
 * @Author: Gong Yanhui
 * @Description: 3158. 求出出现两次数字的 XOR 值
 * @Date: 2024-10-12 23:21
 */

package main

func duplicateNumbersXOR(nums []int) int {
	var res int
	var m = make(map[int]struct{})
	for _, num := range nums {
		if _, ok := m[num]; ok {
			res ^= num
		} else {
			m[num] = struct{}{}
		}
	}
	return res
}
