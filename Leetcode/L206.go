/**
 * @Author: Gong Yanhui
 * @Description: 206. 反转链表
 * @Date: 2024-05-07 10:51
 */

package main

func reverseList(head *ListNode) *ListNode {
	//if head == nil || head.Next == nil {
	//	return head
	//}
	var root = new(ListNode)
	for {
		if head == nil {
			break
		}
		// 将当前节点从原始链表中取下
		cur := head
		head = head.Next
		// 将当前cur节点加入root中
		cur.Next = root.Next
		root.Next = cur
	}
	return root.Next
}
