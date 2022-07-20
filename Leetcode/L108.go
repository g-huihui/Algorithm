/**
 * @Author: Gong Yanhui
 * @Description: 108. 将有序数组转换为二叉搜索树
 * @Date: 2022-07-20 21:15
 */

package main

func sortedArrayToBST(nums []int) *TreeNode {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := left + (right-left)/2
	root := &TreeNode{
		Val:   nums[mid],
		Left:  helper(nums, left, mid-1),
		Right: helper(nums, mid+1, right),
	}
	return root
}
