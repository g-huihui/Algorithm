/**
 * @Author: Gong Yanhui
 * @Description: 2684. 矩阵中移动的最大次数
 * @Date: 2024-03-16 17:41
 */

package main

import "fmt"

func maxMoves(grid [][]int) int {
	var col = len(grid)
	var row = len(grid[0])
	var copyGrid = make([][]int, col)
	for i := 0; i < col; i++ {
		copyGrid[i] = make([]int, row)
		copyGrid[i][0] = 1
	}

	for i := 0; i < row-1; i++ {
		for j := 0; j < col; j++ {
			if copyGrid[j][i] != 0 && j != 0 && grid[j-1][i+1] > grid[j][i] {
				copyGrid[j-1][i+1] = i + 1
			}
			if copyGrid[j][i] != 0 && grid[j][i+1] > grid[j][i] {
				copyGrid[j][i+1] = i + 1
			}
			if copyGrid[j][i] != 0 && j != col-1 && grid[j+1][i+1] > grid[j][i] {
				copyGrid[j+1][i+1] = i + 1
			}
		}
		if isOver(copyGrid, i+1) {
			return i
		}
	}
	return row - 1
}

func isOver(grid [][]int, index int) bool {
	var res = true
	var col = len(grid)
	for i := 0; i < col; i++ {
		if grid[i][index] != 0 {
			res = false
			break
		}
	}
	return res
}

func main() {
	var grid = [][]int{
		{2, 4, 3, 5},
		{5, 4, 9, 3},
		{3, 4, 2, 11},
		{10, 9, 13, 15},
	}
	if maxMoves(grid) != 3 {
		fmt.Println(maxMoves(grid)) // 3
	}

	var grid1 = [][]int{
		{3, 2, 4},
		{2, 1, 9},
		{1, 1, 7},
	}
	if maxMoves(grid1) != 0 {
		fmt.Println(maxMoves(grid1)) // 0
	}

	var grid2 = [][]int{
		{1000000, 92910, 1021, 1022, 1023, 1024, 1025, 1026, 1027, 1028, 1029, 1030, 1031, 1032, 1033, 1034, 1035, 1036, 1037, 1038, 1039, 1040, 1041, 1042, 1043, 1044, 1045, 1046, 1047, 1048, 1049, 1050, 1051, 1052, 1053, 1054, 1055, 1056, 1057, 1058, 1059, 1060, 1061, 1062, 1063, 1064, 1065, 1066, 1067, 1068},
		{1069, 1070, 1071, 1072, 1073, 1074, 1075, 1076, 1077, 1078, 1079, 1080, 1081, 1082, 1083, 1084, 1085, 1086, 1087, 1088, 1089, 1090, 1091, 1092, 1093, 1094, 1095, 1096, 1097, 1098, 1099, 1100, 1101, 1102, 1103, 1104, 1105, 1106, 1107, 1108, 1109, 1110, 1111, 1112, 1113, 1114, 1115, 1116, 1117, 1118},
	}
	if maxMoves(grid2) != 49 {
		fmt.Println(maxMoves(grid2)) // 49
	}
}
