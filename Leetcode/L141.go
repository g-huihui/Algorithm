/**
 * @Author: Gong Yanhui
 * @Description: 141. 环形链表
 * @Date: 2022-10-07 22:01
 */

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	var (
		slow = head
		fast = head.Next
	)
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}
