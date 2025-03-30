/**
 * @Author: Gong Yanhui
 * @Description: 2109. 向字符串添加空格
 * @Date: 2025-03-30 10:53
 */

package main

import "bytes"

// 使用 "+"字符串拼接会超时 使用 bytes.Buffer 优化后通过
func addSpaces(s string, spaces []int) string {

	if len(spaces) == 1 {
		return s[:spaces[0]] + " " + s[spaces[0]:]
	}

	var result bytes.Buffer
	for i := 0; i < len(spaces); i++ {
		var cur string
		if i == 0 {
			cur = s[:spaces[i]]
			result.WriteString(cur)
		} else if i == len(spaces)-1 {
			cur = " " + s[spaces[i-1]:spaces[i]] + " " + s[spaces[i]:]
			result.WriteString(cur)
		} else {
			cur = s[spaces[i-1]:spaces[i]]
			result.WriteString(" " + cur)
		}
	}

	return result.String()
}
