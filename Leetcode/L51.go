/**
 * @Author: Gong Yanhui
 * @Description: 51. N 皇后
 * @Date: 2024-12-01 21:36
 */

package main

import "fmt"

// 暴力破解
func solveNQueens(n int) [][]string {
	var res [][]string

	// 用于生成皇后的结果
	queens := make([]int, 0)
	// 用于判断当前列是否有皇后
	columns := map[int]bool{}
	// 用于判断左斜线(i+j)和又斜线(i-j)是否有皇后
	diagonals1, diagonals2 := map[int]bool{}, map[int]bool{}
	queensDfs(queens, n, 0, columns, diagonals1, diagonals2, &res)

	return res
}

// queensDfs
//
//	@Description: 回溯算法
//	@param queens 皇后的位置 用于生成结果
//	@param size	皇后的数量
//	@param row 当前行
//	@param columns 当前列是否有皇后
//	@param diagonals1 左斜线是否有皇后
//	@param diagonals2 右斜线是否有皇后
//	@param res
func queensDfs(queens []int, size, row int, columns, diagonals1, diagonals2 map[int]bool, res *[][]string) {
	if row == size {
		*res = append(*res, getQueensResArr(queens))
		return
	}

	for i := 0; i < size; i++ {
		if columns[i] { // 当前列有皇后
			continue
		}
		d1 := row + i
		if diagonals1[d1] { // 左斜线有皇后
			continue
		}
		d2 := row - i
		if diagonals2[d2] { // 右斜线有皇后
			continue
		}

		queens = append(queens, i)
		columns[i] = true
		diagonals1[d1] = true
		diagonals2[d2] = true

		queensDfs(queens, size, row+1, columns, diagonals1, diagonals2, res)

		queens = queens[:len(queens)-1]
		delete(columns, i)
		delete(diagonals1, d1)
		delete(diagonals2, d2)
	}
}

func getQueensResArr(arr []int) []string {
	var size = len(arr)
	res := make([]string, size)
	for i := 0; i < size; i++ {
		str := ""
		for j := 0; j < size; j++ {
			if arr[i] == j {
				str += "Q"
			} else {
				str += "."
			}
		}
		res[i] = str
	}
	return res
}

func main() {
	// 测试生成皇后的结果函数
	fmt.Println(getQueensResArr([]int{1, 0}))

	// 测试N皇后
	fmt.Println(solveNQueens(1))
	fmt.Println(solveNQueens(4))
}
