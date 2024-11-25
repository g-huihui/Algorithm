/**
 * @Author: Gong Yanhui
 * @Description: 3238. 求出胜利玩家的数目
 * @Date: 2024-11-25 20:37
 */

package main

func winningPlayerCount(n int, pick [][]int) int {
	var res = make(map[int]struct{})
	var flag = make(map[int]map[int]int)

	for _, its := range pick {
		if _, ok := flag[its[0]]; !ok {
			flag[its[0]] = make(map[int]int)
		}
		flag[its[0]][its[1]]++
		if flag[its[0]][its[1]] > its[0] {
			res[its[0]] = struct{}{}
		}
	}

	return len(res)
}
