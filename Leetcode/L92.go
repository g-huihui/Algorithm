/**
 * @Author: Gong Yanhui
 * @Description: 92. 反转链表 II
 * @Date: 2024-05-07 11:04
 */

package main

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var dummyRoot = new(ListNode)
	dummyRoot.Next = head
	var currentNode = dummyRoot
	for i := 0; i < left-1; i++ {
		if currentNode.Next == nil {
			return head
		}
		currentNode = currentNode.Next
	}
	// 截断
	preNode := currentNode
	currentNode = currentNode.Next
	preNode.Next = nil
	for i := 0; i <= right-left; i++ { // 反转操作
		tmp := currentNode
		currentNode = currentNode.Next
		tmp.Next = preNode.Next
		preNode.Next = tmp
	}
	// 将剩下的节点拼接到后面
	for {
		if preNode.Next == nil {
			preNode.Next = currentNode
			break
		}
		preNode = preNode.Next
	}
	return dummyRoot.Next
}

// 解决不了left=1开始的问题
func reverseBetween2(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil || left == right {
		return head
	}

	var cur = head
	var root = new(ListNode)
	root.Next = cur

	// 拆分第一段
	for i := 1; i < left-1; i++ {
		if cur == nil {
			return head
		}
		cur = cur.Next
	}

	// 拆分第二段
	var oneEnd = cur
	cur = cur.Next
	oneEnd.Next = nil
	var two = cur
	for i := 1; i <= right-left; i++ {
		if cur == nil {
			return head
		}
		cur = cur.Next
	}

	// 拆分第三段
	var three = cur.Next
	cur.Next = nil

	// 反转第二段
	two = reverseList(two)

	//将三段重新拼接
	oneEnd.Next = two
	for {
		if two.Next == nil {
			two.Next = three
			break
		}
		two = two.Next
	}

	return head
}

// 反转链表 通过判断翻转value=Left和value=Right之间的链表
func reverseBetween1(head *ListNode, left int, right int) *ListNode {
	var cur = new(ListNode)

	// 将链表切割为三段

	// 获取第一段 结尾pre
	var pre = new(ListNode)
	pre.Next = head
	for {
		if pre.Next == nil {
			return head
		}
		if pre.Next.Val == left {
			break
		}
		pre = pre.Next
	}

	// 获取第二段
	var two = pre.Next
	pre.Next = nil
	cur = two
	for {
		if cur == nil {
			break
		}
		if cur.Val == right {
			break
		}
		cur = cur.Next
	}

	// 获取第三段
	var three = new(ListNode)
	if cur != nil {
		three = cur.Next
		cur.Next = nil
	}

	// 翻转第二段
	two = reverseList(two)

	//将三段重新拼接
	pre.Next = two
	for {
		if two.Next == nil {
			two.Next = three
			break
		}
		two = two.Next
	}

	return head
}

func main() {
	// 1 2 3 4 5 --> 1 4 3 2 5
	var head = new(ListNode)
	head.Val = 1
	head.Next = new(ListNode)
	head.Next.Val = 2
	head.Next.Next = new(ListNode)
	head.Next.Next.Val = 3
	head.Next.Next.Next = new(ListNode)
	head.Next.Next.Next.Val = 4
	head.Next.Next.Next.Next = new(ListNode)
	head.Next.Next.Next.Next.Val = 5
	reverseBetween(head, 2, 4)

	// 3 5 --> 5 3
	head = new(ListNode)
	head.Val = 3
	head.Next = new(ListNode)
	head.Next.Val = 5
	reverseBetween(head, 1, 2)

	// 3 5 --> 3 5
	head = new(ListNode)
	head.Val = 3
	head.Next = new(ListNode)
	head.Next.Val = 5
	reverseBetween(head, 1, 1)
}
