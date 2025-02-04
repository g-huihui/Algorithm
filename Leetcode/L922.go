/**
 * @Author: Gong Yanhui
 * @Description: 922. 按奇偶排序数组 II
 * @Date: 2025-02-04 14:23
 */

package main

import "fmt"

func sortArrayByParityII(nums []int) []int {
	var size = len(nums)
	if size == 0 {
		return nums
	}

	var even = 0 // 偶数索引
	var odd = 1  // 奇数索引

	for even < size || odd < size {
		// 找到偶数索引上的奇数
		for even < size {
			if nums[even]%2 == 1 {
				break
			}
			even += 2
		}

		// 找到奇数索引上的偶数
		for odd < size {
			if nums[odd]%2 == 0 {
				break
			}
			odd += 2
		}

		if even < size && odd < size {
			nums[even], nums[odd] = nums[odd], nums[even]
		}
	}

	return nums
}

func main() {
	fmt.Println(sortArrayByParityII([]int{4, 2, 5, 7}))
	fmt.Println(sortArrayByParityII([]int{2, 3}))

	fmt.Println(sortArrayByParityII([]int{3, 1, 4, 2}))
}
