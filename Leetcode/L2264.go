/**
 * @Author: Gong Yanhui
 * @Description: 2264. 字符串中最大的 3 位相同数字
 * @Date: 2025-01-08 13:00
 */

package main

func largestGoodInteger(num string) string {
	var res string
	for i := 0; i < len(num)-2; i++ {
		if num[i] == num[i+1] && num[i] == num[i+2] {
			var tmp = num[i : i+3]
			if res == "" || tmp > res {
				res = tmp
			}
		}
	}
	return res
}
