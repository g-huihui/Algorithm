/**
 * @Author: Gong Yanhui
 * @Description: 110. 平衡二叉树
 * @Date: 2022-10-06 20:45
 */

package main

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func treeHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(treeHeight(root.Left), treeHeight(root.Right)) + 1
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(treeHeight(root.Left)-treeHeight(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}
