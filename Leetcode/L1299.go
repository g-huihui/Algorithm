/**
 * @Author: Gong Yanhui
 * @Description: 1299. 将每个元素替换为右侧最大元素
 * @Date: 2025-02-16 16:48
 */

package main

import (
	"fmt"
	"math"
)

func replaceElements(arr []int) []int {
	var maxIndex = -1
	var size = len(arr)
	var res = make([]int, 0, size)
	for i := 0; i < size; i++ {
		if maxIndex <= i {
			maxIndex = get1299(i+1, arr)
		}
		if maxIndex == -1 {
			res = append(res, -1)
		} else {
			res = append(res, arr[maxIndex])
		}
	}

	return res
}

// 根据给定的数组 从开始索引开始遍历找到最大值的索引位置
func get1299(begin int, arr []int) int {
	var maxNum, index = math.MinInt, -1
	for i := begin; i < len(arr); i++ {
		if arr[i] > maxNum {
			maxNum = arr[i]
			index = i
		}
	}

	return index
}

func main() {
	// 输入：arr = [17,18,5,4,6,1]
	// 输出：[18,6,6,6,1,-1]
	fmt.Println(replaceElements([]int{17, 18, 5, 4, 6, 1}))

	// 输入：arr = [400]
	// 输出：[-1]
	fmt.Println(replaceElements([]int{400}))
}
