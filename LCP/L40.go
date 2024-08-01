/**
 * @Author: Gong Yanhui
 * @Description: LCP 40. 心算挑战
 * @Date: 2024-08-01 11:42
 */

package main

import (
	"fmt"
	"sort"
)

func maxmiumScore(cards []int, cnt int) int {
	// 倒序排序
	sort.Slice(cards, func(i, j int) bool {
		return cards[i] > cards[j]
	})

	var sum = 0
	for i := 0; i < cnt; i++ {
		sum += cards[i]
	}
	if checkNumber(sum) == 0 { // 偶数
		return sum
	}
	var length = len(cards)
	var maxRes = 0
	for i := cnt - 1; i >= 0; i-- {
		var flag = checkNumber(cards[i])
		sumTmp := sum - cards[i]
		for j := cnt; j < length; j++ {
			if checkNumber(cards[j]) != flag {
				maxRes = max(sumTmp+cards[j], maxRes)
				break
			}
		}
	}

	return maxRes
}

// 奇数返回1 偶数返回0
func checkNumber(num int) int {
	if num%2 == 0 {
		return 0
	}
	return 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxmiumScore([]int{1, 2, 8, 9}, 3)) // 18
	fmt.Println(maxmiumScore([]int{3, 3, 1}, 1))    // 0

	fmt.Println(maxmiumScore([]int{9, 5, 9, 1, 6, 10, 3, 4, 5, 1}, 2)) // 18
}
