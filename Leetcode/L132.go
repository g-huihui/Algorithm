/**
 * @Author: Gong Yanhui
 * @Description: 132. 分割回文串 II
 * @Date: 2025-03-02 14:12
 */

package main

import (
	"math"
)

// 超时
func minCut(s string) int {
	var res = math.MaxInt
	var path []string

	dfs132(s, 0, path, &res)

	return res
}

func dfs132(s string, start int, path []string, res *int) {
	if start == len(s) && len(path)-1 < *res {
		*res = len(path) - 1
		return
	}

	// 从start开始 尝试每一种可能 start为每次分割的点
	for i := start; i < len(s); i++ {
		if !isPalindrome(s[start : i+1]) {
			continue
		}
		path = append(path, s[start:i+1])
		dfs132(s, i+1, path, res)
		path = path[:len(path)-1]
	}
}

func main() {
	println(minCut("aab")) // 1 ["aa","b"]
	println(minCut("a"))   // 0 ["a"]
	println(minCut("ab"))  // 1 ["a","b"]

	// 超时
	println(minCut("ababababababababababababcbabababababababababababa"))
}
