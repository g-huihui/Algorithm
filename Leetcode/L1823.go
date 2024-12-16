/**
 * @Author: Gong Yanhui
 * @Description: 1823. 找出游戏的获胜者
 * @Date: 2024-12-16 23:40
 */

package main

func findTheWinner(n int, k int) int {

	// 模拟队列
	var queue = make([]int, n)
	for i := 0; i < n; i++ {
		queue[i] = i + 1
	}

	var num int // 计数器
	for n != 1 {
		// 将队头数据取出
		tmp := queue[0]
		queue = queue[1:]
		if num == k-1 {
			n--
		} else {
			// 将数据放回队尾
			queue = append(queue, tmp)
		}
		num = (num + 1) % k
	}

	return queue[0]
}
