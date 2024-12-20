/**
 * @Author: Gong Yanhui
 * @Description: 3285. 找到稳定山的下标
 * @Date: 2024-12-20 21:33
 */

package main

func stableMountains(height []int, threshold int) []int {
	var res []int

	for i := 1; i < len(height); i++ {
		if height[i-1] > threshold {
			res = append(res, i)
		}
	}

	return res
}
