/**
 * @Author: Gong Yanhui
 * @Description: 3146. 两个字符串的排列差
 * @Date: 2024-08-24 22:41
 */

package main

func findPermutationDifference(s string, t string) int {
	var sMap = make(map[byte]int)
	for index, b := range s {
		sMap[byte(b)] = index
	}

	var res int
	for index, b := range t {
		var tmp = index - sMap[byte(b)]
		if tmp > 0 {
			res += tmp
		} else {
			res -= tmp
		}
	}

	return res
}
