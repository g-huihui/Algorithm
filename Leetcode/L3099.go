/**
 * @Author: Gong Yanhui
 * @Description: 3099. 哈沙德数
 * @Date: 2024-07-03 17:45
 */

package main

func sumOfTheDigitsOfHarshadNumber(x int) int {
	var res int

	var tmp = x
	for tmp != 0 {
		res += tmp % 10
		tmp /= 10
	}

	if x%res != 0 {
		return -1
	}

	return res
}
