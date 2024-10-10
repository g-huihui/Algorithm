/**
 * @Author: Gong Yanhui
 * @Description: 3162. 优质数对的总数 I
 * @Date: 2024-10-10 15:15
 */

package main

func numberOfPairs(nums1 []int, nums2 []int, k int) int {
	var count int
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i]%(nums2[j]*k) == 0 {
				count++
			}
		}
	}

	return count
}

func main() {
	println(numberOfPairs([]int{1, 3, 4}, []int{1, 3, 4}, 1))  // 5
	println(numberOfPairs([]int{1, 2, 4, 12}, []int{2, 4}, 3)) // 2
}
