/**
 * @Author: Gong Yanhui
 * @Description: 3258. 统计满足 K 约束的子字符串数量 I
 * @Date: 2024-11-12 20:27
 */

package main

func countKConstraintSubstrings(s string, k int) int {
	var res int
	for begin := 0; begin < len(s); begin++ {
		var flag = make(map[byte]int, 2)
		for end := begin; end < len(s); end++ {
			flag[s[end]]++
			if flag['1'] <= k || flag['0'] <= k {
				res++
			} else {
				break
			}
		}
	}

	return res
}

func main() {
	println(countKConstraintSubstrings("10101", 1))   // 12
	println(countKConstraintSubstrings("1010101", 2)) // 25
	println(countKConstraintSubstrings("11111", 1))   // 15
}
