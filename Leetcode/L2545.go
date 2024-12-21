/**
 * @Author: Gong Yanhui
 * @Description: 2545. 根据第 K 场考试的分数排序
 * @Date: 2024-12-21 10:08
 */

package main

import "sort"

func sortTheStudents(score [][]int, k int) [][]int {
	sort.Slice(score, func(i, j int) bool {
		return score[i][k] > score[j][k]
	})

	return score
}
