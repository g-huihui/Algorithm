/**
 * @Author: Gong Yanhui
 * @Description: 3164. 优质数对的总数 II
 * @Date: 2024-10-11 21:16
 */

package main

// 反向做乘法
func numberOfPairs2(nums1 []int, nums2 []int, k int) int64 {
	var m1, m2 = make(map[int]int), make(map[int]int)
	var nums1Max int
	for _, v := range nums1 {
		m1[v]++
		if v > nums1Max {
			nums1Max = v
		}
	}
	for _, v := range nums2 {
		m2[v]++
	}

	var res int64
	for num, cnt := range m2 {
		for num2 := num * k; num2 <= nums1Max; num2 += num * k {
			if count, ok := m1[num2]; ok {
				res += int64(count) * int64(cnt)
			}
		}
	}

	return res
}

// 超出时间限制
func numberOfPairs2_timeout(nums1 []int, nums2 []int, k int) int64 {
	var m1, m2 = make(map[int]int), make(map[int]int)
	for _, v := range nums1 {
		m1[v]++
	}
	for _, v := range nums2 {
		m2[v]++
	}

	var res int64
	for k1, v1 := range m1 {
		for k2, v2 := range m2 {
			if k1%(k2*k) == 0 {
				res += int64(v1) * int64(v2)
			}
		}
	}

	return res
}

func main() {
	println(numberOfPairs2([]int{1, 3, 4}, []int{1, 3, 4}, 1))  // 5
	println(numberOfPairs2([]int{1, 2, 4, 12}, []int{2, 4}, 3)) // 2
}
