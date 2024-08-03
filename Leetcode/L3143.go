/**
 * @Author: Gong Yanhui
 * @Description: 3143. 正方形中的最多点数
 * @Date: 2024-08-03 17:02
 */

package main

import "math"

func maxPointsInsideSquare(points [][]int, s string) int {
	var minMap = make(map[rune]int)
	var minFlag = math.MaxInt

	for index, c := range s {
		old, ok := minMap[c]
		if ok {
			newNum := getMaxAbs(points[index][0], points[index][1])
			if newNum < old {
				cleanOver(minMap, old)
				minFlag = old
				minMap[c] = getMaxAbs(points[index][0], points[index][1])
			} else if newNum > old {
				if newNum < minFlag {
					cleanOver(minMap, newNum)
					minFlag = newNum
				}
			} else {
				cleanOver(minMap, newNum)
				minFlag = newNum
			}
		} else {
			tmp := getMaxAbs(points[index][0], points[index][1])
			if tmp < minFlag {
				minMap[c] = tmp
			}
		}
	}

	return len(minMap)
}

func cleanOver(minMap map[rune]int, tmp int) {
	var over []rune
	for r, num := range minMap {
		if num >= tmp {
			over = append(over, r)
		}
	}
	for _, o := range over {
		delete(minMap, o)
	}
}

func getMaxAbs(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	if a > b {
		return a
	}
	return b
}

func main() {
	println(maxPointsInsideSquare([][]int{{2, 2}, {-1, -2}, {-4, 4}, {-3, 1}, {3, -3}}, "abdca")) // 2
	println(maxPointsInsideSquare([][]int{{1, 1}, {-2, -2}, {-2, 2}}, "abb"))                     // 1
	println(maxPointsInsideSquare([][]int{{1, 1}, {-1, -1}, {2, -2}}, "ccd"))                     // 0

	println(maxPointsInsideSquare([][]int{{16, 32}, {27, 3}, {23, -14}, {-32, -16}, {-3, 26}, {-14, 33}}, "aaabfc")) // 2
}
