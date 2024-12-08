/**
 * @Author: Gong Yanhui
 * @Description: 3248. 矩阵中的蛇
 * @Date: 2024-11-21 20:46
 */

package main

func finalPositionOfSnake(n int, commands []string) int {
	var count int
	var grid = make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, n)
	}

	// 初始化数据
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			grid[i][j] = count
			count++
		}
	}

	var row, col int
	for _, cmd := range commands {
		switch cmd {
		case "UP":
			row--
		case "DOWN":
			row++
		case "LEFT":
			col--
		case "RIGHT":
			col++
		default:
			panic("no command")
		}
	}

	return grid[row][col]
}
