/**
 * @Author: Gong Yanhui
 * @Description: 2414. 最长的字母序连续子字符串的长度
 * @Date: 2024-09-21 23:55
 */

package main

func longestContinuousSubstring(s string) int {
	var maxRes = 1
	var begin int
	var length = len(s)
	for i := begin; i < length-1; i++ {
		if s[i]+1 != s[i+1] {
			var tmp = i - begin + 1
			if tmp > maxRes {
				maxRes = tmp
			}
			begin = i + 1
		}
		if i == length-2 {
			var tmp = length - begin
			if tmp > maxRes {
				maxRes = tmp
			}
		}
	}
	return maxRes
}

func main() {
	println(longestContinuousSubstring("abacaba")) //2
	println(longestContinuousSubstring("abcde"))   //5
}
