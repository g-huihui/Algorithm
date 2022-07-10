/**
 * @Author: Gong Yanhui
 * @Description: 20. 有效的括号
 * @Date: 2022-07-10 23:09
 */

package main

import "fmt"

func isValid(s string) bool {
	length := len(s)
	if length%2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{} //存放的是前括号类型
	for i := 0; i < length; i++ {
		tmp := s[i] // 获取字符串的每一个字符

		if pairs[tmp] == 0 { //说明是前括号
			stack = append(stack, tmp)
		} else {
			if len(stack) == 0 || pairs[tmp] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()"))     // true
	fmt.Println(isValid("()[]{}")) // true
	fmt.Println(isValid("(]"))     // false
	fmt.Println(isValid("([)]"))   // false
}
