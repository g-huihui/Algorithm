/**
 * @Author: Gong Yanhui
 * @Description: 补种未成活胡杨
 * @Date: 2024-12-16 20:28
 */

package main

/**
1. 题目描述
某沙漠新种植N棵胡杨排成一排,一个月后又M棵胡杨未能成活,现可补种胡杨K课,请问如何补种(只能补种不能新种)可以得到最多连续胡杨树?

2. 输入描述
N 总种植数量
M 未成活胡杨数量
M 个空格分割的数,按编号从小到大排序
K 最多可以补种的数量

3. 输入输出
input : 5 [2,4] 1
output: 3
input : 10 [2,4,7] 1
output: 6
*/

func H2(N int, M []int, K int) int {

	if K >= len(M) {
		return len(M)
	}

	// 用数组代表一排树 0成活 1未成活
	var tree = make([]int, N+1)
	for i := 0; i < len(M); i++ {
		tree[M[i]] = 1
	}

	// 滑动窗口
	var maxNum = M[K]
	for i := K; i < len(M); i++ {
		var tmp int
		if i == len(M)-1 {
			tmp = N - M[i-K] + 1
		} else {
			tmp = M[i] - M[i-K] - 1
		}
		if maxNum < tmp {
			maxNum = tmp
		}
	}

	return maxNum - 1
}

// 没有经过大量数据验证 不确定是否完全正确
func main() {
	println(H2(5, []int{2, 4}, 1))     // 3
	println(H2(10, []int{2, 4, 7}, 1)) // 6
}
