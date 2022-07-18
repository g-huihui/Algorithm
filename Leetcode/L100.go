/**
 * @Author: Gong Yanhui
 * @Description: 100. 相同的树
 * @Date: 2022-07-18 21:57
 */

package main

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	} else if p.Val != q.Val {
		return false
	} else {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
}
