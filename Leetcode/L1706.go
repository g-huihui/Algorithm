/**
 * @Author: Gong Yanhui
 * @Description: 1706. 球会落何处
 * @Date: 2025-02-15 15:50
 */

package main

func findBall(grid [][]int) []int {
	var high = len(grid)    // 网格高度
	var size = len(grid[0]) // 小球数 = 列数

	var res = make([]int, 0, size)    // 要返回的结果
	for num := 0; num < size; num++ { // 从第一个小球开始 循环每个小球
		var flag = true             //	标记小球是否到达最后
		var col = num               // 记录小球到达的列号
		for i := 0; i < high; i++ { // 循环每一层
			// 小球在最左边并且对角线向左
			if col == 0 && grid[i][0] == -1 {
				flag = false
				break
			}
			// 小球在最右边并且对角线向右
			if col == size-1 && grid[i][size-1] == 1 {
				flag = false
				break
			}
			// 小球向右 右边被挡 或者 小球向左 左边被挡
			if (grid[i][col] == 1 && grid[i][col] != grid[i][col+1]) ||
				(grid[i][col] == -1 && grid[i][col] != grid[i][col-1]) {
				flag = false
				break
			}

			// 走到这里说明小球可以向下继续走 判断是向左走还是向右走
			if grid[i][col] == 1 {
				col++
			} else { // grid[i][col] == -1
				col--
			}
		}
		if flag { // 小球走完了每一层
			res = append(res, col)
		} else {
			res = append(res, -1)
		}
	}

	return res
}
