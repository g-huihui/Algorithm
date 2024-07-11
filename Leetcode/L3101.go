/**
 * @Author: Gong Yanhui
 * @Description: 3101. 交替子数组计数
 * @Date: 2024-07-11 15:17
 */

package main

// 超出时间限制
func countAlternatingSubarrays1(nums []int) int64 {
	var count int64

	// 双指针记录索引
	for begin := 0; begin < len(nums); begin++ {
		for end := begin; end < len(nums); end++ {
			if begin == end { // 单个数字符合
				count++
				continue
			}
			if nums[end] == nums[end-1] { // 与前一个相同 则后续都不需要再循环
				break
			}
			count++
		}
	}

	return count
}

func main() {
	println(countAlternatingSubarrays([]int{0, 1, 1, 1})) // 5
	println(countAlternatingSubarrays([]int{1, 0, 1, 0})) // 10

	if getSubNumByArrLen(1) != 1 {
		panic("getSubNumByArrLen(1) != 1")
	}
	if getSubNumByArrLen(2) != 3 {
		panic("getSubNumByArrLen(2) != 3")
	}
	if getSubNumByArrLen(3) != 6 {
		panic("getSubNumByArrLen(3) != 6")
	}
	if getSubNumByArrLen(4) != 10 {
		panic("getSubNumByArrLen(4) != 10")
	}
	if getSubNumByArrLen(5) != 15 {
		panic("getSubNumByArrLen(5) != 15")
	}
}

/** 规律 [数字len] 子数组数
[1] 1
[2] 3
[3] 6
[4] 10
[5] 15
[n] (n-1)数+n
*/
// 找到规律后
func countAlternatingSubarrays(nums []int) int64 {
	if len(nums) == 0 || len(nums) == 1 {
		return int64(len(nums))
	}

	var count int64
	var m = make(map[int]int64) // 加一个缓存提升下性能

	for begin, end := 0, 1; ; end++ {
		if end == len(nums) || nums[end] == nums[end-1] {
			var subLen = end - begin
			num, ok := m[subLen]
			if ok {
				count += num
			} else {
				num = int64(getSubNumByArrLen(subLen))
				m[subLen] = num
				count += num
			}
			if end == len(nums) {
				break
			}
			begin = end
		}
	}

	return count
}

func getSubNumByArrLen(len int) int {
	if len == 0 || len == 1 {
		return len
	}
	return getSubNumByArrLen(len-1) + len
}
