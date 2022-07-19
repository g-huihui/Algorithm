/**
 * @Author: Gong Yanhui
 * @Description: 101. 对称二叉树
 * @Date: 2022-07-19 19:42
 */

package main

func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

func check(q *TreeNode, p *TreeNode) bool {
	if q == nil && p == nil {
		return true
	}
	if q == nil || p == nil {
		return false
	}
	return q.Val == p.Val && check(q.Left, p.Right) && check(q.Right, p.Left)
}
