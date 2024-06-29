/**
 * @Author: Gong Yanhui
 * @Description: 2710. 移除字符串中的尾随零
 * @Date: 2024-06-29 16:21
 */

package main

func removeTrailingZeros(num string) string {
	for len(num) > 0 && num[len(num)-1] == '0' {
		num = num[:len(num)-1]
	}
	return num
}

func main() {
	println(removeTrailingZeros("51230100"))
	println(removeTrailingZeros("123"))
}
