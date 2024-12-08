/**
 * @Author: Gong Yanhui
 * @Description: 226. 翻转二叉树
 * @Date: 2024-12-08 10:21
 */

package main

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)

	return root
}

func main() {
	tmp := invertTree(&TreeNode{})
	_ = tmp
}
