/**
 * @Author: Gong Yanhui
 * @Description: 1502. 判断能否形成等差数列
 * @Date: 2025-01-23 15:44
 */

package main

import "sort"

func canMakeArithmeticProgression(arr []int) bool {
	var size = len(arr)
	if size <= 2 {
		return true
	}

	sort.Ints(arr)
	var diff = arr[1] - arr[0]
	for i := 2; i < size; i++ {
		if arr[i]-arr[i-1] != diff {
			return false
		}
	}

	return true
}
