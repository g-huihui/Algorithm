/**
 * @Author: Gong Yanhui
 * @Description: 第k个排列
 * @Date: 2024-12-22 17:43
 */

package main

import "fmt"

/**
1. 题目描述
给定参数n 从1到n会有n个整数: 1,2,3,…,n 这n个数字共有n!种排列 按大小顺序升序列出所有排列的情况
当n=3时有如下排列: 123, 132, 213, 231, 312, 321
给定n和k 返回第k个排列

2. 输入描述
n	给定的范围[1-n]
k	输出排在第k位置的数字

3. 输入输出
input : 3 3
output: 213
input : 2 2
output: 21
*/

func H5(n, k int) string {
	// 生成所有集合
	var result []string
	// 标记是否使用
	var used = make([]bool, n)

	backtrack5(&result, n, used, "")

	return result[k-1]
}

// 使用回溯算法生成所有结果
func backtrack5(result *[]string, n int, used []bool, cur string) {
	if len(cur) == n {
		*result = append(*result, cur)
		return
	}

	for i := 0; i < n; i++ {
		if !used[i] {
			used[i] = true
			backtrack5(result, n, used, fmt.Sprintf(cur+"%d", i+1))
			used[i] = false
		}
	}
}

func main() {
	println(H5(3, 3)) // 213
	println(H5(2, 2)) // 21
}
