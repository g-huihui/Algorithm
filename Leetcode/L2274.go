/**
 * @Author: Gong Yanhui
 * @Description: 2274. 不含特殊楼层的最大连续楼层数
 * @Date: 2025-01-06 00:07
 */

package main

import "sort"

func maxConsecutive2274(bottom int, top int, special []int) int {
	var maxDistance int

	sort.Ints(special)

	// 先判断特殊楼层之间的最大距离
	for i := 0; i < len(special)-1; i++ {
		maxDistance = max(maxDistance, special[i+1]-special[i]-1)
	}

	// 判断开始楼层和第一个特殊楼层之间的距离
	if special[0] != bottom {
		maxDistance = max(maxDistance, special[0]-bottom)
	}

	// 判断最后一个特殊楼层和结束楼层之间的距离
	if special[len(special)-1] != top {
		maxDistance = max(maxDistance, top-special[len(special)-1])
	}

	return maxDistance
}
