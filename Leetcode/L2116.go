/**
 * @Author: Gong Yanhui
 * @Description: 2116. 判断一个括号字符串是否有效
 * @Date: 2025-03-23 12:18
 */

package main

func main() {
	println(canBeValid("))()))", "010100")) // true
}

// 灵茶山艾府
func canBeValid(s string, locked string) bool {
	var size = len(s)
	if size%2 == 1 {
		// 奇数无效
		return false
	}

	var minLeft, maxLeft int // 记录左括号最大最小值
	for i := 0; i < size; i++ {
		if locked[i] == '1' { // 不可更改
			var diff = -1 // 左括号使用1记录增加 右括号使用-1记录减少
			if s[i] == '(' {
				diff = 1
			}
			maxLeft += diff
			if maxLeft < 0 {
				// 此时右括号过多 无效 直接返回
				return false
			}
			minLeft += diff
			if minLeft < 0 { // 当最大括号值 maxLeft 还有效时 此时左括号数为奇数 改为最小奇数 1 (这里还是有点难理解)
				minLeft = 1
			}
		} else { // 可以为 左括号 or 右括号
			// 当做左括号
			maxLeft++
			// 当做右括号
			minLeft--
			if minLeft < 0 { // 同理上述
				minLeft = 1
			}
		}
	}

	return minLeft == 0
}

// 第一次尝试
func canBeValid1(s string, locked string) bool {
	var size = len(s)
	if size%2 == 1 {
		// 奇数无效
		return false
	}
	// 创建一个栈
	var stack = make([]int, 0, size)
	for i := 0; i < size; i++ {
		if s[i] == ')' {
			// 判断栈中是否有前括号将其弹出
			if stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				// 判断当前位置是否可以转换为前括号入栈
				if locked[i] == '0' {
					stack = append(stack, '(')
				} else {
					return false
				}
			}
		} else {
			// 这里不能将 '(' 无脑往栈压
		}
	}

	return true
}
