/**
 * @Author: Gong Yanhui
 * @Description: 1812. 判断国际象棋棋盘中一个格子的颜色
 * @Date: 2024-12-09 18:17
 */

package main

import "strconv"

func squareIsWhite(coordinates string) bool {
	var x = int(coordinates[0] - 'a' + 1)
	y, _ := strconv.Atoi(string(coordinates[1]))
	return (x+y)%2 == 1
}
