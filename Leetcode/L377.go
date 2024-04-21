/**
 * @Author: Gong Yanhui
 * @Description: 377. 组合总和 Ⅳ
 * @Date: 2024-04-21 21:57
 */

package main

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
