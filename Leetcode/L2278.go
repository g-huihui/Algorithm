/**
 * @Author: Gong Yanhui
 * @Description: 2278. 字母在字符串中的百分比
 * @Date: 2025-03-31 20:29
 */

package main

func percentageLetter(s string, letter byte) int {
	var size = len(s)
	if size == 0 {
		return 0
	}
	var count = 0
	for i := 0; i < size; i++ {
		if s[i] == letter {
			count++
		}
	}

	return (count * 100) / size
}
