/**
 * @Author: Gong Yanhui
 * @Description: 3153. 所有数对中数位差之和
 * @Date: 2024-08-30 20:06
 */

package main

import "strconv"

// 懵的
func sumDigitDifferences(nums []int) int64 {
	var res int64
	cnt := make([][10]int, len(strconv.Itoa(nums[0])))
	for index, num := range nums {
		for i := 0; num > 0; num /= 10 {
			d := num % 10
			res += int64(index - cnt[i][d])
			cnt[i][d]++
			i++
		}
	}
	return res
}

// 超时 未通过
func sumDigitDifferences1(nums []int) int64 {

	var n = make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		var tmpArr []int
		var tmpNum = nums[i]
		for tmpNum > 0 {
			tmpArr = append(tmpArr, tmpNum%10)
			tmpNum = tmpNum / 10
		}
		n[i] = tmpArr
	}

	var res int64
	for i := 0; i < len(n)-1; i++ {
		for j := i + 1; j < len(n); j++ {
			for index := 0; index < len(n[0]); index++ {
				if n[i][index] != n[j][index] {
					res++
				}
			}
		}
	}

	return res
}

func main() {
	println(sumDigitDifferences([]int{13, 23, 12}))     // 4
	println(sumDigitDifferences([]int{10, 10, 10, 10})) // 0
	println(sumDigitDifferences([]int{1, 1, 3, 4, 5}))  // 9
}
