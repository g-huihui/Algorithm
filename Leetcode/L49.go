/**
 * @Author: Gong Yanhui
 * @Description: 49. 字母异位词分组
 * @Date: 2024-12-05 17:01
 */

package main

import "sort"

func groupAnagrams(strs []string) [][]string {

	var m = make(map[string][]string)
	for index := range strs {
		var temp = []byte(strs[index])
		sort.Slice(temp, func(i, j int) bool {
			return temp[i] < temp[j]
		})

		m[string(temp)] = append(m[string(temp)], strs[index])
	}

	var res = make([][]string, 0)
	for _, value := range m {
		res = append(res, value)
	}

	return res
}
