/**
 * @Author: Gong Yanhui
 * @Description: 1. 两数之和
 * @Date: 2022-07-10 22:45
 */

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	var nums = []int{1, 2, 3, 4, 5, 6}

	fmt.Println(twoSum(nums, 3))
}
