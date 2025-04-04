/**
 * @Author: Gong Yanhui
 * @Description: 1123. 最深叶节点的最近公共祖先
 * @Date: 2025-04-04 23:02
 */

package main

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	_, res := f1123(0, root)
	return res
}

func f1123(deep int, root *TreeNode) (int, *TreeNode) {
	// 找到叶子了 从0开始退回加
	if root == nil {
		return 0, nil
	}

	leftDeep, l := f1123(deep, root.Left)
	rightDeep, r := f1123(deep, root.Right)

	if leftDeep > rightDeep {
		return leftDeep + 1, l
	}
	if rightDeep > leftDeep {
		return rightDeep + 1, r
	}

	return leftDeep + 1, root
}

func main() {
	var root = &TreeNode{Val: 0, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 3}}
	q1 := lcaDeepestLeaves(root)
	println(q1)
}
