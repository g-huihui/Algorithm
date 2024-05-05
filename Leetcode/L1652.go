/**
 * @Author: Gong Yanhui
 * @Description: 1652. 拆炸弹
 * @Date: 2024-05-05 22:40
 */

package main

func decrypt(code []int, k int) []int {
	var n = len(code)
	var res = make([]int, n)
	if k == 0 {
		return res
	}
	for i := 0; i < n; i++ {
		res[i] = 0
		if k > 0 {
			for j := 1; j <= k; j++ {
				res[i] += code[(i+j)%n]
			}
		} else {
			tmp := -k
			for j := 1; j <= tmp; j++ {
				res[i] += code[(i-j+n)%n]
			}
		}
	}
	return res
}
