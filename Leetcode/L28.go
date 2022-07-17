/**
 * @Author: Gong Yanhui
 * @Description: 28. 实现 strStr()
 * @Date: 2022-07-17 17:22
 */

package main

func strStr(haystack string, needle string) int {
	n := len(haystack)
	m := len(needle)
	for i := 0; i+m <= n; i++ {
		if haystack[i] != needle[0] {
			continue
		}
		var flag = true
		for j := 0; j < m; j++ {
			if haystack[i+j] != needle[j] {
				flag = false
				break
			}
		}
		if flag {
			return i
		}
	}
	return -1
}
