/**
 * @Author: Gong Yanhui
 * @Description: 报数游戏
 * @Date: 2024-12-16 17:31
 */

package main

import (
	"fmt"
	"sort"
)

/**
1. 题目描述:
100个人围成一圈,每个人有一个编码编号从一开始到一百;他们从一开始依次报数,报道M的人自动退出圈圈,然后下一个人接着从1开始报数一直到剩余人数小于M.
请问最后剩余人在原先的编码为多少?

2. 输入描述
输入一个整数参数M,如果输入参数M小于等于1或者大于等于100,输出"ERROR!";否则按原先的编号从小到大的顺序,以英文逗号分割输出编号字符串

3. 输入输出
input : 3
output:58,91
input : 4
output:34,45,97
*/

func H1(M int) []int {

	var size = 100
	// 使用切片来代替队列使用
	var queue = make([]int, size)
	// 将所有人入队列
	for i := 0; i < size; i++ {
		queue[i] = i + 1
	}

	var num int // 报数的计数器
	for size >= M {
		// 获取队头数据
		var tmp = queue[0]
		// 出队列
		queue = queue[1:]

		if num == M-1 {
			size--
		} else {
			// 入队列
			queue = append(queue, tmp)
		}
		num = (num + 1) % M
	}

	// 排序后返回
	sort.Ints(queue)
	return queue
}

func main() {
	fmt.Println(H1(3)) // 58,91
	fmt.Println(H1(4)) // 34,45,97
}

// H1_ 官方解法
func H1_(M int) []int {
	size := 100

	queue := make([]int, size)
	for i := 0; i < size; i++ {
		queue[i] = i + 1
	}

	for size >= M {
		for i := 0; i < M-1; i++ {
			tmp := queue[0]
			queue = queue[1:]
			queue = append(queue, tmp)
		}
		queue = queue[1:]
		size--
	}

	sort.Ints(queue)
	return queue
}
