/**
 * @Author: Gong Yanhui
 * @Description: 2535. 数组元素和与数字和的绝对差
 * @Date: 2024-09-26 14:39
 */

package main

func differenceOfSum(nums []int) int {
	var totalSum int // 数组元素和
	var totalNum int // 数字和
	for _, num := range nums {
		totalSum += num
		totalNum += getSum(num)
	}
	return abs(totalSum - totalNum)
}

func getSum(num int) int {
	var sum int
	for num != 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}
