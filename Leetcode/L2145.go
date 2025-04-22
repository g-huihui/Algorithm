/**
 * @Author: Gong Yanhui
 * @Description: 2145. 统计隐藏数组数目
 * @Date: 2025-04-21 20:38
 */

package main

// 优化后 复杂度O(len(differences))
func numberOfArrays(differences []int, lower int, upper int) int {
	var size = len(differences)
	var left, right = lower, upper // 将整个区间设置为左右边界 遍历一次数组逐渐缩小区间
	for i := 0; i < size; i++ {
		left = left + differences[i] // 计算下个最小的节点
		if left < lower {            // 最小值小于下边界 则将最小值设置为下边界
			left = lower
		} else if left > upper { // 最小值大于上边界 则区间内所有值都不可能符合条件 直接返回空
			return 0
		}
		right = right + differences[i] // 计算下一个最大的节点
		if right > upper {             // 最大值大于上边界 则将最大值设置为上边界
			right = upper
		} else if right < lower { // 最大值小于下边界 则区间内所有值都不可能符合条件 直接返回空
			return 0
		}
	}

	return right - left + 1
}

// 超时 复杂度O( len(differences)*(upper-lower+1) )
func numberOfArrays1(differences []int, lower int, upper int) int {
	var size = len(differences)
	var count int
	// 枚举这个区间内每个数作为数组的开始元素 遍历出符合范围内的节点
	for i := lower; i <= upper; i++ {
		// 将i设置为数组的第一个元素
		var index, num = 1, i
		for index < size+1 {
			num = num + differences[index-1] // 计算下一个节点
			if num < lower || num > upper {
				// 如果当前节点超出范围 则改数不符合条件 直接退出 枚举下一个数
				break
			}
			index++
		}
		if index == size+1 { // 判断是否遍历完所有节点
			count++
		}
	}

	return count
}
