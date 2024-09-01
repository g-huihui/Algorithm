/**
 * @Author: Gong Yanhui
 * @Description: 1450. 在既定时间做作业的学生人数
 * @Date: 2024-09-01 11:22
 */

package main

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	var result int
	var length = len(startTime)
	for i := 0; i < length; i++ {
		if startTime[i] <= queryTime && endTime[i] >= queryTime {
			result++
		}
	}
	return result
}
