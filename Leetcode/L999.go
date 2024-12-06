/**
 * @Author: Gong Yanhui
 * @Description: 999. 可以被一步捕获的棋子数
 * @Date: 2024-12-06 10:47
 */

package main

func numRookCaptures(board [][]byte) int {
	// 棋盘大小
	const boardLen = 8

	// 记录卒的位置 p
	var pArr = make([][]int, 0)

	// 记录象的位置 B 横向位置和枞向位置
	var bRow = make(map[int][]int)
	var bCol = make(map[int][]int)

	// 记录车的位置 R
	var x, y int

	// 先遍历棋盘 找到所有棋子的位置
	for i := 0; i < boardLen; i++ {
		for j := 0; j < boardLen; j++ {
			switch board[i][j] {
			case 'p':
				pArr = append(pArr, []int{i, j})
			case 'R':
				x, y = i, j
			case 'B':
				bRow[i] = append(bRow[i], j)
				bCol[j] = append(bCol[j], i)
			}
		}
	}

	var count int
	// left right up down
	var left, right, up, down bool
	for _, ints := range pArr {
		// 当前行
		if ints[0] == x {
			between := isBetween(ints[0], y, ints[1], bRow)
			if ints[1] < y && !left && !between {
				// 左边
				count++
				left = true

			} else if ints[1] > y && !right && !between {
				// 右边
				count++
				right = true
			}
		}

		// 当前列
		if ints[1] == y {
			between := isBetween(ints[1], x, ints[0], bCol)
			if ints[0] < x && !up && !between {
				// 上边
				count++
				up = true
			} else if ints[0] > x && !down && !between {
				// 下边
				count++
				down = true
			}
		}
	}

	return count
}

func isBetween(index, x, y int, m map[int][]int) bool {
	tmp, ok := m[index]
	if !ok {
		return false
	}
	for _, v := range tmp {
		if (x < v && v < y) || (y < v && v < x) {
			return true
		}
	}

	return false
}

func main() {
	println(numRookCaptures([][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', 'R', '.', '.', '.', 'p'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'}})) // 3

	println(numRookCaptures([][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', 'p', 'p', 'p', 'p', 'p', '.', '.'},
		{'.', 'p', 'p', 'B', 'p', 'p', '.', '.'},
		{'.', 'p', 'B', 'R', 'B', 'p', '.', '.'},
		{'.', 'p', 'p', 'B', 'p', 'p', '.', '.'},
		{'.', 'p', 'p', 'p', 'p', 'p', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'}})) // 0
}
