/**
 * @Author: Gong Yanhui
 * @Description: 657. 机器人能否返回原点
 * @Date: 2025-01-26 17:02
 */

package main

func judgeCircle(moves string) bool {
	var u, d, l, r = 0, 0, 0, 0
	for i := 0; i < len(moves); i++ {
		switch moves[i] {
		case 'U':
			u++
		case 'D':
			d++
		case 'L':
			l++
		case 'R':
			r++
		}
	}

	return u == d && l == r
}
