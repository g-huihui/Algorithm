/**
 * @Author: Gong Yanhui
 * @Description: 1128. 等价多米诺骨牌对的数量
 * @Date: 2025-05-04 20:34
 */

package main

import "fmt"

func numEquivDominoPairs(dominoes [][]int) int {
	var m = make(map[string]int)
	for _, v := range dominoes {
		key1 := fmt.Sprintf("%d-%d", v[0], v[1])
		if _, ok := m[key1]; ok {
			m[key1]++
			continue
		}
		key2 := fmt.Sprintf("%d-%d", v[1], v[0])
		if _, ok := m[key2]; ok {
			m[key2]++
			continue
		}
		m[key1] = 1
	}

	var count int
	for _, v := range m {
		if v > 1 {
			count += f1128(v)
		}
	}

	return count
}

func f1128(n int) int {
	var res int
	for i := 1; i < n; i++ {
		res += i
	}
	return res
}

// 1 <= dominoes.length <= 4 * 104 超时
func numEquivDominoPairs1(dominoes [][]int) int {
	var count int
	for i := 0; i < len(dominoes)-1; i++ {
		for j := i + 1; j < len(dominoes); j++ {
			if (dominoes[i][0] == dominoes[j][0] && dominoes[i][1] == dominoes[j][1]) ||
				(dominoes[i][0] == dominoes[j][1] && dominoes[i][1] == dominoes[j][0]) {
				count++
			}
		}
	}

	return count
}
