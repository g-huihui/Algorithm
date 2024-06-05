/**
 * @Author: Gong Yanhui
 * @Description: 151. 反转字符串中的单词
 * @Date: 2024-06-05 14:37
 */

package main

import "strings"

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	reverseWord := reverseAndCleanSpace(words)
	return strings.Join(reverseWord, " ")
}

func reverseAndCleanSpace(arr []string) []string {
	var tmpArr []string
	for index := range arr {
		if arr[index] != "" {
			tmpArr = append(tmpArr, arr[index])
		}
	}

	length := len(tmpArr)
	for i := 0; i < length/2; i++ {
		tmpArr[i], tmpArr[length-i-1] = tmpArr[length-i-1], tmpArr[i]
	}

	return tmpArr
}

func main() {
	println(reverseWords("  hello world  "))
	println(reverseWords("a good   example"))
}
