/**
 * @Author: Gong Yanhui
 * @Description: 572. 另一棵树的子树
 * @Date: 2024-08-04 22:02
 */

package main

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root != nil && subRoot != nil {
		if root.Val == subRoot.Val && isEqualTree(root, subRoot) {
			return true
		} else {
			return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
		}
	}
	if root == nil || subRoot == nil {
		return false
	}

	return true
}

func isEqualTree(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}

	return t1.Val == t2.Val && isEqualTree(t1.Left, t2.Left) && isEqualTree(t1.Right, t2.Right)
}
