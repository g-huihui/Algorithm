/**
 * @Author: Gong Yanhui
 * @Description: 3274. 检查棋盘方格颜色是否相同
 * @Date: 2024-12-03 13:09
 */

package main

func checkTwoChessboards(coordinate1 string, coordinate2 string) bool {
	return ((coordinate1[0]-coordinate2[0])+(coordinate1[1]-coordinate2[1]))%2 == 0
}

func main() {
	println(checkTwoChessboards("a1", "c3")) // true
	println(checkTwoChessboards("a1", "h3")) // false

	println(checkTwoChessboards("c2", "g4")) // true
}
