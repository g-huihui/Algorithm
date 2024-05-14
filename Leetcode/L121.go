/**
 * @Author: Gong Yanhui
 * @Description: 121. 买卖股票的最佳时机
 * @Date: 2024-05-14 14:48
 */

package main

func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}
	var minPriceIndex, maxDiff = 0, prices[1] - prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[minPriceIndex] > prices[i] {
			minPriceIndex = i
		}
		if prices[i]-prices[minPriceIndex] > maxDiff {
			maxDiff = prices[i] - prices[minPriceIndex]
		}
	}
	return maxDiff
}
