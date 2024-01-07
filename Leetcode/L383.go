/**
 * @Author: Gong Yanhui
 * @Description: 383. 赎金信
 * @Date: 2024-01-07 23:05
 */

package main

func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[rune]int)
	for _, v := range magazine {
		m[v]++
	}
	for _, v := range ransomNote {
		if m[v] == 0 {
			return false
		}
		m[v]--
	}
	return true
}
