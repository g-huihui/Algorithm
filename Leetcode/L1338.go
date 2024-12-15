/**
 * @Author: Gong Yanhui
 * @Description: 1338. 数组大小减半
 * @Date: 2024-12-15 16:19
 */

package main

import "sort"

func minSetSize(arr []int) int {

	// 记录每个数字出现的次数
	var m = make(map[int]int)
	for _, num := range arr {
		m[num]++
	}

	// 将次数从大到小排序
	var count = make([]int, 0, len(m))
	for _, num := range m {
		count = append(count, num)
	}
	// 这是一种倒序的方式!! 先将切片转换IntSlice为排序专用的格式
	sort.Sort(sort.Reverse(sort.IntSlice(count)))

	var res, total int
	for _, num := range count {
		total += num
		res++
		if total*2 >= len(arr) {
			break
		}
	}

	return res
}
