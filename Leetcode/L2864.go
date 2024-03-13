/**
 * @Author: Gong Yanhui
 * @Description: 2864. 最大二进制奇数
 * @Date: 2024-03-13 21:25
 */

package main

import "strings"

func maximumOddBinaryNumber(s string) string {
	onesCount := 0
	for i, _ := range s {
		if s[i] == '1' {
			onesCount++
		}
	}
	if onesCount == 0 {
		return "0"
	}
	zeroCount := len(s) - onesCount
	onesCount -= 1
	var res string
	for i := 0; i < onesCount; i++ {
		res += "1"
	}
	for i := 0; i < zeroCount; i++ {
		res += "0"
	}
	res += "1"
	return res
}

func main() {
	println(maximumOddBinaryNumber("1000"))
}

// 力扣官方题解
func maximumOddBinaryNumber_(s string) string {
	cnt := strings.Count(s, "1")
	return strings.Repeat("1", cnt-1) + strings.Repeat("0", len(s)-cnt) + "1"
}
