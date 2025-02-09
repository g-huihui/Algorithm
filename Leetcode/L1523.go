/**
 * @Author: Gong Yanhui
 * @Description: 1523. 在区间范围内统计奇数数目
 * @Date: 2025-02-09 15:37
 */

package main

func countOdds(low int, high int) int {
	return (high+1)/2 - low/2
}
