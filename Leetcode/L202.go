/**
 * @Author: Gong Yanhui
 * @Description: 202. 快乐数
 * @Date: 2024-12-01 18:35
 */

package main

func isHappy(n int) bool {
	var slow, fast = n, bitSquareSum(n)
	for slow != fast {
		slow = bitSquareSum(slow)
		fast = bitSquareSum(fast)
		fast = bitSquareSum(fast)
	}

	return slow == 1
}

func bitSquareSum(n int) (res int) {
	for n != 0 {
		tmp := n % 10
		res += tmp * tmp
		n /= 10
	}
	return
}

func main() {
	println(isHappy(19)) // ture
	println(isHappy(2))  // false

	println(isHappy(10)) // ture
}
