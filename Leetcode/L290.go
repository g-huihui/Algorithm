/**
 * @Author: Gong Yanhui
 * @Description: 290. 单词规律
 * @Date: 2024-06-17 16:44
 */

package main

import "strings"

func wordPattern(pattern string, s string) bool {
	split := strings.Split(s, " ")
	if len(pattern) != len(split) {
		return false
	}

	var s2p = make(map[string]byte)
	var p2s = make(map[byte]string)
	for i := 0; i < len(split); i++ {
		b, ok1 := s2p[split[i]]
		str, ok2 := p2s[pattern[i]]
		if ok1 && ok2 && b == pattern[i] && str == split[i] {
			continue
		} else if !ok1 && !ok2 {
			s2p[split[i]] = pattern[i]
			p2s[pattern[i]] = split[i]
		} else {
			return false
		}
	}

	return true
}

func main() {
	println(wordPattern("abba", "dog cat cat dog"))  // true
	println(wordPattern("abba", "dog cat cat fish")) // false
	println(wordPattern("aaaa", "dog cat cat dog"))  // false
	println(wordPattern("abba", "dog dog dog dog"))  // false
}
