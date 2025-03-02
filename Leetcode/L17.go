/**
 * @Author: Gong Yanhui
 * @Description: 17. 电话号码的字母组合
 * @Date: 2025-03-02 12:52
 */

package main

import "fmt"

func letterCombinations(digits string) []string {
	// 数字键盘映射
	var MAPPING = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

	if len(digits) == 0 {
		return []string{}
	}

	var res []string // 结果集
	var path string  // 路径
	dfs17(&MAPPING, digits, 0, path, &res)

	return res
}

func dfs17(MAPPING *[]string, digits string, index int, path string, res *[]string) {
	// 递归终止条件
	if index == len(digits) {
		*res = append(*res, path)
		return
	}

	// 处理当前层逻辑
	var digit = digits[index] - '0' // 遍历的当前数字 - '0' 转换为int
	var letters = (*MAPPING)[digit] // 当前数字对应的字符串
	for i := 0; i < len(letters); i++ {
		// 进入下一层
		dfs17(MAPPING, digits, index+1, path+string(letters[i]), res)
	}
}

func main() {
	// 输入：digits = "23"
	// 输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
	fmt.Println(letterCombinations("23"))

	// 输入：digits = ""
	// 输出：[]
	fmt.Println(letterCombinations(""))

	// 输入：digits = "2"
	// 输出：["a","b","c"]
	fmt.Println(letterCombinations("2"))
}
