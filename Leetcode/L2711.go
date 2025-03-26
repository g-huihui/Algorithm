/**
 * @Author: Gong Yanhui
 * @Description: 2711. 对角线上不同值的数量差
 * @Date: 2025-03-25 21:55
 */

package main

import "fmt"

// 官方解法 此解法暴力破解
func differenceOfDistinctValues(grid [][]int) [][]int {
	m, n := len(grid), len(grid[0])
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			s1 := make(map[int]bool)
			x, y := i+1, j+1
			for x < m && y < n {
				s1[grid[x][y]] = true
				x++
				y++
			}
			s2 := make(map[int]bool)
			x, y = i-1, j-1
			for x >= 0 && y >= 0 {
				s2[grid[x][y]] = true
				x--
				y--
			}
			res[i][j] = abs(len(s1) - len(s2))
		}
	}
	return res
}

// 379 / 709 个通过的测试用例
func differenceOfDistinctValues2(grid [][]int) [][]int {
	var row = len(grid)
	var col = len(grid[0])

	var left = make([][]int, row)
	for i := 0; i < row; i++ {
		left[i] = make([]int, col)
	}
	for i := 1; i < col; i++ {
		var m, n = 1, i
		var cache = make(map[int]struct{})
		for m < row && n < col {
			if _, ok := cache[grid[m][n]]; ok {
				left[m][n] = left[m-1][n-1]
			} else {
				left[m][n] = left[m-1][n-1] + 1
				cache[grid[m][n]] = struct{}{}
			}
			m++
			n++
		}
	}
	for i := 2; i < row; i++ {
		var m, n = i, 1
		var cache = make(map[int]struct{})
		for m < row && n < col {
			if _, ok := cache[grid[m][n]]; ok {
				left[m][n] = left[m-1][n-1]
			} else {
				left[m][n] = left[m-1][n-1] + 1
				cache[grid[m][n]] = struct{}{}
			}
			m++
			n++
		}
	}

	var right = make([][]int, row)
	for i := 0; i < row; i++ {
		right[i] = make([]int, col)
	}
	for i := 0; i < col-1; i++ {
		var m, n = row - 2, i
		var cache = make(map[int]struct{})
		for m >= 0 && n >= 0 {
			if _, ok := cache[grid[m][n]]; ok {
				right[m][n] = right[m+1][n+1]
			} else {
				right[m][n] = right[m+1][n+1] + 1
				cache[grid[m][n]] = struct{}{}
			}
			m--
			n--
		}
	}
	for i := 0; i < row-1; i++ {
		var m, n = i, col - 2
		var cache = make(map[int]struct{})
		for m >= 0 && n >= 0 {
			if _, ok := cache[grid[m][n]]; ok {
				right[m][n] = right[m+1][n+1]
			} else {
				right[m][n] = right[m+1][n+1] + 1
				cache[grid[m][n]] = struct{}{}
			}
			m--
			n--
		}
	}

	// 封装返回结果
	var result = make([][]int, row)
	for i := 0; i < row; i++ {
		result[i] = make([]int, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			result[i][j] = abs(left[i][j] - right[i][j])
		}
	}

	return result
}

// 当前方法只考虑了 len(grid) == len(grid[0]) 的情况
func differenceOfDistinctValues1(grid [][]int) [][]int {
	var size = len(grid)

	// 左上记录
	var left = make([][]int, size)
	for i := 0; i < size; i++ {
		left[i] = make([]int, size)
	}
	// 生成数据
	//0: (0,2)
	//1: (0,1) (1,2)
	//2: (0,0) (1,1) (2,2)
	for i := size - 1; i >= 0; i-- {
		var m, n = 1, i + 1
		var cache = make(map[int]struct{})
		for m < size && n < size {
			if _, ok := cache[grid[m][n]]; ok {
				left[m][n] = left[m-1][n-1]
			} else {
				left[m][n] = left[m-1][n-1] + 1
				cache[grid[m][n]] = struct{}{}
			}
			m++
			n++
		}
	}
	//1: (1,0) (2,1)
	//2: (2,0)
	for i := 1; i < size; i++ {
		var m, n = i + 1, 1
		var cache = make(map[int]struct{})
		for m < size && n < size {
			if _, ok := cache[grid[m][n]]; ok {
				left[m][n] = left[m-1][n-1]
			} else {
				left[m][n] = left[m-1][n-1] + 1
				cache[grid[m][n]] = struct{}{}
			}
			m++
			n++
		}
	}

	// 右下记录
	var right = make([][]int, size)
	for i := 0; i < size; i++ {
		right[i] = make([]int, size)
	}
	// 生成数据
	//0: (0,2)
	//1: (1,2) (0,1)
	//2: (2,2) (1,1) (0,0)
	for i := 0; i < size; i++ {
		var m, n = i - 1, size - 2
		var cache = make(map[int]struct{})
		for m >= 0 && n >= 0 {
			if _, ok := cache[grid[m][n]]; ok {
				right[m][n] = right[m+1][n+1]
			} else {
				right[m][n] = right[m+1][n+1] + 1
				cache[grid[m][n]] = struct{}{}
			}
			m--
			n--
		}
	}
	//1: (2,1) (1,0)
	//2: (2,0)
	for i := size - 2; i >= 0; i-- {
		var m, n = size - 2, size - i - 3
		var cache = make(map[int]struct{})
		for m >= 0 && n >= 0 {
			if _, ok := cache[grid[m][n]]; ok {
				right[m][n] = right[m+1][n+1]
			} else {
				right[m][n] = right[m+1][n+1] + 1
				cache[grid[m][n]] = struct{}{}
			}
			m--
			n--
		}
	}

	// 封装返回结果
	var result = make([][]int, size)
	for i := 0; i < size; i++ {
		result[i] = make([]int, size)
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			result[i][j] = abs(left[i][j] - right[i][j])
		}
	}

	return result
}

func main() {
	fmt.Println(differenceOfDistinctValues([][]int{{1, 2, 3}, {3, 1, 5}, {3, 2, 1}}))
	fmt.Println(differenceOfDistinctValues([][]int{{5, 50, 39, 37}, {2, 19, 36, 26}, {27, 3, 23, 10}, {20, 33, 8, 39}}))
	// result [3,2,1,0] 	left: [0,0,0,0]		right: [3,2,1,0]
	//		  [2,1,0,1]			  [0,1,1,1]			   [2,2,1,0]
	//		  [1,0,1,2]			  [0,1,2,2]			   [1,1,1,0]
	//		  [0,1,2,3]			  [0,1,2,3]			   [0,0,0,0]

	fmt.Println(differenceOfDistinctValues([][]int{
		{6, 28, 37, 34, 12, 30, 43, 35, 6},
		{6, 28, 37, 34, 12, 30, 43, 35, 6},
		{6, 28, 37, 34, 12, 30, 43, 35, 6},
		{6, 28, 37, 34, 12, 30, 43, 35, 6}}))
	// result:
	// [3, 3, 3, 3, 3, 3, 2, 1, 0]
	// [2, 1, 1, 1, 1, 1, 1, 0, 1]
	// [1, 0, 1, 1, 1, 1, 1, 1, 2]
	// [0, 1, 2, 3, 3, 3, 3, 3, 3]
}

/*
原始结构:	-->		最终结构:
[1, 2, 3]			[1, 1, 0]
[3, 1, 5]			[1, 0, 1]
[3, 2, 1]			[0, 1, 1]

分析: 分别根据左上和右下记录对应的计数
Left:				Right:
[0, 0, 0]			[1, 1, 0]
[0, 1, 1]			[1, 1, 0]
[0, 1, 1]			[0, 0, 0]

循环逻辑:
Left:
0: (0,2)
1: (0,1) (1,2)
2: (0,0) (1,1) (2,2)
1: (1,0) (2,1)
2: (2,0)

Right:
0: (0,2)
1: (1,2) (0,1)
2: (2,2) (1,1) (0,0)
1: (2,1) (1,0)
2: (2,0)
*/
