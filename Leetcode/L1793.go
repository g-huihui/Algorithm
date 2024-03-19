/**
 * @Author: Gong Yanhui
 * @Description: 1793. 好子数组的最大分数
 * @Date: 2024-03-19 23:03
 */

package main

import "fmt"

// 初次解法: 暴力解法 超时 时间复杂度O(n^2) 空间复杂度O(1)
func maximumScore1(nums []int, k int) int {
	var maxScore int
	for left := k; left >= 0; left-- {
		for right := k; right < len(nums); right++ {
			value := minInArr1(nums[left:right+1]) * (right - left + 1)
			if value > maxScore {
				maxScore = value
			}
		}
	}
	return maxScore
}

// 查找数组中的最小值
func minInArr1(nums []int) int {
	minNum := nums[0]
	for _, num := range nums {
		if num < minNum {
			minNum = num
		}
	}
	return minNum
}

// 第二次解法: 初次优化 超时
func maximumScore2(nums []int, k int) int {
	var maxScore int
	minInArr := nums[k]
	for left := k; left >= 0; left-- {
		if nums[left] < minInArr {
			minInArr = nums[left]
		}
		minInArrFor := minInArr
		for right := k; right < len(nums); right++ {
			if nums[right] < minInArrFor {
				minInArrFor = nums[right]
			}
			value := minInArrFor * (right - left + 1)
			if value > maxScore {
				maxScore = value
			}
		}
	}
	return maxScore
}

// 第三次解法: 优化 最终版本
func maximumScore(nums []int, k int) int {
	n := len(nums)
	left, right := k-1, k+1
	ans := 0
	for i := nums[k]; ; i-- {
		for left >= 0 && nums[left] >= i {
			left--
		}
		for right < n && nums[right] >= i {
			right++
		}
		// ans = max(ans, (right-left-1)*i)
		if (right-left-1)*i > ans {
			ans = (right - left - 1) * i
		}
		if left == -1 && right == n {
			break
		}
	}
	return ans
}

func main() {
	nums := []int{1, 4, 3, 7, 4, 5}
	k := 3
	fmt.Println(maximumScore(nums, k)) // 15 (1, 5)

	nums = []int{5, 5, 4, 5, 4, 1, 1, 1}
	k = 0
	fmt.Println(maximumScore(nums, k)) // 20 (0, 4)

	nums = []int{6569, 9667, 3148, 7698, 1622, 2194, 793, 9041, 1670, 1872}
	k = 5
	fmt.Println(maximumScore(nums, k)) // 9732 (0, 5)
}
