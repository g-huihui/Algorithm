/**
 * @Author: Gong Yanhui
 * @Description: 43. 字符串相乘
 * @Date: 2025-02-16 18:08
 */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 常规的竖列乘法方式
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	var size1, size2 = len(num1), len(num2)
	if size1 > size2 { // 确保第一个字符串长度更短
		return multiply(num2, num1)
	}

	var total string
	for i := 0; i < size1; i++ {
		tmp, _ := strconv.Atoi(string(num1[size1-i-1]))
		var v = multiplySingle(num2, tmp) + strings.Repeat("0", i) // 单独值乘法结果
		total = subString(total, v)
	}

	return total
}

// 一个字符串数字和单个数字相乘	"123"*2
func multiplySingle(str string, num int) string {
	var size, index = len(str), 0
	var flagNum = 0 // 进位的数字
	var res string
	for index < size {
		n, _ := strconv.Atoi(string(str[size-index-1]))
		n = n * num
		if flagNum > 0 {
			n += flagNum
			flagNum = 0
		}
		if n > 9 {
			flagNum = n / 10
			n %= 10
		}
		res = fmt.Sprintf("%d", n) + res
		index++
	}

	if flagNum > 0 {
		res = fmt.Sprintf("%d", flagNum) + res
	}

	return res
}

// 两个字符串数字相加
func subString(num1, num2 string) string {
	var size1, size2 = len(num1), len(num2)
	var size int
	if size1 < size2 {
		size = size1
	} else {
		size = size2
	}

	var flag = false // 是否进位标记
	var res string
	var index = 0
	for index < size {
		n1, _ := strconv.Atoi(string(num1[size1-index-1]))
		n2, _ := strconv.Atoi(string(num2[size2-index-1]))
		tmp := n1 + n2
		if flag {
			tmp += 1
			flag = false
		}
		if tmp > 9 {
			flag = true
			tmp %= 10
		}
		res = fmt.Sprintf("%d", tmp) + res
		index++
	}

	for index < size1 {
		n1, _ := strconv.Atoi(string(num1[size1-index-1]))
		if flag {
			n1 += 1
			flag = false
		}
		if n1 > 9 {
			flag = true
			n1 %= 10
		}
		res = fmt.Sprintf("%d", n1) + res
		index++
	}

	for index < size2 {
		n2, _ := strconv.Atoi(string(num2[size2-index-1]))
		if flag {
			n2 += 1
			flag = false
		}
		if n2 > 9 {
			flag = true
			n2 %= 10
		}
		res = fmt.Sprintf("%d", n2) + res
		index++
	}

	if flag {
		res = "1" + res
	}

	return res
}

func main() {
	fmt.Println(subString("1", "1")) // 2
	fmt.Println(subString("1", "9")) // 10
	fmt.Println(subString("0", "1")) // 1

	fmt.Println(subString("10", "1"))  // 11
	fmt.Println(subString("99", "1"))  // 100
	fmt.Println(subString("99", "99")) // 198

	fmt.Println(subString("0", "738")) // 198

	// 输入: num1 = "2", num2 = "3"
	// 输出: "6"
	fmt.Println(multiply("2", "3"))

	// 输入: num1 = "123", num2 = "456"
	// 输出: "56088"
	fmt.Println(multiply("123", "456"))
	fmt.Println(multiply("456", "123"))
}
