/**
 * @Author: Gong Yanhui
 * @Description: 2269. 找到一个数字的 K 美丽值
 * @Date: 2025-03-10 22:12
 */

package main

import (
	"fmt"
	"strconv"
)

func divisorSubstrings(num int, k int) int {
	var count int
	var n = fmt.Sprint(num)
	for i := 0; i <= len(n)-k; i++ {
		cur, _ := strconv.Atoi(n[i : i+k])
		if cur != 0 && num%cur == 0 {
			count++
		}
	}

	return count
}
