/**
 * @Author: Gong Yanhui
 * @Description: 688. 骑士在棋盘上的概率
 * @Date: 2024-12-07 20:47
 */

package main

// 未通过
func knightProbability(n int, k int, row int, column int) float64 {

	// 代表马的8种走向
	var dir = []struct {
		x, y int
	}{
		{1, 2}, {2, 1},
		{-1, 2}, {-2, 1},
		{1, -2}, {2, -1},
		{-1, -2}, {-2, -1},
	}

	var dfs func(k, x, y int) float64
	dfs = func(k, x, y int) float64 {
		// 查看当前索引是否已经跳出棋盘
		if x < 0 || y < 0 || x >= n || y >= n {
			return 0
		}
		// 已经没有吓一跳 并且还在棋盘中
		if k == 0 {
			return 1
		}
		var res float64
		var count int
		for _, s := range dir {
			var tmp = dfs(k-1, x+s.x, y+s.y)
			if tmp != 0 {
				count++
				res += tmp
			}
		}
		res /= float64(count)
		return res
	}

	return dfs(k, row, column)
}

func main() {
	println(knightProbability(3, 2, 0, 0)) // 0.0625
	println(knightProbability(1, 0, 0, 0)) // 1.00000
}
