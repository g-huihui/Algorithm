/**
 * @Author: Gong Yanhui
 * @Description: 205. 同构字符串
 * @Date: 2024-06-15 12:57
 */

package main

func isIsomorphic(s string, t string) bool {

	var sMap = make(map[byte]int)
	var tMap = make(map[byte]int)

	for index, _ := range s {
		b1, ok1 := sMap[s[index]]
		b2, ok2 := tMap[t[index]]
		if ok1 && ok2 && b1 == b2 {
			// 都存在并且相同
		} else if !ok1 && !ok2 {
			// 都不存在
			sMap[s[index]] = index
			tMap[t[index]] = index
		} else {
			return false
		}
	}
	return true
}

func main() {
	println(isIsomorphic("egg", "add"))     //true
	println(isIsomorphic("foo", "bar"))     //false
	println(isIsomorphic("paper", "title")) //true
}
