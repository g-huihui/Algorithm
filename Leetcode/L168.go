/**
 * @Author: Gong Yanhui
 * @Description: 168. Excel表列名称
 * @Date: 2024-01-09 16:56
 */

package main

func convertToTitle(columnNumber int) string {
	var res string
	for columnNumber > 0 {
		char := num2char((columnNumber-1)%26 + 1)
		res = string(char) + res
		columnNumber = (columnNumber - 1) / 26
	}
	return res
}

func num2char(num int) rune {
	return rune(num + 64)
}

func main() {
	println(num2char(1))

	println(convertToTitle(28))
	println(convertToTitle(701))
	println(convertToTitle(52))
}
