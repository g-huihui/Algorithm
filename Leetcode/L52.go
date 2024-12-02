/**
 * @Author: Gong Yanhui
 * @Description: 52. N 皇后 II
 * @Date: 2024-12-02 09:19
 */

package main

func totalNQueens(n int) int {
	var count int

	// 用于判断当前列是否有皇后
	columns := map[int]bool{}
	// 用于判断左斜线(i+j)和右斜线(i-j)是否有皇后
	diagonals1, diagonals2 := map[int]bool{}, map[int]bool{}
	queensDfsCount(n, 0, columns, diagonals1, diagonals2, &count)

	return count
}

func queensDfsCount(n int, i int, columns map[int]bool, diagonals1 map[int]bool, diagonals2 map[int]bool, count *int) {
	if i == n {
		*count++
		return
	}

	for j := 0; j < n; j++ {
		if columns[j] {
			continue
		}
		d1 := i + j
		if diagonals1[d1] {
			continue
		}
		d2 := i - j
		if diagonals2[d2] {
			continue
		}

		columns[j] = true
		diagonals1[d1] = true
		diagonals2[d2] = true

		queensDfsCount(n, i+1, columns, diagonals1, diagonals2, count)

		columns[j] = false
		diagonals1[d1] = false
		diagonals2[d2] = false
	}
}

func main() {
	println(totalNQueens(1))
	println(totalNQueens(4))
	println(totalNQueens(8))
}
