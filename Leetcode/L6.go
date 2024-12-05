/**
 * @Author: Gong Yanhui
 * @Description: 6. Z 字形变换
 * @Date: 2024-11-30 15:24
 */

package main

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	var arr = make([]string, numRows)
	var circleNum = numRows*2 - 2 // 一个循环的长度

	for index, ch := range s {
		tmpNum := index % circleNum
		if tmpNum < numRows {
			arr[tmpNum] += string(ch)
		} else {
			arr[circleNum-tmpNum] += string(ch)
		}
	}

	// 拼接结果
	var res string
	for _, a := range arr {
		res += a
	}

	return res
}
