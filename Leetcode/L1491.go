/**
 * @Author: Gong Yanhui
 * @Description: 1491. 去掉最低工资和最高工资后的工资平均值
 * @Date: 2024-05-03 16:05
 */

package main

import "sort"

func average(salary []int) float64 {
	sort.Ints(salary)
	salary = salary[1 : len(salary)-1]
	var sum int
	for _, v := range salary {
		sum += v
	}
	return float64(sum) / float64(len(salary))
}
