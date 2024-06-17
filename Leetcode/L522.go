/**
 * @Author: Gong Yanhui
 * @Description: 522. 最长特殊序列 II
 * @Date: 2024-06-17 14:27
 */

package main

func findLUSlength(strs []string) int {
	var res = -1
	for i := 0; i < len(strs); i++ {
		var flag = false
		for j := 0; j < len(strs); j++ {
			if i != j && isSubseq(strs[i], strs[j]) {
				flag = true
				break
			}
		}
		if !flag && len(strs[i]) > res {
			res = len(strs[i])
		}
	}
	return res
}

func main() {
	println(findLUSlength([]string{"aba", "cdc", "eae"})) // 3
	println(findLUSlength([]string{"aaa", "aaa", "aa"}))  // -1

	println(findLUSlength([]string{"aabbcc", "aabbcc", "cb"})) // 2
}

func isSubseq(s, t string) bool {
	ptS := 0
	for ptT := range t {
		if s[ptS] == t[ptT] {
			if ptS++; ptS == len(s) {
				return true
			}
		}
	}
	return false
}
