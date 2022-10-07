/**
 * @Author: Gong Yanhui
 * @Description: 136. 只出现一次的数字
 * @Date: 2022-10-07 21:50
 */

package main

import "fmt"

func singleNumber(nums []int) int {
	var result = 0
	for i := 0; i < len(nums); i++ {
		result = result ^ nums[i]
	}
	return result
}

func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
}
