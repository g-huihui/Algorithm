/**
 * @Author: Gong Yanhui
 * @Description: 112. 路径总和
 * @Date: 2022-10-06 21:15
 */

package main

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	targetSum -= root.Val
	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}
