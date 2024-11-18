/**
 * @Author: Gong Yanhui
 * @Description: 661. 图片平滑器
 * @Date: 2024-11-18 16:23
 */

package main

import "fmt"

func imageSmoother(img [][]int) [][]int {
	var row, col = len(img), len(img[0])

	var res = make([][]int, row)
	for i := 0; i < row; i++ {
		res[i] = make([]int, col)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			var sum, cnt = img[i][j], 1
			// 计算周围8个像素的值
			// 1. 左上
			if i-1 >= 0 && j-1 >= 0 {
				sum += img[i-1][j-1]
				cnt++
			}
			// 2. 上
			if i-1 >= 0 {
				sum += img[i-1][j]
				cnt++
			}
			// 3. 右上
			if i-1 >= 0 && j+1 < col {
				sum += img[i-1][j+1]
				cnt++
			}
			// 4. 左
			if j-1 >= 0 {
				sum += img[i][j-1]
				cnt++
			}
			// 5. 右
			if j+1 < col {
				sum += img[i][j+1]
				cnt++
			}
			// 6. 左下
			if i+1 < row && j-1 >= 0 {
				sum += img[i+1][j-1]
				cnt++
			}
			// 7. 下
			if i+1 < row {
				sum += img[i+1][j]
				cnt++
			}
			// 8. 右下
			if i+1 < row && j+1 < col {
				sum += img[i+1][j+1]
				cnt++
			}
			res[i][j] = sum / cnt
		}
	}

	return res
}

func main() {
	fmt.Println(imageSmoother([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}))
	// 输出: [[0, 0, 0],[0, 0, 0], [0, 0, 0]]

	fmt.Println(imageSmoother([][]int{{100, 200, 100}, {200, 50, 200}, {100, 200, 100}}))
	// 输出: [[137,141,137],[141,138,141],[137,141,137]]
}
