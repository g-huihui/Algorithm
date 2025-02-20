/**
 * @Author: Gong Yanhui
 * @Description: 2595. 奇偶位数
 * @Date: 2025-02-20 22:51
 */

package main

/** 二进制转换
假设 n = 5 -> 101
通过 n & 1 取二进制最低位
然后 n >> 1 把n向右移动1位
通过这种方式循环实现十进制转二进制 直到n==0未知
*/

func evenOddBit(n int) []int {
	var even, odd int
	var index int
	for n != 0 {
		if (n & 1) == 1 {
			if index%2 == 0 {
				even++
			} else {
				odd++
			}
		}
		n = n >> 1
		index++
	}

	return []int{even, odd}
}
