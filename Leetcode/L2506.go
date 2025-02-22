/**
 * @Author: Gong Yanhui
 * @Description: 2506. 统计相似字符串对的数目
 * @Date: 2025-02-22 12:24
 */

package main

func similarPairs(words []string) int {
	var size = len(words)
	var m = make([]map[uint8]struct{}, 0, size)
	for i := 0; i < size; i++ {
		m = append(m, str2map2506(words[i]))
	}

	var count int
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			if isEqual2506(m[i], m[j]) {
				count++
			}
		}
	}

	return count
}

// 将str字符串转换为map map的key为出现的字符
func str2map2506(str string) (m map[uint8]struct{}) {
	m = make(map[uint8]struct{})
	var size = len(str)
	for i := 0; i < size; i++ {
		m[str[i]] = struct{}{}
	}
	return
}

func isEqual2506(m1, m2 map[uint8]struct{}) bool {
	if len(m1) != len(m2) {
		return false
	}
	for v := range m1 {
		if _, ok := m2[v]; !ok {
			return false
		}
	}
	return true
}

func main() {
	println(similarPairs([]string{"aba", "aabb", "abcd", "bac", "aabc"})) // 2
	println(similarPairs([]string{"aabb", "ab", "ba"}))                   // 3
	println(similarPairs([]string{"nba", "cba", "dba"}))                  // 0
}
