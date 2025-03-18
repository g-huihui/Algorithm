/**
 * @Author: Gong Yanhui
 * @Description: 2614. 对角线上的质数
 * @Date: 2025-03-18 17:26
 */

package main

func diagonalPrime(nums [][]int) int {
	var size = len(nums)
	var res int
	for i := 0; i < size; i++ {
		if nums[i][i] > res && isPrime(nums[i][i]) {
			res = nums[i][i]
		}
		if nums[i][size-i-1] > res && isPrime(nums[i][size-i-1]) {
			res = nums[i][size-i-1]
		}
	}

	return res
}

// 判断 n 是否是质数
func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n >= 2 // 1 不是质数
}
