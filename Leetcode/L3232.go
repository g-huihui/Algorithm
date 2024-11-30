/**
 * @Author: Gong Yanhui
 * @Description: 3232. 判断是否可以赢得数字游戏
 * @Date: 2024-11-30 14:48
 */

package main

func canAliceWin(nums []int) bool {
	var single, singleTotal int
	var double, doubleTotal int
	for i := 0; i < len(nums); i++ {
		if nums[i] < 10 {
			single += nums[i]
		} else {
			singleTotal += nums[i]
		}
		if nums[i] >= 10 && nums[i] < 100 {
			double += nums[i]
		} else {
			doubleTotal += nums[i]
		}
	}

	return single > singleTotal || double > doubleTotal
}
