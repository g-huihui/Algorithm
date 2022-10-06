/**
 * @Author: Gong Yanhui
 * @Description: 111. 二叉树的最小深度
 * @Date: 2022-10-06 20:55
 */

package main

import "math"

//
// minDepth
//  @Description: 深度优先搜索
//  @param root
//  @return int
//
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minRes := math.MaxInt
	if root.Left != nil {
		minRes = min(minDepth(root.Left), minRes)
	}
	if root.Right != nil {
		minRes = min(minDepth(root.Right), minRes)
	}
	// 返回需要包括根节点
	return minRes + 1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
