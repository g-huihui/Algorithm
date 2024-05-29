/**
 * @Author: Gong Yanhui
 * @Description: 274. H 指数
 * @Date: 2024-05-29 20:16
 */

package main

import "sort"

func hIndex(citations []int) int {
	var h int
	sort.Ints(citations)
	for i := len(citations) - 1; i >= 0 && citations[i] > h; i-- {
		h++
	}
	return h
}

// [0, 1, 3, 5, 6]	3
// [1, 1, 3]		1
