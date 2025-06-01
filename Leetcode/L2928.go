/**
 * @Author: Gong Yanhui
 * @Description: 2928. 给小朋友们分糖果 I
 * @Date: 2025-06-01 16:42
 */

package main

import "fmt"

func distributeCandies2928(n int, limit int) int {

	var res int
	for i := 0; i <= limit; i++ { // 第一位小朋友分的糖果数
		for j := 0; j <= limit; j++ { // 第二位小朋友分的糖果数
			var last = n - (i + j) // 剩余糖果数
			if last < 0 || last > limit {
				continue
			}
			res++
			//fmt.Printf("%d %d %d\n", i, j, last)
		}
	}

	return res
}

func main() {
	fmt.Println(distributeCandies2928(5, 2)) // 3
	fmt.Println(distributeCandies2928(3, 3)) // 10
}
