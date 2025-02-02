/**
 * @Author: Gong Yanhui
 * @Description: 598. 区间加法 II
 * @Date: 2025-02-02 15:50
 */

package main

func maxCount(m int, n int, ops [][]int) int {
	for _, ints := range ops {
		m = min(m, ints[0])
		n = min(n, ints[1])
	}

	return m * n
}
