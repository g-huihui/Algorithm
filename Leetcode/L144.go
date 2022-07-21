/**
 * @Author: Gong Yanhui
 * @Description: 144. 二叉树的前序遍历
 * @Date: 2022-07-21 20:06
 */

package main

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0, 0)
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 前序遍历
		res = append(res, node.Val)
		inorder(node.Left)
		inorder(node.Right)
	}

	inorder(root)
	return res
}
