/**
 * @Author: Gong Yanhui
 * @Description: 392. 判断子序列
 * @Date: 2024-06-05 15:12
 */

package main

func isSubsequence(s string, t string) bool {
	sLen := len(s)
	tLen := len(t)
	if sLen > tLen {
		return false
	}
	if s == t || sLen == 0 {
		return true
	}

	for sIndex, tIndex := 0, 0; tIndex < tLen; tIndex++ {
		if s[sIndex] == t[tIndex] {
			sIndex++
		}
		if sIndex == sLen {
			return true
		}
	}
	return false
}

func main() {
	println(isSubsequence("abc", "ahbgdc"))
	println(isSubsequence("axc", "ahbgdc"))
	println(isSubsequence("", "ahbgdc"))
}
