/**
 * @Author: Gong Yanhui
 * @Description: 3083. 字符串及其反转中是否存在同一子字符串
 * @Date: 2024-12-26 19:59
 */

package main

func isSubstringPresent(s string) bool {
	if len(s) < 2 {
		return false
	}
	var m = make(map[rune]map[rune]struct{})
	for i := 1; i < len(s); i++ {
		// 先放入新的字符
		if _, ok := m[rune(s[i-1])]; !ok {
			m[rune(s[i-1])] = make(map[rune]struct{})
		}
		m[rune(s[i-1])][rune(s[i])] = struct{}{}
		// 再检查当前字符是否在之前的字符中
		if _, ok := m[rune(s[i])]; ok {
			if _, okIn := m[rune(s[i])][rune(s[i-1])]; okIn {
				return true
			}
		}
	}

	return false
}
