/**
 * @Author: Gong Yanhui
 * @Description: 680. 验证回文串 II
 * @Date: 2025-02-03 14:11
 */

package main

func validPalindrome(s string) bool {
	return f680(s, false)
}

func f680(str string, isDeleted bool) bool {
	var size = len(str)

	if size <= 1 {
		return true
	}

	if str[0] == str[size-1] {
		return f680(str[1:size-1], isDeleted)
	} else {
		if isDeleted {
			return false
		}
		return f680(str[1:size], true) || f680(str[:size-1], true)
	}
}
