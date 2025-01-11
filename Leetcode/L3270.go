/**
 * @Author: Gong Yanhui
 * @Description: 3270. 求出数字答案
 * @Date: 2025-01-11 09:48
 */

package main

import "math"

func generateKey(num1 int, num2 int, num3 int) int {
	var key int
	for i := 0; i < 4 && num1 > 0 && num2 > 0 && num3 > 0; i++ {
		key += min3270(num1%10, num2%10, num3%10) * int(math.Pow(10, float64(i)))
		num1 /= 10
		num2 /= 10
		num3 /= 10
	}

	return key
}

func min3270(x, y, z int) int {
	return min(min(x, y), z)
}
