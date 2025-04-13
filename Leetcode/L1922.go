/**
 * @Author: Gong Yanhui
 * @Description: 1922. 统计好数字的数目
 * @Date: 2025-04-13 13:01
 */

package main

func countGoodNumbers(n int64) int {
	var mod int64 = 1000000007

	// 偶数位 可能的值 var even = []int{0, 2, 4, 6, 8}
	// 长度为 n 的 偶数位有 (n+1)/2

	// 奇数位 可能的值 var odd = []int{1, 3, 5, 7}
	// 长度为 n 的 奇数位有 n/2

	// 即 长度为 n 的可能有 5^(n+1)/2 * 4^n/2

	// 封装求幂函数 x^y
	var mulFunc = func(x, y int64) int64 {
		var res int64 = 1
		var num = x // 底数
		for y > 0 {
			if y%2 == 1 {
				res = res * num % mod
			}
			num = num * num % mod
			y /= 2
		}
		return res
	}

	return int(mulFunc(5, (n+1)/2) * mulFunc(4, n/2) % mod)
}
