/**
 * @Author: Gong Yanhui
 * @Description: 119. 杨辉三角 II
 * @Date: 2022-10-07 11:27
 */

package main

import "fmt"

func getRow(rowIndex int) []int {
	var res []int
	if rowIndex < 0 {
		return res
	}
	res = append(res, 1)
	for i := 0; i < rowIndex; i++ {
		res = nextRow(res)
	}
	return res
}

func nextRow(cur []int) []int {
	var res []int
	res = append(res, 1)
	length := len(cur)
	for i := 0; i < length; i++ {
		if len(cur) < 2 {
			res = append(res, cur[0])
			break
		}
		num := cur[0]
		cur = cur[1:]
		res = append(res, num+cur[0])
	}
	return res
}

func main() {
	fmt.Println(nextRow([]int{1, 1}))
	fmt.Println(getRow(0))
	fmt.Println(getRow(1))
}
