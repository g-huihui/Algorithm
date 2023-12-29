/**
 * @Author: Gong Yanhui
 * @Description: 2706. 购买两块巧克力
 * @Date: 2023-12-29 17:14
 */

package main

import "sort"

func buyChoco(prices []int, money int) int {
	if len(prices) < 2 {
		return money
	}
	sort.Ints(prices)
	if prices[0]+prices[1] > money {
		return money
	}
	return money - prices[0] - prices[1]
}
