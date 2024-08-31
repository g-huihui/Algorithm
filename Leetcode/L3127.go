/**
 * @Author: Gong Yanhui
 * @Description: 3127. 构造相同颜色的正方形
 * @Date: 2024-08-31 21:33
 */

package main

func canMakeSquare(grid [][]byte) bool {
	var length = len(grid)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1; j++ {
			white, black := colorCount(grid, i, j)
			if white >= 3 || black >= 3 {
				return true
			}
		}
	}
	return false
}

func colorCount(grid [][]byte, x, y int) (int, int) {
	var white, black int
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if grid[x+i][y+j] == 'W' {
				white++
			} else {
				black++
			}
		}
	}
	return white, black
}
