/**
 * @Author: Gong Yanhui
 * @Description: 624. 数组列表中的最大距离
 * @Date: 2025-02-19 23:25
 */

package main

func maxDistance(arrays [][]int) int {
	var res = 0
	var n = len(arrays)
	var minValue, maxValue = arrays[0][0], arrays[0][len(arrays[0])-1]

	for i := 1; i < n; i++ {
		var size = len(arrays[i])
		res = max(res, max(arrays[i][size-1]-minValue, maxValue-arrays[i][0]))
		minValue = min(minValue, arrays[i][0])
		maxValue = max(maxValue, arrays[i][size-1])
	}

	return res
}

func main() {
	// 输入：[[1,2,3],[4,5],[1,2,3]]
	// 输出：4
	println(maxDistance([][]int{{1, 2, 3}, {4, 5}, {1, 2, 3}}))

	// 输入：arrays = [[1],[1]]
	// 输出：0
	println(maxDistance([][]int{{1}, {1}}))

	// [[1,5],[3,4]] => 3
	println(maxDistance([][]int{{1, 5}, {3, 4}}))
}
