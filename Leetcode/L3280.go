/**
 * @Author: Gong Yanhui
 * @Description: 3280. 将日期转换为二进制表示
 * @Date: 2025-01-03 11:00
 */

package main

import (
	"strconv"
	"strings"
)

// 力扣 十进制转二进制
func binary(x int) string {
	var s []byte
	for ; x != 0; x >>= 1 {
		s = append(s, '0'+byte(x&1))
	}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}

func convertDateToBinary(date string) string {
	s := strings.Split(date, "-")
	year, _ := strconv.Atoi(s[0])
	month, _ := strconv.Atoi(s[1])
	day, _ := strconv.Atoi(s[2])
	return binary(year) + "-" + binary(month) + "-" + binary(day)
}
