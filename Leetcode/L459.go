/**
 * @Author: Gong Yanhui
 * @Description: 459. 重复的子字符串
 * @Date: 2025-01-11 11:24
 */

package main

func repeatedSubstringPattern(s string) bool {
	// 创建一个临时拷贝的字符串
	var tmp = s
	// 逐个将字符串的第一个字符移到最后 判断是否与原字符串相等 只需要判断一半即可
	for i := 0; i < len(s)/2; i++ {
		tmp = tmp[1:] + string(tmp[0])
		if tmp == s {
			return true
		}
	}
	return false
	// 该方法可以执行通过 但是时间复杂度较高 在字符串长度过长时不好
}
