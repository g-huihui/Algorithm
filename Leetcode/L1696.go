/**
 * @Author: Gong Yanhui
 * @Description: 1696. 跳跃游戏 VI Shopee 面试题
 * @Date: 2025-03-31 20:25
 */

package main

// 动态规划 超时 72 / 73 个通过的测试用例
func maxResult(nums []int, k int) int {
	var size = len(nums)
	var dp = make([]int, len(nums))

	dp[0] = nums[0]
	var head, tail int //	定义指向队列的对头队尾指针

	for i := 1; i < size; i++ {
		dp[i] = nums[i] + getMax(dp, &head, &tail, i, k)
	}

	return dp[size-1]
}

// 滑动窗口 双指针
func getMax(arr []int, head, tail *int, index, k int) int {
	if index-*head > k {
		*head = index - k
	}
	// 先找到当前节点前范围内最大值
	var result = arr[*head]
	for i := *head + 1; i <= *tail; i++ {
		if arr[i] > result {
			*head = i
			result = arr[i]
		}
	}
	*tail = index
	return result
}

func main() {
	println(maxResult([]int{1, -1, -2, 4, -7, 3}, 2))          // 7
	println(maxResult([]int{10, -5, -2, 4, 0, 3}, 3))          // 17
	println(maxResult([]int{1, -5, -20, 4, -1, 3, -6, -3}, 2)) // 0
}
