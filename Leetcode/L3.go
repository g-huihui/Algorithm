/**
 * @Author: Gong Yanhui
 * @Description: 3. 无重复字符的最长子串
 * @Date: 2024-11-30 16:21
 */

package main

func lengthOfLongestSubstring(s string) int {
	var maxDistance int
	var mFlag = make(map[byte]int)
	for i := 0; i < len(s); i++ {
		_, ok := mFlag[s[i]]
		if !ok {
			mFlag[s[i]] = i
			if len(mFlag) > maxDistance {
				maxDistance = len(mFlag)
			}
		} else {
			// 找到重复字符的位置 删除重复字符之前的字符
			for j := i - len(mFlag); j < i; j++ {
				delete(mFlag, s[j])
				if s[j] == s[i] {
					break
				}
			}
			mFlag[s[i]] = i
		}
	}

	return maxDistance
}

func main() {
	println(lengthOfLongestSubstring("abcabcbb")) // 3
	println(lengthOfLongestSubstring("bbbbb"))    // 1
	println(lengthOfLongestSubstring("pwwkew"))   // 3
}
