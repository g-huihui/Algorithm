/**
 * @Author: Gong Yanhui
 * @Description: 2181. 合并零之间的节点
 * @Date: 2024-09-09 20:03
 */

package main

func mergeNodes(head *ListNode) *ListNode {
	if head.Val != 0 {
		return head
	}

	var root = new(ListNode)
	var cur = root
	var total int
	for tmp := head; tmp != nil; tmp = tmp.Next {
		if tmp.Val == 0 {
			if total != 0 {
				cur.Next = &ListNode{Val: total}
				cur = cur.Next
				total = 0
			}
		} else {
			total += tmp.Val
		}
	}

	return root.Next
}
