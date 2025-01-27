/**
 * @Author: Gong Yanhui
 * @Description: 1275. 找出井字棋的获胜者
 * @Date: 2025-01-27 22:58
 */

package main

func tictactoe(moves [][]int) string {
	var chessboard = make([][]string, 3)
	for i := 0; i < 3; i++ {
		chessboard[i] = make([]string, 3)
	}

	// 循环下每一步棋 判断是否有获胜者
	for i, move := range moves {
		var player string
		if i%2 == 0 {
			player = "A"
		} else {
			player = "B"
		}
		chessboard[move[0]][move[1]] = player
		ok, res := check1275(chessboard)
		if ok {
			return res
		}
	}

	if len(moves) < 9 {
		return "Pending"
	}

	return "Draw"
}

// 判断当前棋盘是否有获胜者
func check1275(chessboard [][]string) (bool, string) {
	// 横向判断
	for i := 0; i < 3; i++ {
		if chessboard[i][0] == chessboard[i][1] && chessboard[i][1] == chessboard[i][2] && chessboard[i][0] != "" {
			return true, chessboard[i][0]
		}
	}

	// 纵向判断
	for i := 0; i < 3; i++ {
		if chessboard[0][i] == chessboard[1][i] && chessboard[1][i] == chessboard[2][i] && chessboard[0][i] != "" {
			return true, chessboard[0][i]
		}
	}

	// 斜向判断
	if chessboard[0][0] == chessboard[1][1] && chessboard[1][1] == chessboard[2][2] && chessboard[0][0] != "" {
		return true, chessboard[0][0]
	}
	if chessboard[0][2] == chessboard[1][1] && chessboard[1][1] == chessboard[2][0] && chessboard[0][2] != "" {
		return true, chessboard[0][2]
	}

	return false, ""
}
