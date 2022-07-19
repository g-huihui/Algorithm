/**
 * @Author: Gong Yanhui
 * @Description: 104. 二叉树的最大深度
 * @Date: 2022-07-19 19:50
 */

package main

func maxDepth(root *TreeNode) int {
	return checkDeep(root, 0)
}

func checkDeep(root *TreeNode, deep int) int {
	if root == nil {
		return deep
	}
	deep++
	return max(checkDeep(root.Left, deep), checkDeep(root.Right, deep))
}
