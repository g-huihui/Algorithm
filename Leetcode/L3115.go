/**
 * @Author: Gong Yanhui
 * @Description: 3115. 质数的最大距离
 * @Date: 2024-07-04 10:33
 */

package main

func maximumPrimeDifference(nums []int) int {

	// 预处理所有不大于100的质数
	var arr = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	var m = make(map[int]struct{})
	for _, num := range arr {
		m[num] = struct{}{}
	}

	var begin, end = -1, -1
	for index, num := range nums {
		_, ok := m[num]
		if !ok {
			continue
		}
		if begin == -1 {
			begin = index
			end = index
			continue
		}
		end = index
	}

	return end - begin
}

func main() {
	println(maximumPrimeDifference([]int{4, 2, 9, 5, 3})) // 3
	println(maximumPrimeDifference([]int{4, 8, 2, 8}))    // 0
}
