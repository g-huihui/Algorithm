/**
 * @Author: Gong Yanhui
 * @Description: 2187. 完成旅途的最少时间
 * @Date: 2024-10-05 14:57
 */

package main

import (
	"math"
)

// 二分查找法
func minimumTime(time []int, totalTrips int) int64 {
	var left, right = 0, math.MinInt
	for _, num := range time {
		if num*totalTrips > right {
			right = num * totalTrips
		}
	}

	for left < right {
		mid := left + (right-left)/2
		if checkIsOk(time, totalTrips, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return int64(left)
}

func checkIsOk(time []int, totalTrips int, mid int) bool {
	var count int
	for _, num := range time {
		count += mid / num
	}
	return count >= totalTrips
}

// 每次走一步 超时
func minimumTime2(time []int, totalTrips int) int64 {
	var length = len(time)
	var resTime int64
	var count int
	var numFlag = make([]int, length)
	copy(numFlag, time)
	for count < totalTrips {
		resTime++
		for i := 0; i < length; i++ {
			numFlag[i]--
			if numFlag[i] == 0 {
				count++
				numFlag[i] = time[i]
				if count == totalTrips {
					return resTime
				}
			}
		}
	}
	return resTime
}

// 有问题
func minimumTime1(time []int, totalTrips int) int64 {
	var length = len(time)
	var minStep = math.MinInt
	for i := 0; i < length; i++ {
		if time[i] < minStep {
			minStep = time[i]
		}
	}
	var resTime int64
	var count int
	var numFlag = make([]int, length)
	copy(numFlag, time)
	for count < totalTrips {
		resTime += int64(minStep)
		for i := 0; i < length; i++ {
			if numFlag[i] >= minStep {
				numFlag[i] -= minStep
				if numFlag[i] == 0 {
					count++
					numFlag[i] = time[i]
					if count == totalTrips {
						return resTime
					}
				}
			} else {
				count++
				numFlag[i] = numFlag[i] - minStep + time[i]
				if count == totalTrips { // 如果减多了要回退
					return resTime - int64(numFlag[i]+minStep-time[i])
				}
			}
		}
	}

	return resTime
}

func main() {
	println(minimumTime([]int{1, 2, 3}, 5))     // 3
	println(minimumTime([]int{2}, 1))           // 2
	println(minimumTime([]int{9, 3, 10, 5}, 2)) // 5
}
