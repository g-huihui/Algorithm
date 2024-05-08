/**
 * @Author: Gong Yanhui
 * @Description: LCR 023. 相交链表
 * @Date: 2024-05-08 17:53
 */

package LCR

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var m = make(map[*ListNode]bool)
	for tmp := headA; tmp != nil; tmp = tmp.Next {
		m[tmp] = true
	}
	for tmp := headB; tmp != nil; tmp = tmp.Next {
		if m[tmp] {
			return tmp
		}
	}
	return nil
}
