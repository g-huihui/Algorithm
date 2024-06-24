/**
 * @Author: Gong Yanhui
 * @Description: 496. 下一个更大元素 I
 * @Date: 2024-06-24 21:21
 */

package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {

	var m = make(map[int]int) // key: num => value: index
	for i := 0; i < len(nums2); i++ {
		m[nums2[i]] = i
	}

	var res []int
	for i := 0; i < len(nums1); i++ {
		index := m[nums1[i]]
		if index == len(nums2)-1 {
			res = append(res, -1)
		} else if num, ok := nums2AfterBiggest(nums2, index); ok {
			res = append(res, num)
		} else {
			res = append(res, -1)
		}
	}

	return res
}

func nums2AfterBiggest(nums []int, index int) (int, bool) {
	for i := index + 1; i < len(nums); i++ {
		if nums[i] > nums[index] {
			return nums[i], true
		}
	}

	return 0, false
}

func main() {
	res1 := nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}) // [-1, 3, -1]
	println(res1)

	res2 := nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}) // [3, -1]
	println(res2)

	res3 := nextGreaterElement([]int{1, 3, 5, 2, 4}, []int{6, 5, 4, 3, 2, 1, 7}) // [7, 7, 7, 7, 7]
	println(res3)
}
