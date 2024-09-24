/**
 * @Author: Gong Yanhui
 * @Description: 2207. 字符串中最多数目的子序列
 * @Date: 2024-09-24 15:11
 */

package main

func maximumSubsequenceCount(text string, pattern string) int64 {
	var result, prev, next int64
	for _, ch := range text {
		if ch == rune(pattern[1]) {
			result += prev
			next++
		}
		if ch == rune(pattern[0]) {
			prev++
		}
	}
	if prev > next {
		result += prev
	} else {
		result += next
	}

	return result
}

func main() {
	println(maximumSubsequenceCount("abdcdbc", "ac")) // 4
	println(maximumSubsequenceCount("aabb", "ab"))    // 6
}
