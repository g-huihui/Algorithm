/**
 * @Author: Gong Yanhui
 * @Description: 445. 两数相加 II
 * @Date: 2025-03-06 23:02
 */

package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	// 使用两个栈 将链表压入栈
	s1, s2 := list2Stack445(l1), list2Stack445(l2)

	// 创建一个新的返回链表
	var head = &ListNode{}
	var flag bool // 是否进位
	for len(s1) > 0 || len(s2) > 0 || flag {
		var tmp = &ListNode{}
		if len(s1) > 0 {
			tmp.Val += s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
		}
		if len(s2) > 0 {
			tmp.Val += s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
		}
		if flag {
			tmp.Val += 1
			flag = false
		}
		if tmp.Val > 9 {
			tmp.Val %= 10
			flag = true
		}

		// 使用头插法将新的节点放入返回值中
		var n = head.Next
		head.Next = tmp
		tmp.Next = n
	}

	return head.Next
}

func list2Stack445(l *ListNode) []int {
	var res = make([]int, 0)
	for l != nil {
		res = append(res, l.Val)
		l = l.Next
	}

	return res
}
