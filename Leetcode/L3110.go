/**
 * @Author: Gong Yanhui
 * @Description: 3110. 字符串的分数
 * @Date: 2025-03-16 10:57
 */

package main

func scoreOfString(s string) int {
	var res int
	for i := 0; i < len(s)-1; i++ {
		res += abs(int(s[i]) - int(s[i+1]))
	}
	return res
}
