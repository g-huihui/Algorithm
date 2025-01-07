/**
 * @Author: Gong Yanhui
 * @Description: 3019. 按键变更的次数
 * @Date: 2025-01-07 14:34
 */

package main

import "strings"

func countKeyChanges(s string) int {
	lower := strings.ToLower(s)
	var count int
	for i := 0; i < len(s)-1; i++ {
		if lower[i] != lower[i+1] {
			count++
		}
	}

	return count
}
