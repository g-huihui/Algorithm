/**
 * @Author: Gong Yanhui
 * @Description: LCP 61. 气温变化趋势
 * @Date: 2024-06-21 14:57
 */

package main

func temperatureTrend(temperatureA []int, temperatureB []int) int {

	var maxDays int // 最大连续天数
	var day = 0     // 当前连续天数

	for i := 1; i < len(temperatureA); i++ {
		if (temperatureA[i] > temperatureA[i-1] && temperatureB[i] > temperatureB[i-1]) ||
			(temperatureA[i] < temperatureA[i-1] && temperatureB[i] < temperatureB[i-1]) ||
			(temperatureA[i] == temperatureA[i-1] && temperatureB[i] == temperatureB[i-1]) {
			day++
		} else {
			if day > maxDays {
				maxDays = day
			}
			day = 0
		}
	}
	if day > maxDays {
		maxDays = day
	}

	return maxDays
}

func main() {
	temperatureA1 := []int{21, 18, 18, 18, 31}
	temperatureB1 := []int{34, 32, 16, 16, 17}
	println(temperatureTrend(temperatureA1, temperatureB1)) // 2

	temperatureA2 := []int{5, 10, 16, -6, 15, 11, 3}
	temperatureB2 := []int{16, 22, 23, 23, 25, 3, -16}
	println(temperatureTrend(temperatureA2, temperatureB2)) // 3

	temperatureA3 := []int{1, -15, 3, 14, -1, 4, 35, 36}
	temperatureB3 := []int{-15, 32, 20, 9, 33, 4, -1, -5}
	println(temperatureTrend(temperatureA3, temperatureB3)) // 0

	temperatureA4 := []int{-14, 7, -19, 9, 13, 40, 19, 15, -18}
	temperatureB4 := []int{3, 16, 28, 32, 25, 12, 13, -6, 4}
	println(temperatureTrend(temperatureA4, temperatureB4)) // 1
}
