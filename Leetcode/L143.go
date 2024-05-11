/**
 * @Author: Gong Yanhui
 * @Description: 143. 重排链表
 * @Date: 2024-05-08 18:06
 */

package main

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	var arrList = make([]*ListNode, 0)
	for ; head != nil; head = head.Next {
		arrList = append(arrList, head)
	}
	var left, right = 0, len(arrList) - 1
	var root = new(ListNode)
	var currentNode = root
	for left <= right {
		currentNode.Next = arrList[left]
		currentNode = currentNode.Next
		left++
		if left > right {
			break
		}
		currentNode.Next = arrList[right]
		currentNode = currentNode.Next
		right--
	}
	currentNode.Next = nil
	head = root.Next
}
