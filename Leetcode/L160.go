/**
 * @Author: Gong Yanhui
 * @Description: 160. 相交链表
 * @Date: 2022-10-22 16:40
 */

package main

//
// getIntersectionNode1
//  @Description: map集合解法
//  @param headA
//  @param headB
//  @return *ListNode
//
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	hash := make(map[*ListNode]bool)
	for ; headA != nil; headA = headA.Next {
		hash[headA] = true
	}
	for ; headB != nil; headB = headB.Next {
		if hash[headB] {
			return headB
		}
	}
	return nil
}

//
// getIntersectionNode2
//  @Description: 双指针
//  @param headA
//  @param headB
//
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}
