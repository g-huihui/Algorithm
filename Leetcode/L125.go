/**
 * @Author: Gong Yanhui
 * @Description: 125. 验证回文串
 * @Date: 2022-08-20 11:37
 */

package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	if len(s) < 2 {
		return true
	}
	left, right := 0, len(s)-1
	for left <= right {
		if !isABC(s[left]) {
			left++
			continue
		}
		if !isABC(s[right]) {
			right--
			continue
		}
		if strings.EqualFold(string(s[left]), string(s[right])) {
			left++
			right--
			continue
		}
		return false
	}
	return true
}

func isABC(s byte) bool {
	if (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z') || (s >= '0' && s <= '9') {
		return true
	}
	return false
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("0P"))
}
