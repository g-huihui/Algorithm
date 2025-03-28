/**
 * @Author: Gong Yanhui
 * @Description: 2716. 最小化字符串长度
 * @Date: 2025-03-28 22:03
 */

package main

func minimizedStringLength(s string) int {
	var m = make(map[uint8]struct{})

	for i := 0; i < len(s); i++ {
		m[s[i]] = struct{}{}
	}

	return len(m)
}

func main() {
	println(minimizedStringLength("aaabc"))  // 3
	println(minimizedStringLength("cbbd"))   // 3
	println(minimizedStringLength("dddaaa")) // 2
}
