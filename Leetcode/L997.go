/**
 * @Author: Gong Yanhui
 * @Description: 997. 找到小镇的法官
 * @Date: 2024-09-22 13:24
 */

package main

func findJudge(n int, trust [][]int) int {
	if n == 1 && len(trust) == 0 {
		return 1
	}

	var arr = make([]int, n+1)
	for i := 0; i < len(trust); i++ {
		arr[trust[i][1]]++
	}

	for i, count := range arr {
		if count == n-1 {
			var flag = true
			for _, ints := range trust {
				if i == ints[0] {
					flag = false
					break
				}
			}
			if flag {
				return i
			}
		}
	}

	return -1
}

// failed 2
func findJudge2(n int, trust [][]int) int {
	var err = -1

	var arr = make([]int, n+1)
	for i := 0; i < len(trust); i++ {
		if arr[trust[i][0]] == err {
			return err
		}
		arr[trust[i][0]] = err
	}

	var m = make(map[int]struct{})
	for i := 1; i < len(arr); i++ {
		if arr[i] != err {
			m[i] = struct{}{}
		}
	}

	if len(m) != 1 {
		return err
	}

	for v, _ := range m {
		return v
	}

	return err
}

// failed 1
func findJudge1(n int, trust [][]int) int {
	var res = -1
	if n-1 != len(trust) {
		return res
	}

	for i := 0; i < len(trust)-1; i++ {
		if trust[i][1] != trust[i+1][1] {
			return res
		}
	}

	return trust[0][1]
}

func main() {
	println(findJudge(2, [][]int{{1, 2}}))                 // 2
	println(findJudge(3, [][]int{{1, 3}, {2, 3}}))         // 3
	println(findJudge(3, [][]int{{1, 3}, {2, 3}, {3, 1}})) // -1
	println(findJudge(3, [][]int{{1, 2}, {2, 3}}))         // -1
}
