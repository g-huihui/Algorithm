/**
 * @Author: Gong Yanhui
 * @Description: 2610. 转换二维数组
 * @Date: 2025-03-19 19:54
 */

package main

func findMatrix(nums []int) [][]int {
	var m = make(map[int]int) // 记录数组中每个数字出现的次数
	var size = 0              // 记录最大出现次数 作为二维数组的行数
	for _, num := range nums {
		m[num]++
		if m[num] > size {
			size = m[num]
		}
	}
	var res = make([][]int, size)
	for i := 0; i < size; i++ {
		for num, count := range m {
			if count > 0 {
				res[i] = append(res[i], num)
				m[num]--
			}
		}
	}

	return res
}
