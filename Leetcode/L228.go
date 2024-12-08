/**
 * @Author: Gong Yanhui
 * @Description: 228. 汇总区间
 * @Date: 2024-12-06 14:24
 */

package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	var res []string

	for index, begin := 0, 0; index < len(nums); index++ {
		if index+1 < len(nums) && nums[index+1] == nums[index]+1 {
			continue
		}

		if begin == index {
			res = append(res, strconv.Itoa(nums[begin]))
		} else {
			res = append(res, strconv.Itoa(nums[begin])+"->"+strconv.Itoa(nums[index]))
		}

		begin = index + 1
	}

	return res
}

func main() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
}
