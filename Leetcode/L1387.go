/**
 * @Author: Gong Yanhui
 * @Description: 1387. 将整数按权重排序
 * @Date: 2024-12-22 12:41
 */

package main

import "sort"

func getKth(lo int, hi int, k int) int {
	var m = make([][]int, hi-lo+1)
	for i := lo; i <= hi; i++ {
		m[i-lo] = []int{i, get1387(i)}
	}
	sort.Slice(m, func(i, j int) bool {
		if m[i][1] < m[j][1] {
			return true
		} else if m[i][1] == m[j][1] {
			return m[i][0] < m[j][0]
		}
		return false
	})

	return m[k-1][0]
}

// 根据num获取对应的权重
func get1387(num int) int {
	var count int
	for num != 1 {
		count++
		if num%2 == 0 {
			num /= 2
		} else {
			num = num*3 + 1
		}
	}
	return count
}

func main() {
	println(getKth(12, 15, 2)) // 13
	println(getKth(7, 11, 4))  // 7
}
