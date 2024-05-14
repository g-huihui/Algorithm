/**
 * @Author: Gong Yanhui
 * @Description: 2244. 完成所有任务需要的最少轮数
 * @Date: 2024-05-14 10:52
 */

package main

func minimumRounds(tasks []int) int {
	var m = make(map[int]int)
	for _, num := range tasks {
		count, ok := m[num]
		if ok {
			m[num] = count + 1
		} else {
			m[num] = 1
		}
	}

	var resCount = 0
	for _, count := range m {
		if count == 1 {
			return -1
		}
		if count%3 == 0 {
			resCount = count/3 + resCount
		} else {
			resCount = count/3 + resCount + 1
		}
	}

	return resCount
}
