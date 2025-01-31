/**
 * @Author: Gong Yanhui
 * @Description: 541. 反转字符串 II
 * @Date: 2025-01-31 14:50
 */

package main

import "fmt"

func reverseStr(s string, k int) string {
	var b = []byte(s)
	for i := 0; i < len(s); i += 2 * k {
		if i+k < len(s) {
			b = reverseStrRange541(b, i, i+k-1)
		} else {
			b = reverseStrRange541(b, i, len(s)-1)
		}
	}
	return string(b)
}

// 反转字符串 s 的 [begin, end] 区间的字符
func reverseStrRange541(str []byte, begin, end int) []byte {
	for begin < end {
		str[begin], str[end] = str[end], str[begin]
		begin++
		end--
	}
	return str
}

func main() {
	fmt.Println(string(reverseStrRange541([]byte("abcdef"), 1, 3)))  // acbdef
	fmt.Println(string(reverseStrRange541([]byte("abcdefg"), 0, 2))) // aedcbf

	// s = "abcdefg", k = 2
	println(reverseStr("abcdefg", 2)) // bacdfeg

	// s = "abcd", k = 2
	println(reverseStr("abcd", 2)) // bacd
}
