/**
 * @Author: Gong Yanhui
 * @Description: 2391. 收集垃圾的最少总时间
 * @Date: 2024-05-11 20:39
 */

package main

import "strings"

func garbageCollection(garbage []string, travel []int) int {
	var totalTime = 0
	var mFlag, pFlag, gFlag bool

	for i := len(garbage) - 1; i >= 0; i-- {
		if !mFlag && strings.Contains(garbage[i], "M") {
			for j := 0; j < i; j++ {
				totalTime += travel[j]
			}
			mFlag = true
		}
		if !pFlag && strings.Contains(garbage[i], "P") {
			for j := 0; j < i; j++ {
				totalTime += travel[j]
			}
			pFlag = true
		}
		if !gFlag && strings.Contains(garbage[i], "G") {
			for j := 0; j < i; j++ {
				totalTime += travel[j]
			}
			gFlag = true
		}
		totalTime += len(garbage[i])
	}

	return totalTime
}
