/**
 * @Author: Gong Yanhui
 * @Description: 2288. 价格减免
 * @Date: 2024-06-18 20:14
 */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func discountPrices(sentence string, discount int) string {
	split := strings.Split(sentence, " ")
	for i := 0; i < len(split); i++ {
		if split[i][0] == '$' {
			res := discountFunc(split[i][1:], discount)
			if res != "" {
				split[i] = res
			}
		}
	}
	return strings.Join(split, " ")
}

// 接收 num 返回 num*(1-discount/100) 保留小数点后两位
func discountFunc(numStr string, discount int) string {
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return ""
	}
	discountPrice := float64(num) * (1 - float64(discount)/100)
	return fmt.Sprintf("$%.2f", discountPrice)
}
