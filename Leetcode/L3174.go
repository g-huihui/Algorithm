/**
 * @Author: Gong Yanhui
 * @Description: 3174. 清除数字
 * @Date: 2024-09-05 20:21
 */

package main

import (
	"unicode"
)

func clearDigits(s string) string {
	var res []byte
	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			res = res[:len(res)-1]
		} else {
			res = append(res, s[i])
		}
	}

	return string(res)
}
