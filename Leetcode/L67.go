/**
 * @Author: Gong Yanhui
 * @Description: 67. 二进制求和
 * @Date: 2022-07-17 22:32
 */

package main

import "strconv"

func addBinary(a string, b string) string {
	var res string
	lenA, lenB := len(a), len(b)
	maxLen := max(lenA, lenB)
	carry := 0
	for i := 0; i < maxLen; i++ {
		if i < lenA {
			carry += int(a[lenA-i-1] - '0')
		}
		if i < lenB {
			carry += int(b[lenB-i-1] - '0')
		}
		res = strconv.Itoa(carry%2) + res
		carry /= 2
	}
	if carry > 0 {
		res = "1" + res
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
