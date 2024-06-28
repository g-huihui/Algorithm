/**
 * @Author: Gong Yanhui
 * @Description: 1103. 分糖果 II
 * @Date: 2024-06-28 14:52
 */

package main

func distributeCandies2(candies int, num_people int) []int {
	var res = make([]int, num_people)
	for count := 1; ; count++ {
		if candies > count {
			res[(count-1)%num_people] += count
			candies -= count
		} else {
			res[(count-1)%num_people] += candies
			break
		}
	}
	return res
}

func main() {
	res1 := distributeCandies2(7, 4)
	println(res1) // [1,2,3,1]

	res2 := distributeCandies2(10, 3)
	println(res2) // [5,2,3]
}
