/**
 * @Author: Gong Yanhui
 * @Description: 565. 数组嵌套
 * @Date: 2022-07-17 11:55
 */

package main

import "fmt"

// 有向图 每条路线都是一个环状 只需遍历数组 找到最大的环即可
func arrayNesting(nums []int) int {
	var max int = 0
	vis := make([]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		cur := 0
		for !vis[i] {
			vis[i] = true
			i = nums[i]
			cur++
		}
		if max < cur {
			max = cur
		}
	}
	return max
}

func main() {
	var nums []int = []int{5, 4, 0, 3, 1, 6, 2}
	fmt.Println(arrayNesting(nums))
}
