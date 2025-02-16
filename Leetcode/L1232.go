/**
 * @Author: Gong Yanhui
 * @Description: 1232. 缀点成线
 * @Date: 2025-02-16 17:24
 */

package main

func checkStraightLine(coordinates [][]int) bool {
	var size = len(coordinates)
	for i := 0; i < size-2; i++ {
		if (coordinates[i][0]-coordinates[i+1][0])*(coordinates[i][1]-coordinates[i+2][1]) !=
			(coordinates[i][0]-coordinates[i+2][0])*(coordinates[i][1]-coordinates[i+1][1]) {
			return false
		}
	}
	return true
}

// (2,1) (4,2) (6,3)
// 通过斜率相等比较 (2-4)/(1-2) == (2-6)/(1-3)
// ==> 数学换算 (2-4)(1-3) == (2-6)(1-2)
