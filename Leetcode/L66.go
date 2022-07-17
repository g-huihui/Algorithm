/**
 * @Author: Gong Yanhui
 * @Description: 66. 加一
 * @Date: 2022-07-17 19:50
 */

package main

import "fmt"

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		// 一直往前找到不为9的数
		if digits[i] != 9 {
			digits[i]++
			// 把后面的数全部置零
			for j := i + 1; j < len(digits); j++ {
				digits[j] = 0
			}
			return digits
		}
	}
	//没有在上面的for里return 说明全部为9
	var res []int = make([]int, len(digits)+1)
	res[0] = 1
	return res
}

func main() {
	var digits []int = []int{8, 9, 9, 9}
	fmt.Println(plusOne(digits))
}
