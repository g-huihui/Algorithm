/**
 * @Author: Gong Yanhui
 * @Description: 3206. 交替组 I
 * @Date: 2024-11-26 15:44
 */

package main

func numberOfAlternatingGroups(colors []int) int {
	if len(colors) < 3 {
		return 0
	}
	var res int
	var size = len(colors)
	colors = append(colors, colors[0])
	colors = append(colors, colors[1])

	for i := 0; i < size; i++ {
		if colors[i] != colors[i+1] && colors[i] == colors[i+2] {
			res++
		}
	}

	return res
}
