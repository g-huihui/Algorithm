/**
 * @Author: Gong Yanhui
 * @Description: 58. 最后一个单词的长度
 * @Date: 2022-07-17 19:03
 */

package main

func lengthOfLastWord(s string) int {
	var index = len(s) - 1
	for s[index] == ' ' {
		index--
	}
	var wordLength = 0
	for index >= 0 && s[index] != ' ' {
		index--
		wordLength++
	}
	return wordLength
}
