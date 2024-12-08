/**
 * @Author: Gong Yanhui
 * @Description: 3001. 捕获黑皇后需要的最少移动次数
 * @Date: 2024-12-05 15:43
 */

package main

func minMovesToCaptureTheQueen(a int, b int, c int, d int, e int, f int) int {
	// 理论上最多只有2步移动
	// (a,b)车 (c,d)象 (e,f)皇后

	// 先尝试用车去吃皇后 是否能1步吃掉
	var one bool
	// 车和皇后在 同一行 并且中间无障碍
	// 中间有障碍 多1步将障碍移走
	if a == e && !(c == e && inBetween(b, d, f)) {
		one = true
	}
	// 车和皇后在 同一列 并且中间无障碍
	if b == f && !(d == f && inBetween(a, c, e)) {
		one = true
	}

	// 象和皇后在同一对角线上
	var x, y, z = a + b, c + d, e + f
	if y == z && !(x == z && inBetween(c, a, e)) {
		one = true
	}
	x, y, z = a-b, c-d, e-f
	if y == z && !(x == z && inBetween(c, a, e)) {
		one = true
	}

	if one {
		return 1
	}
	return 2
}

// 判断 b 在不在 a和c之间
func inBetween(a, b, c int) bool {
	return (a < b && b < c) || (c < b && b < a)
}

func main() {
	println(minMovesToCaptureTheQueen(1, 1, 8, 8, 2, 3)) // 2
	println(minMovesToCaptureTheQueen(5, 3, 3, 4, 5, 2)) // 1

	println(minMovesToCaptureTheQueen(4, 3, 3, 4, 5, 2)) // 2
	println(minMovesToCaptureTheQueen(1, 1, 1, 4, 1, 8)) // 2
	println(minMovesToCaptureTheQueen(7, 8, 7, 7, 7, 3)) // 2
	/**
	r  0  0  b  0  0  0  q
	0  0  0  0  0  0  0  0
	0  0  0  0  0  0  0  0
	0  0  0  0  0  0  0  0
	0  0  0  0  0  0  0  0
	0  0  0  0  0  0  0  0
	0  0  0  0  0  0  0  0
	0  0  0  0  0  0  0  0
	*/
}
