/**
 * @Author: Gong Yanhui
 * @Description: 1290. 二进制链表转整数
 * @Date: 2024-05-08 17:39
 */

package main

func getDecimalValue(head *ListNode) int {
	_, value := getLenAndValue(head)
	return value
}

func getLenAndValue(head *ListNode) (int, int) {
	if head.Next == nil {
		return 1, head.Val
	}
	length, value := getLenAndValue(head.Next)
	return length + 1, value + head.Val*(1<<length)
}
