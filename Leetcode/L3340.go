/**
 * @Author: Gong Yanhui
 * @Description: 3340. 检查平衡字符串
 * @Date: 2025-03-16 11:23
 */

package main

import "strconv"

func isBalanced3340(num string) bool {
	var odd int  // 奇数和
	var even int // 偶数和
	for i, ch := range num {
		n, _ := strconv.Atoi(string(ch))
		if i%2 == 0 {
			even += n
		} else {
			odd += n
		}
	}

	return odd == even
}
