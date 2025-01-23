/**
 * @Author: Gong Yanhui
 * @Description: 2218. 从栈中取出 K 个硬币的最大面值和
 * @Date: 2025-01-23 11:20
 */

package main

func main() {
	// piles = [[1,100,3],[7,8,9]], k = 2
	// 输出：101
	println(maxValueOfCoins([][]int{{1, 100, 3}, {7, 8, 9}}, 2))

	// piles = [[100],[100],[100],[100],[100],[100],[1,1,1,1,1,1,700]], k = 7
	// 输出：706
	println(maxValueOfCoins([][]int{{100}, {100}, {100}, {100}, {100}, {100}, {1, 1, 1, 1, 1, 1, 700}}, 7))

	// piles = [[80,62,78,78,40,59,98,35],[79,19,100,15],[79,2,27,73,12,13,11,37,27,55,54,55,87,10,97,26,78,20,75,23,46,94,56,32,14,70,70,37,60,46,1,53]], k = 25
	// 超时
}

// 通过回溯算法完成 超时
func maxValueOfCoins(piles [][]int, k int) int {
	var res int
	f(piles, 0, k, &res)

	return res
}

func f(piles [][]int, sum, lastCount int, res *int) {
	if lastCount == 0 {
		if sum > *res {
			*res = sum
		}
		return
	}
	for i := 0; i < len(piles); i++ {
		for j := 0; j < len(piles[i]); j++ {
			if piles[i][j] == 0 { // 已经被取过了
				continue
			}
			if j > 0 && piles[i][j-1] != 0 { // 栈前面没有取过
				break
			}
			tmp := piles[i][j]
			piles[i][j] = 0 // 取走的标记为0
			f(piles, sum+tmp, lastCount-1, res)
			piles[i][j] = tmp
		}
	}
}
