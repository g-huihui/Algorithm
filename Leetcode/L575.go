/**
 * @Author: Gong Yanhui
 * @Description: 575. 分糖果
 * @Date: 2024-06-27 21:21
 */

package main

func distributeCandies(candyType []int) int {
	var t = make(map[int]struct{})
	for i := 0; i < len(candyType); i++ {
		t[candyType[i]] = struct{}{}
	}

	only := len(candyType) / 2
	typeCount := len(t)
	if only > typeCount {
		return typeCount
	}

	return only
}
