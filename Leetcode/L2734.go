/**
 * @Author: Gong Yanhui
 * @Description: 2734. 执行子串操作后的字典序最小字符串
 * @Date: 2024-06-27 20:21
 */

package main

func smallestString(s string) string {

	var length = len(s)
	var byteStr = []byte(s)
	for i := 0; i < length; i++ {
		if byteStr[i] > 'a' {
			for j := i; j < length; j++ {
				if byteStr[j] > 'a' {
					byteStr[j]--
				} else {
					break
				}
			}
			return string(byteStr)
		}
	}

	byteStr[length-1] = 'z'
	return string(byteStr)
}

func main() {
	println(smallestString("cbabc"))    // "baabc"
	println(smallestString("acbbc"))    // "abaab"
	println(smallestString("leetcode")) // "kddsbncd"
	println(smallestString("aaaaa"))    // "aaaaz"
}
