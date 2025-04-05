/**
 * @Author: Gong Yanhui
 * @Description: 1863. 找出所有子集的异或总和再求和
 * @Date: 2025-04-05 10:42
 */

package main

func subsetXORSum(nums []int) int {
	return dfs1863(0, 0, nums)
}

// 代表索引递归到index之前的值res
func dfs1863(res, index int, nums []int) int {
	// 结束
	if index == len(nums) {
		return res
	}

	// 递归当前值和下一个值的和
	return dfs1863(res^nums[index], index+1, nums) + dfs1863(res, index+1, nums)
}
