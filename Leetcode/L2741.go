/**
 * @Author: Gong Yanhui
 * @Description: 2741. 特别的排列
 * @Date: 2024-06-26 21:17
 */

package main

func specialPerm(nums []int) int {
	var count int

	for i := 0; i < len(nums); i++ {
		var flag = make([]int, len(nums))
		flag[i] = 1
		count += backendFunc(nums, flag, append([]int{}, nums[i]), 0)
	}

	return count % (1000000000 + 7)
}

func backendFunc(nums, flag, cur []int, count int) int {
	if len(cur) == len(flag) {
		return 1
	}

	for i := 0; i < len(nums); i++ {
		if flag[i] == 1 {
			continue
		}
		if nums[i]%cur[len(cur)-1] == 0 || cur[len(cur)-1]%nums[i] == 0 {
			var tmpFlag = make([]int, len(flag))
			copy(tmpFlag, flag)
			tmpFlag[i] = 1
			count += backendFunc(nums, tmpFlag, append(cur, nums[i]), 0)
		}
	}

	return count
}

func main() {
	println(specialPerm([]int{2, 4}))
	println(specialPerm([]int{2, 3, 6}))                     // [3,6,2] 和 [2,6,3]
	println(specialPerm([]int{1, 4, 3}))                     // [3,1,4] 和 [4,1,3]
	println(specialPerm([]int{20, 100, 50, 5, 10, 70, 7}))   // 48
	println(specialPerm([]int{6, 30, 3, 9, 36, 72, 18, 54})) // 1440

	println(specialPerm([]int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192})) // timeout
}
