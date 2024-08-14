/**
 * @Author: Gong Yanhui
 * @Description: 3152. 特殊数组 II
 * @Date: 2024-08-14 10:25
 */

package main

import "fmt"

func isArraySpecial2(nums []int, queries [][]int) []bool {
	var flags = make([]int, len(nums))
	for i := 1; i < len(flags); i++ {
		flags[i] = flags[i-1]
		if nums[i-1]%2 == nums[i]%2 {
			flags[i]++
		}
	}
	var res []bool
	for _, between := range queries {
		if flags[between[1]] == flags[between[0]] {
			res = append(res, true)
		} else {
			res = append(res, false)
		}
	}
	return res
}

// 超出时间限制
func isArraySpecial2Tmp(nums []int, queries [][]int) []bool {
	var flags = make([]int, len(nums))
	var res []bool
next:
	for _, between := range queries {
		for i := between[0] + 1; i <= between[1]; i++ {
			if flags[i-1] == 0 {
				flags[i-1] = checkOdd(nums[i-1])
			}
			if flags[i] == 0 {
				flags[i] = checkOdd(nums[i])
			}
			if flags[i-1] == flags[i] {
				res = append(res, false)
				continue next
			}
		}
		res = append(res, true)
	}

	return res
}

// 判断奇偶数 奇数返回1 偶数返回2
func checkOdd(num int) int {
	if num%2 == 1 {
		return 1
	}
	return 2
}

func main() {
	// nums = [3,4,1,2,6], queries = [[0,4]]
	fmt.Println(isArraySpecial2([]int{3, 4, 1, 2, 6}, [][]int{{0, 4}})) // false

	// nums = [4,3,1,6], queries = [[0,2],[2,3]]
	fmt.Println(isArraySpecial2([]int{4, 3, 1, 6}, [][]int{{0, 2}, {2, 3}})) // false true

	// num = [1,1], queries = [[0,1]]
	fmt.Println(isArraySpecial2([]int{1, 1}, [][]int{{0, 1}})) // false
}
