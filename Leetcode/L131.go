/**
 * @Author: Gong Yanhui
 * @Description: 131. 分割回文串
 * @Date: 2025-03-02 13:53
 */

package main

func partition(s string) [][]string {
	var res [][]string
	var path []string

	dfs131(s, 0, path, &res)

	return res
}

func dfs131(s string, start int, path []string, res *[][]string) {
	if start == len(s) {
		*res = append(*res, append([]string{}, path...))
		return
	}

	// 从start开始 尝试每一种可能 start为每次分割的点
	for i := start; i < len(s); i++ {
		if !isPalindrome(s[start : i+1]) {
			continue
		}
		path = append(path, s[start:i+1])
		dfs131(s, i+1, path, res)
		path = path[:len(path)-1]
	}
}
