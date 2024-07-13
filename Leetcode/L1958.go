/**
 * @Author: Gong Yanhui
 * @Description: 1958. 检查操作是否合法
 * @Date: 2024-07-12 21:00
 */

package main

// 不想写了 累! 心累！
func checkMove(board [][]byte, rMove int, cMove int, color byte) bool {

	var row, col = len(board), len(board[0])

	// 左边行
	if cMove >= 2 {
		for i := cMove - 2; i >= 0; i-- {
			if board[rMove][i] == color && board[rMove][i+1] != color &&
				thisRowRangeIsSame(board, rMove, i+1, cMove-1) {
				return true
			}
		}
	}
	// 右边行
	if (cMove + 2) < col {
		for i := cMove + 2; i < col; i++ {
			if board[rMove][i] == color && board[rMove][i-1] != color &&
				thisRowRangeIsSame(board, rMove, cMove+1, i-1) {
				return true
			}
		}
	}

	// 上边列
	if rMove >= 2 {
		for i := rMove - 2; i >= 0; i-- {
			if board[i][cMove] == color && board[i+1][cMove] != color &&
				thisColRangeIsSame(board, cMove, i+1, rMove-1) {
				return true
			}
		}
	}
	// 下边列
	if (rMove + 2) < row {
		for i := rMove + 2; i < row; i++ {
			if board[i][cMove] == color && board[i-1][cMove] != color &&
				thisColRangeIsSame(board, cMove, rMove+1, i-1) {
				return true
			}
		}
	}

	// 左上斜线
	if rMove >= 2 && cMove >= 2 {
		minTmp := min(rMove, cMove)
		for rang := 0; (rang + 2) <= minTmp; rang++ {
			if board[rMove-(rang+2)][cMove-(rang+2)] == color && board[rMove-(rang+2)+1][cMove-(rang+2)+1] != color &&
				thisLeftSlashRangeIsSame(board, rMove, cMove, rang, minTmp) {
				return true
			}
		}
	}

	return false
}

// 当前行 范围内是否都相同 闭区间
func thisRowRangeIsSame(board [][]byte, row int, begin, end int) bool {
	// 如果中间只有一个节点不会进入循环
	for index := begin + 1; index <= end; index++ {
		if board[row][index-1] != board[row][index] {
			return false
		}
	}
	return true
}

// 当前列 范围内是否都相同 闭区间
func thisColRangeIsSame(board [][]byte, col int, begin, end int) bool {
	for index := begin + 1; index <= end; index++ {
		if board[index-1][col] != board[index][col] {
			return false
		}
	}
	return true
}

// 左斜线\ 范围内是否都相同 闭区间
func thisLeftSlashRangeIsSame(board [][]byte, row, col int, rang, minTmp int) bool {
	for index := 0; index <= rang-1; index++ {
		if board[row-(minTmp+index)+1][col-(minTmp+index)+1] != board[row+index][col-index] {
			return false
		}
	}
	return true
}

func main() {
	board1 := [][]byte{
		{'.', '.', '.', 'B', '.', '.', '.', '.'},
		{'.', '.', '.', 'W', '.', '.', '.', '.'},
		{'.', '.', '.', 'W', '.', '.', '.', '.'},
		{'.', '.', '.', 'W', '.', '.', '.', '.'},
		{'W', 'B', 'B', '.', 'W', 'W', 'W', 'B'},
		{'.', '.', '.', 'B', '.', '.', '.', '.'},
		{'.', '.', '.', 'B', '.', '.', '.', '.'},
		{'.', '.', '.', 'W', '.', '.', '.', '.'}}
	println(checkMove(board1, 4, 3, 'B')) // true

	board2 := [][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', 'B', '.', '.', 'W', '.', '.', '.'},
		{'.', '.', 'W', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'W', 'B', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', 'B', 'W', '.', '.'},
		{'.', '.', '.', '.', '.', '.', 'W', '.'},
		{'.', '.', '.', '.', '.', '.', '.', 'B'}}
	println(checkMove(board2, 4, 4, 'W')) // false
}
