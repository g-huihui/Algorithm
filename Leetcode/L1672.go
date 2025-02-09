/**
 * @Author: Gong Yanhui
 * @Description: 1672. 最富有客户的资产总量
 * @Date: 2025-02-09 15:16
 */

package main

func maximumWealth(accounts [][]int) int {
	var res int
	for _, account := range accounts {
		var total int
		for _, money := range account {
			total += money
		}
		if total > res {
			res = total
		}
	}

	return res
}
