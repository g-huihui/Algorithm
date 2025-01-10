/**
 * @Author: Gong Yanhui
 * @Description: 1768. 交替合并字符串
 * @Date: 2025-01-10 22:58
 */

package main

func mergeAlternately(word1 string, word2 string) string {
	var res string
	var index1, index2 int
	for index1 < len(word1) && index2 < len(word2) {
		res += string(word1[index1]) + string(word2[index2])
		index1++
		index2++
	}
	if index1 < len(word1) {
		res += word1[index1:]
	}
	if index2 < len(word2) {
		res += word2[index2:]
	}

	return res
}
