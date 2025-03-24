/**
 * @Author: Gong Yanhui
 * @Description: 2255. 统计是给定字符串前缀的字符串数目
 * @Date: 2025-03-24 12:18
 */

package main

import "strings"

// 方法1 使用go内置库
// 执行用时 1ms 内存消耗5.15MB
func countPrefixes1(words []string, s string) int {
	var count int
	for _, w := range words {
		if strings.HasPrefix(s, w) {
			count++
		}
	}

	return count
}

// 方法2 使用map缓存 复杂度O(len(words) + len(s))
// 执行用时 0ms 内存消耗6.01MB
func countPrefixes(words []string, s string) int {
	var count int
	var cache = make(map[string]int, len(words))
	for _, w := range words {
		cache[w]++
	}

	var size = len(s)
	for i := 0; i < size; i++ {
		if v, ok := cache[s[:i+1]]; ok {
			count += v
		}
	}

	return count
}
