/**
 * @Author: Gong Yanhui
 * @Description: 单词接龙
 * @Date: 2024-12-21 16:27
 */

package main

import (
	"fmt"
	"sort"
)

/**
1. 题目描述
可用于接龙的单词首字母必须要前一个单词的尾字母相同
当存在多个首字母相同的单词时 取长度最长的单词 如果长度也相等 则取字典序最小的单词(已经参与接龙的单词不能重复使用)
现给定一组全部由小写字母组成单词数组 并指定其中的一个单词作为起始单词 进行单词接龙
请输出最长的单词串 单词串是单词拼接而成 中间没有空格

2. 输入描述
index	起始单词索引下标
arr		单词组

3. 输入输出
input : 0 ["word","dd","da","dc","dword","d"]
output: worddwordda (word dword da)
input : 4 ["word","dd","da","dc","dword","d"]
output: dwordda (dword da)
*/

func H4(index int, arr []string) string {

	// 使用map存储每个单词的开始字符 并且需要将每个字符串数组按要求排序
	var tmp = make(map[uint8][]string)
	for i, str := range arr {
		if i == index { // 排除起始单词
			continue
		}
		tmp[str[0]] = append(tmp[str[0]], str)
	}
	var m = make(map[byte][]string)
	for b, strs := range tmp {
		sort.Slice(strs, func(i, j int) bool {
			if len(strs[i]) > len(strs[j]) {
				return true
			} else if len(strs[i]) == len(strs[j]) {
				return strs[i] < strs[j]
			}

			return false
		})
		m[b] = strs
	}

	var res = arr[index]
	var last = res[len(res)-1]
	for len(m[last]) > 0 {
		strs := m[last]
		res += strs[0]
		m[last] = strs[1:]
		last = strs[0][len(strs[0])-1]
	}

	return res
}

func main() {
	if H4(0, []string{"word", "dd", "da", "dc", "dword", "d"}) != "worddwordda" {
		println("err: 0 [\"word\",\"dd\",\"da\",\"dc\",\"dword\",\"d\"]")
	}
	if H4(4, []string{"word", "dd", "da", "dc", "dword", "d"}) != "dwordda" {
		println("err: 4, []string{\"word\", \"dd\", \"da\", \"dc\", \"dword\", \"d\"}")
	}

	fmt.Println("如果没有其他数据输出则为正常!")
}
