/**
 * @Author: Gong Yanhui
 * @Description: 377. 组合总和 Ⅳ
 * @Date: 2024-04-21 21:57
 */

package main

//
// combinationSum4
//  @Description: 给你一个由 不同 整数组成的数组 nums 和一个目标整数 target 请你从 nums 中找出并返回总和为 target 的元素组合的个数
//  @param nums
//  @param target
//  @return int
//
func combinationSum4(nums []int, target int) int {
	var dp = make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i >= num {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}

/*
// 按照之前的解法会超时
func combinationSum4(nums []int, target int) int {
	var res int
	sort.Ints(nums)
	backtrack4(nums, target, 0, &res)
	return res
}
func backtrack4(nums []int, target int, index int, res *int) {
	if target == 0 {
		*res++
		return
	}
	for i := index; i < len(nums); i++ {
		if target < nums[i] {
			break
		}
		backtrack4(nums, target-nums[i], index, res)
	}
}
*/

func main() {
	var nums = []int{1, 2, 3}
	var target = 4
	println(combinationSum4(nums, target)) // 7

	nums = []int{2, 1, 3}
	target = 35
	println(combinationSum4(nums, target)) // 1132436852
}
