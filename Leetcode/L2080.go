/**
 * @Author: Gong Yanhui
 * @Description: 2080. 区间内查询数字的频率
 * @Date: 2025-02-18 22:23
 */

package main

type RangeFreqQuery struct {
	// 记录每种键在数组中出现的索引下标
	cache map[int][]int
}

func Constructor(arr []int) RangeFreqQuery {
	var m = make(map[int][]int)
	for index, key := range arr {
		// 保证每个key的index列是有序的可以做二分查找
		m[key] = append(m[key], index)
	}

	return RangeFreqQuery{cache: m}
}

// Query1 方案1 不做二分查找 直接查找 会超时
func (this *RangeFreqQuery) Query1(left int, right int, value int) int {
	arr, ok := this.cache[value]
	if !ok {
		return 0
	}

	var count int
	for _, v := range arr {
		if v > right {
			// 找完区间后快速退出
			return count
		}
		if v < left {
			continue
		}
		count++
	}

	return count
}

// Query 方案2 做二分查找
func (this *RangeFreqQuery) Query(left int, right int, value int) int {
	arr, ok := this.cache[value]
	if !ok {
		return 0
	}

	// 查找第一个大于等于left边界的位置
	l := lowerBound(arr, left)

	// 查找第一个大于right边界的位置
	r := upperBound(arr, right)

	return r - l
}

func lowerBound(arr []int, target int) int {
	var left, right = 0, len(arr) - 1
	for left <= right {
		var mid = left + (right-left)/2
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}

func upperBound(arr []int, target int) int {
	var left, right = 0, len(arr) - 1
	for left <= right {
		var mid = left + (right-left)/2
		if arr[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}

/**
 * Your RangeFreqQuery object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,value);
 */
