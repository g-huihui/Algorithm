/**
 * @Author: Gong Yanhui
 * @Description: 2766. 重新放置石块
 * @Date: 2024-07-24 12:05
 */

package main

import "sort"

func relocateMarbles(nums []int, moveFrom []int, moveTo []int) []int {

	var set = make(map[int]struct{})
	for _, num := range nums {
		set[num] = struct{}{}
	}

	for index, num := range moveFrom {
		delete(set, num)
		set[moveTo[index]] = struct{}{}
	}

	var res []int
	for num, _ := range set {
		res = append(res, num)
	}
	sort.Ints(res)

	return res
}
