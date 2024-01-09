/**
 * @Author: Gong Yanhui
 * @Description: 2744. 最大字符串配对数目
 * @Date: 2024-01-09 16:31
 */

package main

func maximumNumberOfStringPairs(words []string) int {
	length := len(words)
	count := 0
	flag := make([]bool, length)
	for i := 0; i < length-1; i++ {
		if flag[i] {
			continue
		}
		for j := i + 1; j < length; j++ {
			if flag[j] {
				continue
			}
			if isReverse(words[i], words[j]) {
				count++
			}
		}
	}
	return count
}

func isReverse(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	length := len(str1)
	for i := 0; i < length; i++ {
		if str1[i] != str2[length-i-1] {
			return false
		}
	}

	return true
}
