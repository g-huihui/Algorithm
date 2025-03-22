/**
 * @Author: Gong Yanhui
 * @Description: 2643. 一最多的行
 * @Date: 2025-03-22 10:27
 */

package main

import "fmt"

func rowAndMaximumOnes(mat [][]int) []int {
	var res = make([]int, 2)
	for i, ints := range mat {
		var count int
		for _, num := range ints {
			if num == 1 {
				count++
			}
		}
		if count > res[1] {
			res[0] = i
			res[1] = count
		}
	}

	return res
}

func main() {
	fmt.Println(rowAndMaximumOnes([][]int{{0, 1}, {1, 0}}))         // [0,1]
	fmt.Println(rowAndMaximumOnes([][]int{{0, 0, 0}, {0, 1, 1}}))   // [1,2]
	fmt.Println(rowAndMaximumOnes([][]int{{0, 0}, {1, 1}, {0, 0}})) // [1,2]
}
