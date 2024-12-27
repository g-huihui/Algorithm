/**
 * @Author: Gong Yanhui
 * @Description: 3159. 查询数组中元素的出现位置
 * @Date: 2024-12-27 15:40
 */

package main

func occurrencesOfElement(nums []int, queries []int, x int) []int {

	var size int
	var m = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if nums[i] == x {
			size++
			m[size] = i
		}
	}

	var res []int
	for _, n := range queries {
		tmp, ok := m[n]
		if ok {
			res = append(res, tmp)
		} else {
			res = append(res, -1)
		}
	}

	return res
}
